---
page_title: "delete_user_from_current_organization"
description: |-
  Delete User
---

# delete_user_from_current_organization

See the [`scalegrid_operation` resource documentation](../resources/operation#delete_user_from_current_organization) for the full reference.

**HTTP:** `DELETE /identity/organizations/users/{userId}`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "delete_user_from_current_organization"
  path_params = {
    userId = "00000000-0000-0000-0000-000000000001"
  }
}
```
