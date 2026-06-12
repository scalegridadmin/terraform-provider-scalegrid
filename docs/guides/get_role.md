---
page_title: "get_role"
description: |-
  Get Role
---

# get_role

See the [`scalegrid_operation` resource documentation](../resources/operation#get_role) for the full reference.

**HTTP:** `GET /identity/organizations/roles/{roleId}`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "get_role"
  path_params = {
    roleId = "00000000-0000-0000-0000-000000000001"
  }
}
```
