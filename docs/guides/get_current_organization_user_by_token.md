---
page_title: "get_current_organization_user_by_token"
description: |-
  Get Current User
---

# get_current_organization_user_by_token

See the [`scalegrid_operation` data source documentation](../data-sources/operation#get_current_organization_user_by_token) for the full reference.

**HTTP:** `GET /identity/organizations/users/user`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "get_current_organization_user_by_token"
}
```
