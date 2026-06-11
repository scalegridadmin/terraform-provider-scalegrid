---
page_title: "get_backup_locale_list"
description: |-
  Get Backup Locale
---

# get_backup_locale_list

See the [`scalegrid_operation` data source documentation](../data-sources/operation#get_backup_locale_list) for the full reference.

**HTTP:** `POST /backuplocales/list`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "get_backup_locale_list"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```
