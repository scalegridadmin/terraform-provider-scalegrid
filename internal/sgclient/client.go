// Package sgclient provides the HTTP client shared by generated Terraform
// resources and data sources when talking to sg-api.
package sgclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	defaultEndpoint = "https://scalegrid-api.scalegrid.ai"
	tokenTTL        = 50 * time.Minute
)

var tokenCache sync.Map

type cachedToken struct {
	value     string
	expiresAt time.Time
}

// Config contains the provider-level connection settings for sg-api.
type Config struct {
	Endpoint    string
	BearerToken string
	Username    string
	Password    string
	HTTPClient  *http.Client
}

// Client wraps an authenticated HTTP client and the sg-api endpoint it targets.
type Client struct {
	Endpoint   string
	HTTPClient *http.Client
}

// New builds an authenticated sg-api client for exactly one credential mode.
func New(cfg Config) (*Client, error) {
	endpoint, err := resolveEndpoint(cfg.Endpoint)
	if err != nil {
		return nil, err
	}

	if err := validateCredentials(cfg); err != nil {
		return nil, err
	}

	baseClient := http.DefaultClient
	if cfg.HTTPClient != nil {
		baseClient = cfg.HTTPClient
	}

	baseTransport := baseClient.Transport
	if baseTransport == nil {
		baseTransport = http.DefaultTransport
	}

	mintClient := *baseClient
	mintClient.Transport = baseTransport

	client := *baseClient
	client.Transport = &authTransport{
		base:       baseTransport,
		mintClient: &mintClient,
		endpoint:   endpoint,
		cfg:        cfg,
	}

	return &Client{
		Endpoint:   endpoint,
		HTTPClient: &client,
	}, nil
}

// NewRequest builds a request relative to the configured sg-api endpoint.
func (c *Client) NewRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	if c == nil {
		return nil, errors.New("sgclient: nil client")
	}

	endpoint, err := url.Parse(c.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("sgclient: invalid endpoint: %w", err)
	}

	rel, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("sgclient: invalid request path: %w", err)
	}

	return http.NewRequestWithContext(ctx, method, endpoint.ResolveReference(rel).String(), body)
}

func resolveEndpoint(configured string) (string, error) {
	endpoint := firstNonEmpty(configured, os.Getenv("SCALEGRID_ENDPOINT"), os.Getenv("SG_API_BASE_URL"), defaultEndpoint)
	parsed, err := url.Parse(endpoint)
	if err != nil {
		return "", fmt.Errorf("sgclient: invalid endpoint: %w", err)
	}
	if parsed.Scheme == "" || parsed.Host == "" {
		return "", fmt.Errorf("sgclient: endpoint must include scheme and host")
	}
	return strings.TrimRight(endpoint, "/"), nil
}

func validateCredentials(cfg Config) error {
	if (cfg.Username == "") != (cfg.Password == "") {
		return errors.New("sgclient: username and password must be configured together")
	}

	modeCount := 0
	for _, configured := range []bool{
		cfg.BearerToken != "",
		cfg.Username != "" && cfg.Password != "",
	} {
		if configured {
			modeCount++
		}
	}

	if modeCount != 1 {
		return errors.New("sgclient: exactly one of bearer_token or username/password must be configured")
	}

	return nil
}

type authTransport struct {
	base       http.RoundTripper
	mintClient *http.Client
	endpoint   string
	cfg        Config
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	body, err := readBody(req)
	if err != nil {
		return nil, err
	}

	resp, err := t.roundTrip(req, body, false)
	if err != nil || resp == nil || resp.StatusCode != http.StatusUnauthorized || t.cfg.Username == "" {
		return resp, err
	}

	_, _ = io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()
	tokenCache.Delete(t.cfg.Username)

	return t.roundTrip(req, body, true)
}

func (t *authTransport) roundTrip(req *http.Request, body []byte, forceRefresh bool) (*http.Response, error) {
	cloned := cloneRequest(req, body)
	header, err := t.authorizationHeader(cloned.Context(), forceRefresh)
	if err != nil {
		return nil, err
	}

	tflog.Trace(cloned.Context(), "sg-api request", map[string]any{
		"method": cloned.Method,
		"path":   cloned.URL.EscapedPath(),
	})
	cloned.Header.Set("Authorization", header)

	return t.base.RoundTrip(cloned)
}

func (t *authTransport) authorizationHeader(ctx context.Context, forceRefresh bool) (string, error) {
	switch {
	case t.cfg.BearerToken != "":
		return bearerHeader(t.cfg.BearerToken), nil
	default:
		token, err := t.getOrMint(ctx, forceRefresh)
		if err != nil {
			return "", err
		}
		return bearerHeader(token), nil
	}
}

func (t *authTransport) getOrMint(ctx context.Context, forceRefresh bool) (string, error) {
	if !forceRefresh {
		if cached, ok := tokenCache.Load(t.cfg.Username); ok {
			token := cached.(cachedToken)
			if time.Now().Before(token.expiresAt) {
				return token.value, nil
			}
		}
	}

	token, err := t.mintToken(ctx)
	if err != nil {
		return "", err
	}
	tokenCache.Store(t.cfg.Username, cachedToken{
		value:     token,
		expiresAt: time.Now().Add(tokenTTL),
	})
	return token, nil
}

func (t *authTransport) mintToken(ctx context.Context) (string, error) {
	body, err := json.Marshal(map[string]string{
		"clientId":     t.cfg.Username,
		"clientSecret": t.cfg.Password,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, t.endpoint+"/identity/users/token", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := t.mintClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		_, _ = io.Copy(io.Discard, resp.Body)
		return "", fmt.Errorf("sgclient: identity token mint failed with status %d", resp.StatusCode)
	}

	var decoded struct {
		Value string `json:"value"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		return "", err
	}
	if decoded.Value == "" {
		return "", errors.New("sgclient: identity token response is missing value")
	}

	return decoded.Value, nil
}

func readBody(req *http.Request) ([]byte, error) {
	if req.Body == nil {
		return nil, nil
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	_ = req.Body.Close()
	return body, nil
}

func cloneRequest(req *http.Request, body []byte) *http.Request {
	cloned := req.Clone(req.Context())
	if body != nil {
		cloned.Body = io.NopCloser(bytes.NewReader(body))
		cloned.ContentLength = int64(len(body))
	}
	return cloned
}

func bearerHeader(value string) string {
	if strings.HasPrefix(strings.ToLower(value), "bearer ") {
		return value
	}
	return "Bearer " + value
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}
