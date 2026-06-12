---
page_title: "get_scalegrid_service_types"
description: |-
  Get Service Types
---

# get_scalegrid_service_types

See the [`scalegrid_operation` resource documentation](../resources/operation#get_scalegrid_service_types) for the full reference.

**HTTP:** `POST /scalegridservices/scalegridservicetypes/list`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "get_scalegrid_service_types"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```
