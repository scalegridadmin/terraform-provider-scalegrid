---
page_title: "get_organization_cluster"
description: |-
  Get Cluster
---

# get_organization_cluster

See the [`scalegrid_operation` resource documentation](../resources/operation#get_organization_cluster) for the full reference.

**HTTP:** `GET /clusters/organizationClusters/{organizationClusterUid}`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "get_organization_cluster"
  path_params = {
    organizationClusterUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
