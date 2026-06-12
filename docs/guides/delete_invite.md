---
page_title: "delete_invite"
description: |-
  Delete User Invite
---

# delete_invite

See the [`scalegrid_operation` resource documentation](../resources/operation#delete_invite) for the full reference.

**HTTP:** `DELETE /identity/organizations/users/invite/{inviteUid}`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "delete_invite"
  path_params = {
    inviteUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
