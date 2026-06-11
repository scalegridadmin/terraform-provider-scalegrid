---
page_title: "delete_cluster"
description: |-
  Delete Cluster
---

# delete_cluster

See the [`scalegrid_operation` data source documentation](../data-sources/operation#delete_cluster) for the full reference.

**HTTP:** `DELETE /clusters/{organizationClusterUid}`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "delete_cluster"
  path_params = {
    organizationClusterUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
