---
page_title: "delete_role"
description: |-
  Delete Role
---

# delete_role

See the [`scalegrid_operation` resource documentation](../resources/operation#delete_role) for the full reference.

**HTTP:** `DELETE /identity/organizations/roles/{roleId}`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "delete_role"
  path_params = {
    roleId = "00000000-0000-0000-0000-000000000001"
  }
}
```
