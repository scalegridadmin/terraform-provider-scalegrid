---
page_title: "get_ingresses"
description: |-
  Get Ingresses
---

# get_ingresses

See the [`scalegrid_operation` data source documentation](../data-sources/operation#get_ingresses) for the full reference.

**HTTP:** `POST /scalegridservices/ingresstypes/list`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "get_ingresses"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```
