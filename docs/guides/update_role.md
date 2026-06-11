---
page_title: "update_role"
description: |-
  Update Role
---

# update_role

See the [`scalegrid_operation` data source documentation](../data-sources/operation#update_role) for the full reference.

**HTTP:** `PUT /identity/organizations/roles/{roleId}`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "update_role"
  path_params = {
    roleId = "00000000-0000-0000-0000-000000000001"
  }
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"permissions\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"name\":\"e2e-smoke\",\"category\":\"e2e-smoke\",\"permissionName\":\"e2e-smoke\",\"resourceServerIdentifier\":\"e2e-smoke\",\"valid\":false}],\"modificationAllowed\":false,\"valid\":false}"
}
```
