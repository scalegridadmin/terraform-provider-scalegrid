---
page_title: "get_dashboards"
description: |-
  Get Dashboards
---

# get_dashboards

See the [`scalegrid_operation` data source documentation](../data-sources/operation#get_dashboards) for the full reference.

**HTTP:** `POST /dashboards/scalegridservices/dashboards/list`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "get_dashboards"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```
