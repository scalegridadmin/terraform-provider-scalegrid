---
page_title: "get_instance_state"
description: |-
  Get Instance State
---

# get_instance_state

See the [`scalegrid_operation` data source documentation](../data-sources/operation#get_instance_state) for the full reference.

**HTTP:** `GET /scalegridserviceclusters/{scalegridServiceClusterUid}/instances/{instanceUid}/state`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "get_instance_state"
  path_params = {
    scalegridServiceClusterUid = "00000000-0000-0000-0000-000000000001"
    instanceUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
