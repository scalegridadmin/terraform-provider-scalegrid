---
page_title: "get_backup_locale_state"
description: |-
  Get Backup Locale State
---

# get_backup_locale_state

See the [`scalegrid_operation` data source documentation](../data-sources/operation#get_backup_locale_state) for the full reference.

**HTTP:** `GET /backuplocales/{backupLocaleUid}/state`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "get_backup_locale_state"
  path_params = {
    backupLocaleUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
