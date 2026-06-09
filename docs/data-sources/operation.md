---
page_title: "scalegrid_operation Data Source"
subcategory: ""
description: |-
  Calls a ScaleGrid API operation by operation ID.
---

# scalegrid_operation (Data Source)

Invokes a ScaleGrid API operation. Operation IDs and request bodies match the ScaleGrid OpenAPI specification.

## Example Usage

```terraform
data "scalegrid_operation" "pg_versions" {
  operation_id = "get_scalegrid_service_versions"
  body_json = jsonencode({
    pageIndex = 0
    pageSize  = 10
  })
}
```

## Schema

### Required

- `operation_id` (String) OpenAPI operation identifier.

### Optional

- `path_params` (Map of String) Path parameter values for templated URLs.
- `body_json` (String) JSON request body.

### Read-Only

- `id` (String) Synthetic identifier for Terraform state.
- `status_code` (Number) HTTP status code from the API response.
- `response` (String) JSON response body from the API.
