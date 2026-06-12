---
page_title: "create_restore"
description: |-
  Restore Backup
---

# create_restore

See the [`scalegrid_operation` resource documentation](../resources/operation#create_restore) for the full reference.

**HTTP:** `POST /backups/restore`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "create_restore"
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"targetInstances\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"backupUid\":\"e2e-smoke\",\"configurationDetails\":\"e2e-smoke\",\"startTime\":\"e2e-smoke\",\"endTime\":\"e2e-smoke\",\"status\":\"WIP\",\"statusDetails\":\"e2e-smoke\",\"targetInstanceUid\":\"e2e-smoke\",\"restoreClusterUid\":\"e2e-smoke\",\"targetInstanceName\":\"e2e-smoke\",\"valid\":false}],\"valid\":false}"
}
```
