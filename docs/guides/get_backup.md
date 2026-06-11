---
page_title: "get_backup"
description: |-
  Get Backup
---

# get_backup

See the [`scalegrid_operation` data source documentation](../data-sources/operation#get_backup) for the full reference.

**HTTP:** `GET /backups/{backupUid}`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "get_backup"
  path_params = {
    backupUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
