---
page_title: "get_cluster_state_1"
description: |-
  Get Cluster State
---

# get_cluster_state_1

See the [`scalegrid_operation` data source documentation](../data-sources/operation#get_cluster_state_1) for the full reference.

**HTTP:** `GET /clusters/{organizationClusterUid}/state`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "get_cluster_state_1"
  path_params = {
    organizationClusterUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
