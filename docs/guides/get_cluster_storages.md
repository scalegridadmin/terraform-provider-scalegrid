---
page_title: "get_cluster_storages"
description: |-
  Get Storages
---

# get_cluster_storages

See the [`scalegrid_operation` data source documentation](../data-sources/operation#get_cluster_storages) for the full reference.

**HTTP:** `GET /clusters/{organizationClusterUid}/storages`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "get_cluster_storages"
  path_params = {
    organizationClusterUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
