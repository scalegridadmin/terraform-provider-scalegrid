---
page_title: "onboard_current_user"
description: |-
  Onboard User
---

# onboard_current_user

See the [`scalegrid_operation` data source documentation](../data-sources/operation#onboard_current_user) for the full reference.

**HTTP:** `POST /identity/users/onboard`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "onboard_current_user"
}
```
