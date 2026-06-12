---
page_title: "get_backup_locale"
description: |-
  Get Backup Locale
---

# get_backup_locale

See the [`scalegrid_operation` resource documentation](../resources/operation#get_backup_locale) for the full reference.

**HTTP:** `GET /backuplocales/{backupLocaleUid}`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "get_backup_locale"
  path_params = {
    backupLocaleUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
