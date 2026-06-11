---
page_title: "get_backup_list"
description: |-
  Get Backups
---

# get_backup_list

See the [`scalegrid_operation` data source documentation](../data-sources/operation#get_backup_list) for the full reference.

**HTTP:** `POST /backups/list`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "get_backup_list"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```
