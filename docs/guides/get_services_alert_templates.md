---
page_title: "get_services_alert_templates"
description: |-
  Get Alert Templates
---

# get_services_alert_templates

See the [`scalegrid_operation` resource documentation](../resources/operation#get_services_alert_templates) for the full reference.

**HTTP:** `POST /alerts/templates/list`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "get_services_alert_templates"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```
