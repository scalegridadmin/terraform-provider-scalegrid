---
page_title: "get_roles"
description: |-
  Get Roles
---

# get_roles

See the [`scalegrid_operation` data source documentation](../data-sources/operation#get_roles) for the full reference.

**HTTP:** `POST /identity/organizations/roles/list`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "get_roles"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```
