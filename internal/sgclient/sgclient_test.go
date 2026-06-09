package sgclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
)

func TestClientAuthorizesBearerTokenMode(t *testing.T) {
	t.Parallel()

	// Given
	server := authAssertServer(t, "Bearer configured-bearer")
	defer server.Close()

	client, err := New(Config{
		Endpoint:    server.URL,
		BearerToken: "configured-bearer",
		HTTPClient:  server.Client(),
	})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// When
	resp, err := client.HTTPClient.Get(server.URL + "/v3/projects")
	if err != nil {
		t.Fatalf("GET error = %v", err)
	}
	defer resp.Body.Close()

	// Then
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", resp.StatusCode, http.StatusOK)
	}
}

func TestClientMintsUsernamePasswordTokenAndCachesIt(t *testing.T) {
	clearTokenCache()

	var mintCalls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/identity/users/token":
			// Given
			mintCalls.Add(1)
			assertJSONBody(t, r, `{"clientId":"user@example.com","clientSecret":"pw"}`)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"value":"minted-token"}`))
		case "/v3/projects":
			// Then
			if got, want := r.Header.Get("Authorization"), "Bearer minted-token"; got != want {
				t.Fatalf("Authorization = %q, want %q", got, want)
			}
			w.WriteHeader(http.StatusOK)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client, err := New(Config{
		Endpoint:   server.URL,
		Username:   "user@example.com",
		Password:   "pw",
		HTTPClient: server.Client(),
	})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// When
	for range 2 {
		resp, err := client.HTTPClient.Get(server.URL + "/v3/projects")
		if err != nil {
			t.Fatalf("GET error = %v", err)
		}
		_ = resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("status = %d, want %d", resp.StatusCode, http.StatusOK)
		}
	}

	// Then
	if got, want := mintCalls.Load(), int32(1); got != want {
		t.Fatalf("mint calls = %d, want %d", got, want)
	}
}

func TestClientRefreshesOnceAndRetriesOnUnauthorized(t *testing.T) {
	clearTokenCache()

	var mintCalls atomic.Int32
	var resourceCalls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/identity/users/token":
			call := mintCalls.Add(1)
			w.Header().Set("Content-Type", "application/json")
			_, _ = fmt.Fprintf(w, `{"value":"minted-token-%d"}`, call)
		case "/v3/projects":
			// Given
			assertTextBody(t, r, "payload")
			call := resourceCalls.Add(1)
			if call == 1 {
				if got, want := r.Header.Get("Authorization"), "Bearer minted-token-1"; got != want {
					t.Fatalf("first Authorization = %q, want %q", got, want)
				}
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Then
			if got, want := r.Header.Get("Authorization"), "Bearer minted-token-2"; got != want {
				t.Fatalf("retry Authorization = %q, want %q", got, want)
			}
			w.WriteHeader(http.StatusOK)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client, err := New(Config{
		Endpoint:   server.URL,
		Username:   "retry@example.com",
		Password:   "pw",
		HTTPClient: server.Client(),
	})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// When
	req, err := http.NewRequest(http.MethodPost, server.URL+"/v3/projects", strings.NewReader("payload"))
	if err != nil {
		t.Fatalf("NewRequest error = %v", err)
	}
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		t.Fatalf("Do error = %v", err)
	}
	defer resp.Body.Close()

	// Then
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if got, want := mintCalls.Load(), int32(2); got != want {
		t.Fatalf("mint calls = %d, want %d", got, want)
	}
	if got, want := resourceCalls.Load(), int32(2); got != want {
		t.Fatalf("resource calls = %d, want %d", got, want)
	}
}

func TestClientPassesRawAuthorizationBearerVerbatim(t *testing.T) {
	t.Parallel()

	// Given a credential that already carries the "Bearer " scheme (Mode C raw passthrough)
	server := authAssertServer(t, "Bearer raw-passthrough")
	defer server.Close()

	client, err := New(Config{
		Endpoint:    server.URL,
		BearerToken: "Bearer raw-passthrough",
		HTTPClient:  server.Client(),
	})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// When the client issues a request
	resp, err := client.HTTPClient.Get(server.URL + "/v3/projects")
	if err != nil {
		t.Fatalf("GET error = %v", err)
	}
	defer resp.Body.Close()

	// Then the credential is forwarded verbatim without a duplicated scheme
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", resp.StatusCode, http.StatusOK)
	}
}

func TestBearerHeaderNormalisesScheme(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name  string
		value string
		want  string
	}{
		// Given a bare credential, When formatting, Then the scheme is prepended
		{name: "bare value", value: "abc", want: "Bearer abc"},
		// Given a value already carrying the canonical scheme, Then it is left verbatim
		{name: "canonical scheme", value: "Bearer abc", want: "Bearer abc"},
		// Given a lower-case scheme prefix, Then it is treated as already-scoped and left verbatim
		{name: "lowercase scheme", value: "bearer abc", want: "bearer abc"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// When
			got := bearerHeader(tc.value)

			// Then
			if got != tc.want {
				t.Fatalf("bearerHeader(%q) = %q, want %q", tc.value, got, tc.want)
			}
		})
	}
}

func TestNewResolvesEndpointFromEnvironment(t *testing.T) {
	// Given only the SCALEGRID_ENDPOINT env var (env-only configuration shape)
	t.Setenv("SCALEGRID_ENDPOINT", "https://env.example.com/")
	t.Setenv("SG_API_BASE_URL", "")

	// When the client is constructed without an explicit endpoint
	client, err := New(Config{BearerToken: "configured-bearer"})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Then the endpoint is read from the environment with the trailing slash trimmed
	if got, want := client.Endpoint, "https://env.example.com"; got != want {
		t.Fatalf("endpoint = %q, want %q", got, want)
	}
}

func TestNewFallsBackToLegacyBaseURLEnvironment(t *testing.T) {
	// Given the legacy SG_API_BASE_URL env var only
	t.Setenv("SCALEGRID_ENDPOINT", "")
	t.Setenv("SG_API_BASE_URL", "https://legacy.example.com")

	// When the client is constructed without an explicit endpoint
	client, err := New(Config{BearerToken: "configured-bearer"})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Then the legacy endpoint env var is honoured
	if got, want := client.Endpoint, "https://legacy.example.com"; got != want {
		t.Fatalf("endpoint = %q, want %q", got, want)
	}
}

func TestNewRejectsInvalidCredentialCombinations(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		cfg  Config
		want string
	}{
		// Given a username without a password, Then construction fails
		{
			name: "username without password",
			cfg:  Config{Endpoint: "https://example.com", Username: "user@example.com"},
			want: "username and password must be configured together",
		},
		// Given no credential at all, Then construction fails
		{
			name: "no credentials",
			cfg:  Config{Endpoint: "https://example.com"},
			want: "exactly one of bearer_token or username/password",
		},
		// Given more than one credential mode, Then construction fails
		{
			name: "multiple credentials",
			cfg:  Config{Endpoint: "https://example.com", BearerToken: "b", Username: "u", Password: "p"},
			want: "exactly one of bearer_token or username/password",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// When
			_, err := New(tc.cfg)

			// Then
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("New() error = %v, want substring %q", err, tc.want)
			}
		})
	}
}

func TestNewRejectsInvalidEndpoints(t *testing.T) {
	cases := []struct {
		name     string
		endpoint string
		want     string
	}{
		// Given an endpoint missing scheme and host, Then construction fails
		{name: "missing scheme and host", endpoint: "not-a-url", want: "endpoint must include scheme and host"},
		// Given an unparseable endpoint, Then construction fails
		{name: "unparseable", endpoint: "http://%zz", want: "invalid endpoint"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// resolveEndpoint consults env fallbacks; clear them so the configured value is used.
			t.Setenv("SCALEGRID_ENDPOINT", "")
			t.Setenv("SG_API_BASE_URL", "")

			// When
			_, err := New(Config{Endpoint: tc.endpoint, BearerToken: "configured-bearer"})

			// Then
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("New() error = %v, want substring %q", err, tc.want)
			}
		})
	}
}

func TestNewRequestBuildsAbsoluteURL(t *testing.T) {
	t.Parallel()

	// Given a client targeting a configured endpoint
	client := &Client{Endpoint: "https://example.com/base"}

	// When a relative request is built
	req, err := client.NewRequest(context.Background(), http.MethodGet, "v3/projects", nil)
	if err != nil {
		t.Fatalf("NewRequest() error = %v", err)
	}

	// Then the path is resolved against the endpoint and the method is preserved
	if got, want := req.URL.String(), "https://example.com/v3/projects"; got != want {
		t.Fatalf("url = %q, want %q", got, want)
	}
	if got, want := req.Method, http.MethodGet; got != want {
		t.Fatalf("method = %q, want %q", got, want)
	}
}

func TestNewRequestRejectsNilClient(t *testing.T) {
	t.Parallel()

	// Given a nil client
	var client *Client

	// When a request is built
	_, err := client.NewRequest(context.Background(), http.MethodGet, "v3/projects", nil)

	// Then construction fails
	if err == nil || !strings.Contains(err.Error(), "nil client") {
		t.Fatalf("NewRequest() error = %v, want nil client error", err)
	}
}

func TestNewRequestRejectsInvalidInputs(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		endpoint string
		path     string
		want     string
	}{
		// Given an unparseable endpoint, Then request construction fails
		{name: "invalid endpoint", endpoint: "http://%zz", path: "v3/projects", want: "invalid endpoint"},
		// Given an unparseable path, Then request construction fails
		{name: "invalid path", endpoint: "https://example.com", path: "%zz", want: "invalid request path"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Given
			client := &Client{Endpoint: tc.endpoint}

			// When
			_, err := client.NewRequest(context.Background(), http.MethodGet, tc.path, nil)

			// Then
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("NewRequest() error = %v, want substring %q", err, tc.want)
			}
		})
	}
}

func TestClientReturnsErrorWhenMintFails(t *testing.T) {
	clearTokenCache()

	// Given an identity endpoint that rejects the mint request
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/identity/users/token" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.NotFound(w, r)
	}))
	defer server.Close()

	client, err := New(Config{
		Endpoint:   server.URL,
		Username:   "mint-fail@example.com",
		Password:   "pw",
		HTTPClient: server.Client(),
	})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// When a request triggers minting
	_, err = client.HTTPClient.Get(server.URL + "/v3/projects")

	// Then the mint failure is surfaced and nothing is cached
	if err == nil || !strings.Contains(err.Error(), "identity token mint failed with status 500") {
		t.Fatalf("GET error = %v, want mint failure", err)
	}
	if _, ok := tokenCache.Load("mint-fail@example.com"); ok {
		t.Fatalf("token unexpectedly cached after mint failure")
	}
}

func TestClientReturnsErrorWhenMintResponseMissingValue(t *testing.T) {
	clearTokenCache()

	// Given an identity endpoint that returns a body without a token value
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/identity/users/token" {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"value":""}`))
			return
		}
		http.NotFound(w, r)
	}))
	defer server.Close()

	client, err := New(Config{
		Endpoint:   server.URL,
		Username:   "empty-value@example.com",
		Password:   "pw",
		HTTPClient: server.Client(),
	})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// When a request triggers minting
	_, err = client.HTTPClient.Get(server.URL + "/v3/projects")

	// Then the empty value is reported as an error
	if err == nil || !strings.Contains(err.Error(), "identity token response is missing value") {
		t.Fatalf("GET error = %v, want missing value error", err)
	}
}

func TestClientReturnsErrorWhenMintResponseMalformed(t *testing.T) {
	clearTokenCache()

	// Given an identity endpoint that returns malformed JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/identity/users/token" {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{not-json`))
			return
		}
		http.NotFound(w, r)
	}))
	defer server.Close()

	client, err := New(Config{
		Endpoint:   server.URL,
		Username:   "malformed@example.com",
		Password:   "pw",
		HTTPClient: server.Client(),
	})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// When a request triggers minting
	_, err = client.HTTPClient.Get(server.URL + "/v3/projects")

	// Then the decode failure is surfaced
	if err == nil {
		t.Fatalf("GET error = nil, want decode failure")
	}
}

func TestClientReturnsErrorWhenRequestBodyUnreadable(t *testing.T) {
	t.Parallel()

	// Given a request whose body cannot be read
	server := authAssertServer(t, "Bearer configured-bearer")
	defer server.Close()

	client, err := New(Config{
		Endpoint:    server.URL,
		BearerToken: "configured-bearer",
		HTTPClient:  server.Client(),
	})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, server.URL+"/v3/projects", io.NopCloser(errReader{}))
	if err != nil {
		t.Fatalf("NewRequest error = %v", err)
	}

	// When the request is sent
	_, err = client.HTTPClient.Do(req)

	// Then the read failure is surfaced
	if err == nil || !strings.Contains(err.Error(), "boom") {
		t.Fatalf("Do error = %v, want body read failure", err)
	}
}

type errReader struct{}

func (errReader) Read([]byte) (int, error) {
	return 0, fmt.Errorf("boom")
}

func authAssertServer(t *testing.T, expectedAuth string) *httptest.Server {
	t.Helper()

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Then
		if got := r.Header.Get("Authorization"); got != expectedAuth {
			t.Fatalf("Authorization = %q, want %q", got, expectedAuth)
		}
		w.WriteHeader(http.StatusOK)
	}))
}

func assertJSONBody(t *testing.T, r *http.Request, expected string) {
	t.Helper()

	got, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("ReadAll() error = %v", err)
	}
	if string(got) != expected {
		t.Fatalf("body = %q, want %q", got, expected)
	}
}

func assertTextBody(t *testing.T, r *http.Request, expected string) {
	t.Helper()

	got, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("ReadAll() error = %v", err)
	}
	if string(got) != expected {
		t.Fatalf("body = %q, want %q", got, expected)
	}
}

func clearTokenCache() {
	tokenCache = sync.Map{}
}
