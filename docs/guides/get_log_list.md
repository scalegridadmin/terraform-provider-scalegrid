---
page_title: "get_log_list"
description: |-
  Get Instance Logs
---

# get_log_list

See the [`scalegrid_operation` data source documentation](../data-sources/operation#get_log_list) for the full reference.

**HTTP:** `POST /scalegridserviceclusters/{scalegridServiceClusterUid}/instances/{instanceUid}/logs/list`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "get_log_list"
  path_params = {
    scalegridServiceClusterUid = "00000000-0000-0000-0000-000000000001"
    instanceUid = "00000000-0000-0000-0000-000000000001"
  }
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```
