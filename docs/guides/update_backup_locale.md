---
page_title: "update_backup_locale"
description: |-
  Update Backup Locale
---

# update_backup_locale

See the [`scalegrid_operation` data source documentation](../data-sources/operation#update_backup_locale) for the full reference.

**HTTP:** `PUT /backuplocales`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "update_backup_locale"
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"creationTime\":\"e2e-smoke\",\"labels\":[\"e2e-smoke\"],\"backupLocaleRoles\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"permissions\":[{\"errors\":[{}],\"name\":\"e2e-smoke\",\"category\":\"e2e-smoke\",\"permissionName\":\"e2e-smoke\",\"resourceServerIdentifier\":\"e2e-smoke\",\"valid\":false}],\"modificationAllowed\":false,\"valid\":false}],\"organizationUid\":\"e2e-smoke\",\"cloudProvider\":\"AWS\",\"platformId\":\"e2e-smoke\",\"backupStorages\":[\"STORAGE_CLASS\"],\"objectStorageName\":\"e2e-smoke\",\"valid\":false}"
}
```
