---
page_title: "create_role"
description: |-
  Create Role
---

# create_role

See the [`scalegrid_operation` data source documentation](../data-sources/operation#create_role) for the full reference.

**HTTP:** `POST /identity/organizations/roles`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "create_role"
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"permissions\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"name\":\"e2e-smoke\",\"category\":\"e2e-smoke\",\"permissionName\":\"e2e-smoke\",\"resourceServerIdentifier\":\"e2e-smoke\",\"valid\":false}],\"modificationAllowed\":false,\"valid\":false}"
}
```
