package datasources

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/scalegridadmin/terraform-provider-scalegrid/internal/sgclient"
)

var pathParamPattern = regexp.MustCompile(`\{([^}]+)\}`)

var _ datasource.DataSource = &operationDataSource{}
var _ datasource.DataSourceWithConfigure = &operationDataSource{}

type operationDataSource struct {
	client *sgclient.Client
}

type operationDataSourceModel struct {
	OperationID types.String `tfsdk:"operation_id"`
	PathParams  types.Map    `tfsdk:"path_params"`
	BodyJSON    types.String `tfsdk:"body_json"`
	ID          types.String `tfsdk:"id"`
	StatusCode  types.Int64  `tfsdk:"status_code"`
	Response    types.String `tfsdk:"response"`
}

func NewOperationDataSource() datasource.DataSource {
	return &operationDataSource{}
}

func (d *operationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_operation"
}

func (d *operationDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"operation_id": schema.StringAttribute{
				Required:    true,
				Description: "OpenAPI operation key from the auto-generated catalog (for example, get_scalegrid_service_versions).",
			},
			"path_params": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Path parameter values substituted into the OpenAPI path template.",
			},
			"body_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON request body for POST, PUT, and PATCH operations.",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Internal ID for the operation invocation.",
			},
			"status_code": schema.Int64Attribute{
				Computed:    true,
				Description: "HTTP status code returned by sg-api.",
			},
			"response": schema.StringAttribute{
				Computed:    true,
				Description: "Raw response body returned by sg-api.",
			},
		},
	}
}

func (d *operationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sgclient.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected provider data type",
			fmt.Sprintf("Expected *sgclient.Client but got %T", req.ProviderData),
		)
		return
	}

	d.client = client
}

func (d *operationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError(
			"Unconfigured ScaleGrid client",
			"Provider configuration is missing; data source cannot call sg-api.",
		)
		return
	}

	var state operationDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	operationID := strings.TrimSpace(state.OperationID.ValueString())
	spec, ok := operationCatalog[operationID]
	if !ok {
		resp.Diagnostics.AddAttributeError(
			path.Root("operation_id"),
			"Unknown operation_id",
			fmt.Sprintf("operation_id %q is not present in the OpenAPI operation catalog", operationID),
		)
		return
	}

	pathParams, diags := pathParamsFromState(ctx, state.PathParams)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resolvedPath, diags := renderPathTemplate(spec.Path, pathParams)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var body io.Reader
	if methodAllowsBody(spec.Method) {
		bodyJSON := strings.TrimSpace(state.BodyJSON.ValueString())
		if bodyJSON != "" {
			if !json.Valid([]byte(bodyJSON)) {
				resp.Diagnostics.AddAttributeError(
					path.Root("body_json"),
					"Invalid JSON in body_json",
					"body_json must contain valid JSON when set.",
				)
				return
			}
			body = strings.NewReader(bodyJSON)
		}
	}

	httpReq, err := d.client.NewRequest(ctx, spec.Method, resolvedPath, body)
	if err != nil {
		resp.Diagnostics.AddError("Unable to build sg-api request", err.Error())
		return
	}
	if body != nil {
		httpReq.Header.Set("Content-Type", "application/json")
	}
	httpReq.Header.Set("Accept", "application/json")

	httpResp, err := d.client.HTTPClient.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Unable to call sg-api", err.Error())
		return
	}
	defer httpResp.Body.Close()

	bodyBytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		resp.Diagnostics.AddError("Unable to read sg-api response", err.Error())
		return
	}

	state.ID = types.StringValue(operationID)
	state.StatusCode = types.Int64Value(int64(httpResp.StatusCode))
	state.Response = types.StringValue(string(bodyBytes))

	if httpResp.StatusCode < http.StatusOK || httpResp.StatusCode >= http.StatusMultipleChoices {
		resp.Diagnostics.AddError(
			"sg-api request failed",
			fmt.Sprintf(
				"%s %s returned status %d: %s",
				spec.Method,
				resolvedPath,
				httpResp.StatusCode,
				truncateForDiagnostic(string(bodyBytes)),
			),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func pathParamsFromState(ctx context.Context, value types.Map) (map[string]string, diag.Diagnostics) {
	var diags diag.Diagnostics
	if value.IsNull() || value.IsUnknown() {
		return map[string]string{}, diags
	}

	raw := map[string]string{}
	diags.Append(value.ElementsAs(ctx, &raw, false)...)
	return raw, diags
}

func renderPathTemplate(template string, params map[string]string) (string, diag.Diagnostics) {
	var diags diag.Diagnostics
	missing := []string{}

	resolved := pathParamPattern.ReplaceAllStringFunc(template, func(match string) string {
		name := strings.TrimSuffix(strings.TrimPrefix(match, "{"), "}")
		value, ok := params[name]
		if !ok || strings.TrimSpace(value) == "" {
			missing = append(missing, name)
			return match
		}
		return value
	})

	if len(missing) > 0 {
		diags.AddAttributeError(
			path.Root("path_params"),
			"Missing required path parameters",
			fmt.Sprintf("Provide path_params values for: %s", strings.Join(missing, ", ")),
		)
	}

	return resolved, diags
}

func methodAllowsBody(method string) bool {
	switch strings.ToUpper(method) {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		return true
	default:
		return false
	}
}

func truncateForDiagnostic(value string) string {
	const maxLen = 600
	if len(value) <= maxLen {
		return value
	}
	return value[:maxLen] + "...(truncated)"
}
