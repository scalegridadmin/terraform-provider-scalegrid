---
page_title: "get_cluster_state"
description: |-
  Get Service State
---

# get_cluster_state

See the [`scalegrid_operation` data source documentation](../data-sources/operation#get_cluster_state) for the full reference.

**HTTP:** `GET /scalegridserviceclusters/{scalegridServiceClusterUid}/state`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "get_cluster_state"
  path_params = {
    scalegridServiceClusterUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
