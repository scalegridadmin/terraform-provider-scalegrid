package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestMetadataReportsStampedVersion(t *testing.T) {
	// Given a provider built from the New factory with a stamped version
	p := New("2026.05.29-abc1234")()

	// When Metadata is requested
	var resp provider.MetadataResponse
	p.Metadata(context.Background(), provider.MetadataRequest{}, &resp)

	// Then the type name is fixed and the stamped version is surfaced
	if want := "scalegrid"; resp.TypeName != want {
		t.Fatalf("Metadata TypeName = %q, want %q", resp.TypeName, want)
	}
	if want := "2026.05.29-abc1234"; resp.Version != want {
		t.Fatalf("Metadata Version = %q, want %q", resp.Version, want)
	}
}

func TestConfiguredStringPrefersConfigValueOverEnvironment(t *testing.T) {
	// Given a configured attribute value and a populated env fallback
	t.Setenv("SCALEGRID_BEARER_TOKEN", "env-bearer")
	var diags diag.Diagnostics

	// When the value is resolved
	got := configuredString(types.StringValue("  config-bearer  "), "SCALEGRID_BEARER_TOKEN", "bearer_token", &diags)

	// Then the trimmed config value wins and no diagnostics are raised
	if want := "config-bearer"; got != want {
		t.Fatalf("configuredString = %q, want %q", got, want)
	}
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
}

func TestConfiguredStringFallsBackToEnvironmentWhenConfigUnset(t *testing.T) {
	cases := []struct {
		name  string
		value types.String
	}{
		// Given the attribute is null (env-only SCALEGRID_* credential shape)
		{name: "null value", value: types.StringNull()},
		// Given the attribute is an empty/whitespace string
		{name: "blank value", value: types.StringValue("   ")},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Given a populated env fallback and no usable config value
			t.Setenv("SCALEGRID_BEARER_TOKEN", "  env-bearer  ")
			var diags diag.Diagnostics

			// When the value is resolved
			got := configuredString(tc.value, "SCALEGRID_BEARER_TOKEN", "bearer_token", &diags)

			// Then the trimmed env value is used and no diagnostics are raised
			if want := "env-bearer"; got != want {
				t.Fatalf("configuredString = %q, want %q", got, want)
			}
			if diags.HasError() {
				t.Fatalf("unexpected diagnostics: %v", diags)
			}
		})
	}
}

func TestConfiguredStringReturnsEmptyWhenNeitherConfigured(t *testing.T) {
	// Given neither a config value nor an env fallback
	t.Setenv("SCALEGRID_USERNAME", "")
	var diags diag.Diagnostics

	// When the value is resolved
	got := configuredString(types.StringNull(), "SCALEGRID_USERNAME", "username", &diags)

	// Then the result is empty and no diagnostics are raised
	if got != "" {
		t.Fatalf("configuredString = %q, want empty string", got)
	}
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
}

func TestConfiguredStringReportsUnknownValue(t *testing.T) {
	// Given an unknown (not-yet-resolved) attribute value
	var diags diag.Diagnostics

	// When the value is resolved
	got := configuredString(types.StringUnknown(), "SCALEGRID_BEARER_TOKEN", "bearer_token", &diags)

	// Then an attribute error is raised and an empty string is returned
	if got != "" {
		t.Fatalf("configuredString = %q, want empty string", got)
	}
	if !diags.HasError() {
		t.Fatalf("expected an error diagnostic for an unknown value, got none")
	}
}

func TestAuthEnvConfigured(t *testing.T) {
	cases := []struct {
		name string
		env  string
	}{
		// Given each SCALEGRID_* auth env var in turn, Then env auth is detected
		{name: "bearer token", env: "SCALEGRID_BEARER_TOKEN"},
		{name: "username", env: "SCALEGRID_USERNAME"},
		{name: "password", env: "SCALEGRID_PASSWORD"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			clearAuthEnv(t)

			// Given a single auth env var is populated
			t.Setenv(tc.env, "value")

			// When env auth is probed
			// Then it is reported as configured
			if !authEnvConfigured() {
				t.Fatalf("authEnvConfigured() = false, want true with %s set", tc.env)
			}
		})
	}
}

func TestAuthEnvConfiguredFalseWhenNoAuthEnv(t *testing.T) {
	// Given no SCALEGRID_* auth env vars are set
	clearAuthEnv(t)

	// When env auth is probed
	// Then it is reported as not configured
	if authEnvConfigured() {
		t.Fatalf("authEnvConfigured() = true, want false with no auth env set")
	}
}

func TestConfigValidatorsBypassedWhenAuthEnvConfigured(t *testing.T) {
	// Given the env-only SCALEGRID_* credential shape (auth supplied via env)
	clearAuthEnv(t)
	t.Setenv("SCALEGRID_BEARER_TOKEN", "env-bearer")

	// When the provider reports its config validators
	got := (&scalegridProvider{}).ConfigValidators(context.Background())

	// Then the ExactlyOneOf validator is bypassed so env-only auth is accepted
	if len(got) != 0 {
		t.Fatalf("ConfigValidators() returned %d validators, want 0 when auth env is configured", len(got))
	}
}

func TestConfigValidatorsRequireExactlyOneCredentialWithoutEnv(t *testing.T) {
	// Given no auth env vars (credentials must come from the provider block)
	clearAuthEnv(t)

	// When the provider reports its config validators
	got := (&scalegridProvider{}).ConfigValidators(context.Background())

	// Then a single ExactlyOneOf validator guards the credential shapes
	if len(got) != 1 {
		t.Fatalf("ConfigValidators() returned %d validators, want 1 without auth env", len(got))
	}
}

func clearAuthEnv(t *testing.T) {
	t.Helper()

	for _, name := range []string{
		"SCALEGRID_BEARER_TOKEN",
		"SCALEGRID_USERNAME",
		"SCALEGRID_PASSWORD",
	} {
		t.Setenv(name, "")
	}
}
