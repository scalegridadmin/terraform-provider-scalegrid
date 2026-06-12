---
page_title: "get_latest_cluster_version"
description: |-
  Get Latest Cluster Version
---

# get_latest_cluster_version

See the [`scalegrid_operation` resource documentation](../resources/operation#get_latest_cluster_version) for the full reference.

**HTTP:** `GET /clusters/clusterversion`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "get_latest_cluster_version"
}
```
