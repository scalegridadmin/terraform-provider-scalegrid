---
page_title: "create_backup_locale"
description: |-
  Create Backup Locale
---

# create_backup_locale

See the [`scalegrid_operation` resource documentation](../resources/operation#create_backup_locale) for the full reference.

**HTTP:** `POST /backuplocales`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "create_backup_locale"
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"creationTime\":\"e2e-smoke\",\"labels\":[\"e2e-smoke\"],\"backupLocaleRoles\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"permissions\":[{\"errors\":[{}],\"name\":\"e2e-smoke\",\"category\":\"e2e-smoke\",\"permissionName\":\"e2e-smoke\",\"resourceServerIdentifier\":\"e2e-smoke\",\"valid\":false}],\"modificationAllowed\":false,\"valid\":false}],\"organizationUid\":\"e2e-smoke\",\"cloudProvider\":\"AWS\",\"platformId\":\"e2e-smoke\",\"backupStorages\":[\"STORAGE_CLASS\"],\"objectStorageName\":\"e2e-smoke\",\"valid\":false}"
}
```
