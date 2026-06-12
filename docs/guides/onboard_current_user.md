---
page_title: "onboard_current_user"
description: |-
  Onboard User
---

# onboard_current_user

See the [`scalegrid_operation` resource documentation](../resources/operation#onboard_current_user) for the full reference.

**HTTP:** `POST /identity/users/onboard`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "onboard_current_user"
}
```
