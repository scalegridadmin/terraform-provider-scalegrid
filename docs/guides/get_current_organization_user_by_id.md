---
page_title: "get_current_organization_user_by_id"
description: |-
  Get User
---

# get_current_organization_user_by_id

See the [`scalegrid_operation` resource documentation](../resources/operation#get_current_organization_user_by_id) for the full reference.

**HTTP:** `GET /identity/organizations/users/{userId}`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "get_current_organization_user_by_id"
  path_params = {
    userId = "00000000-0000-0000-0000-000000000001"
  }
}
```
