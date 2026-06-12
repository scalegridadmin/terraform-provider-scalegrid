---
page_title: "delete_backup"
description: |-
  Delete Backup
---

# delete_backup

See the [`scalegrid_operation` resource documentation](../resources/operation#delete_backup) for the full reference.

**HTTP:** `DELETE /backups/{backupConfigurationsUid}/backups/{backupUid}`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "delete_backup"
  path_params = {
    backupConfigurationsUid = "00000000-0000-0000-0000-000000000001"
    backupUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
