---
page_title: "delete_backup_locale"
description: |-
  Delete Backup Locale
---

# delete_backup_locale

See the [`scalegrid_operation` data source documentation](../data-sources/operation#delete_backup_locale) for the full reference.

**HTTP:** `DELETE /backuplocales/{backupLocaleUid}`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "delete_backup_locale"
  path_params = {
    backupLocaleUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
