// Package provider wires the tfplugingen-framework generated
// ScalegridProviderSchema (in scalegrid_provider_gen.go) into a fully
// functional provider.Provider implementation. Hand-written on purpose:
// the framework codegen emits schema helpers only, and a Terraform
// provider needs Metadata / Schema / Configure / Resources / DataSources
// glue to load.
package provider

import (
	"context"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/providervalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	providerschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/scalegridadmin/terraform-provider-scalegrid/internal/datasources"
	"github.com/scalegridadmin/terraform-provider-scalegrid/internal/sgclient"
)

var _ provider.ProviderWithConfigValidators = &scalegridProvider{}

// scalegridProvider satisfies provider.Provider. The factory returned
// by New() is consumed by main.go via providerserver.Serve.
type scalegridProvider struct {
	// version is the build-stamped provider version surfaced through
	// Metadata so `terraform version`/provider diagnostics report the
	// release packaged by the `Build Terraform Provider` pipeline step.
	version string
}

type scalegridProviderModel struct {
	BearerToken types.String `tfsdk:"bearer_token"`
	Username    types.String `tfsdk:"username"`
	Password    types.String `tfsdk:"password"`
	Endpoint    types.String `tfsdk:"endpoint"`
}

// New is the provider factory entry point.
func New(version string) func() provider.Provider {
	return func() provider.Provider { return &scalegridProvider{version: version} }
}

func (p *scalegridProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "scalegrid"
	resp.Version = p.version
}

func (p *scalegridProvider) Schema(ctx context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	providerSchema := ScalegridProviderSchema(ctx)
	if providerSchema.Attributes == nil {
		providerSchema.Attributes = map[string]providerschema.Attribute{}
	}

	providerSchema.Attributes["bearer_token"] = providerschema.StringAttribute{
		Optional:    true,
		Sensitive:   true,
		Description: "Bearer credential used for sg-api requests. May also be set with SCALEGRID_BEARER_TOKEN.",
	}
	providerSchema.Attributes["username"] = providerschema.StringAttribute{
		Optional:    true,
		Sensitive:   true,
		Description: "Username for minting an sg-api bearer token. May also be set with SCALEGRID_USERNAME.",
	}
	providerSchema.Attributes["password"] = providerschema.StringAttribute{
		Optional:    true,
		Sensitive:   true,
		Description: "Password for minting an sg-api bearer token. May also be set with SCALEGRID_PASSWORD.",
	}
	providerSchema.Attributes["endpoint"] = providerschema.StringAttribute{
		Optional:    true,
		Description: "Base URL for sg-api. May also be set with SCALEGRID_ENDPOINT.",
	}

	resp.Schema = providerSchema
}

func (p *scalegridProvider) ConfigValidators(_ context.Context) []provider.ConfigValidator {
	if authEnvConfigured() {
		return nil
	}

	return []provider.ConfigValidator{
		providervalidator.ExactlyOneOf(
			path.MatchRoot("bearer_token"),
			path.MatchRoot("username"),
		),
	}
}

func (p *scalegridProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config scalegridProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, err := sgclient.New(sgclient.Config{
		BearerToken: configuredString(config.BearerToken, "SCALEGRID_BEARER_TOKEN", "bearer_token", &resp.Diagnostics),
		Username:    configuredString(config.Username, "SCALEGRID_USERNAME", "username", &resp.Diagnostics),
		Password:    configuredString(config.Password, "SCALEGRID_PASSWORD", "password", &resp.Diagnostics),
		Endpoint:    configuredString(config.Endpoint, "SCALEGRID_ENDPOINT", "endpoint", &resp.Diagnostics),
	})
	if resp.Diagnostics.HasError() {
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Unable to configure ScaleGrid client", err.Error())
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *scalegridProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}

func (p *scalegridProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		datasources.NewOperationDataSource,
	}
}

func configuredString(value types.String, envName, attrName string, diags *diag.Diagnostics) string {
	if value.IsUnknown() {
		diags.AddAttributeError(
			path.Root(attrName),
			"Unknown provider configuration value",
			"The ScaleGrid provider cannot configure the sg-api client until this value is known.",
		)
		return ""
	}

	if !value.IsNull() && strings.TrimSpace(value.ValueString()) != "" {
		return strings.TrimSpace(value.ValueString())
	}

	return strings.TrimSpace(os.Getenv(envName))
}

func authEnvConfigured() bool {
	return strings.TrimSpace(os.Getenv("SCALEGRID_BEARER_TOKEN")) != "" ||
		strings.TrimSpace(os.Getenv("SCALEGRID_USERNAME")) != "" ||
		strings.TrimSpace(os.Getenv("SCALEGRID_PASSWORD")) != ""
}
