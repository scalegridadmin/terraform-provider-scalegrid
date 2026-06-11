---
page_title: "delete_scalegrid_service_cluster"
description: |-
  Delete Service
---

# delete_scalegrid_service_cluster

See the [`scalegrid_operation` data source documentation](../data-sources/operation#delete_scalegrid_service_cluster) for the full reference.

**HTTP:** `DELETE /scalegridserviceclusters/{scalegridServiceClusterUid}`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "delete_scalegrid_service_cluster"
  path_params = {
    scalegridServiceClusterUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
