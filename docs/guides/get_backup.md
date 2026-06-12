---
page_title: "get_backup"
description: |-
  Get Backup
---

# get_backup

See the [`scalegrid_operation` resource documentation](../resources/operation#get_backup) for the full reference.

**HTTP:** `GET /backups/{backupUid}`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "get_backup"
  path_params = {
    backupUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
