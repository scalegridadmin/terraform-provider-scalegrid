---
page_title: "get_current_organization_user_by_token"
description: |-
  Get Current User
---

# get_current_organization_user_by_token

See the [`scalegrid_operation` resource documentation](../resources/operation#get_current_organization_user_by_token) for the full reference.

**HTTP:** `GET /identity/organizations/users/user`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "get_current_organization_user_by_token"
}
```
