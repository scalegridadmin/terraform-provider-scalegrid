---
page_title: "get_users"
description: |-
  Get Users
---

# get_users

See the [`scalegrid_operation` resource documentation](../resources/operation#get_users) for the full reference.

**HTTP:** `POST /identity/organizations/users/list`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "get_users"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```
