---
page_title: "get_event_list"
description: |-
  Get Instance Events
---

# get_event_list

See the [`scalegrid_operation` resource documentation](../resources/operation#get_event_list) for the full reference.

**HTTP:** `POST /scalegridserviceclusters/{scalegridServiceClusterUid}/instances/{instanceUid}/events/list`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "get_event_list"
  path_params = {
    scalegridServiceClusterUid = "00000000-0000-0000-0000-000000000001"
    instanceUid = "00000000-0000-0000-0000-000000000001"
  }
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```
