---
page_title: "create_backup"
description: |-
  Create Backup
---

# create_backup

See the [`scalegrid_operation` resource documentation](../resources/operation#create_backup) for the full reference.

**HTTP:** `POST /backupconfigurations/{backupConfigurationUid}/backups/run`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "create_backup"
  path_params = {
    backupConfigurationUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
