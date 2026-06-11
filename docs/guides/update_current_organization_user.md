---
page_title: "update_current_organization_user"
description: |-
  Update User
---

# update_current_organization_user

See the [`scalegrid_operation` data source documentation](../data-sources/operation#update_current_organization_user) for the full reference.

**HTTP:** `PUT /identity/organizations/users/{userId}`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "update_current_organization_user"
  path_params = {
    userId = "00000000-0000-0000-0000-000000000001"
  }
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"email\":\"e2e-smoke\",\"password\":\"e2e-smoke\",\"organizationId\":\"e2e-smoke\",\"connection\":\"e2e-smoke\",\"roles\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"permissions\":[{\"errors\":[{}],\"name\":\"e2e-smoke\",\"category\":\"e2e-smoke\",\"permissionName\":\"e2e-smoke\",\"resourceServerIdentifier\":\"e2e-smoke\",\"valid\":false}],\"modificationAllowed\":false,\"valid\":false}],\"creationDate\":\"e2e-smoke\",\"mfaEnabled\":false,\"mfaKey\":\"e2e-smoke\",\"valid\":false}"
}
```
