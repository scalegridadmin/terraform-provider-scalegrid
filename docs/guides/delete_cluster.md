---
page_title: "delete_cluster"
description: |-
  Delete Cluster
---

# delete_cluster

See the [`scalegrid_operation` resource documentation](../resources/operation#delete_cluster) for the full reference.

**HTTP:** `DELETE /clusters/{organizationClusterUid}`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "delete_cluster"
  path_params = {
    organizationClusterUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
