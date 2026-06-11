---
page_title: "get_latest_cluster_version"
description: |-
  Get Latest Cluster Version
---

# get_latest_cluster_version

See the [`scalegrid_operation` data source documentation](../data-sources/operation#get_latest_cluster_version) for the full reference.

**HTTP:** `GET /clusters/clusterversion`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "get_latest_cluster_version"
}
```
