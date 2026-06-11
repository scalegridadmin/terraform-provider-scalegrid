---
page_title: "get_cluster_metadata"
description: |-
  Get Metadata
---

# get_cluster_metadata

See the [`scalegrid_operation` data source documentation](../data-sources/operation#get_cluster_metadata) for the full reference.

**HTTP:** `GET /clusters/{organizationClusterUid}/metadata`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "get_cluster_metadata"
  path_params = {
    organizationClusterUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
