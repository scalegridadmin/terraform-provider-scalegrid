---
page_title: "invite_user_for_current_organization"
description: |-
  Invite User
---

# invite_user_for_current_organization

See the [`scalegrid_operation` resource documentation](../resources/operation#invite_user_for_current_organization) for the full reference.

**HTTP:** `POST /identity/organizations/users/invite/{email}`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "invite_user_for_current_organization"
  path_params = {
    email = "e2e-smoke@example.com"
  }
}
```
