---
page_title: "failover_instance"
description: |-
  Failover Instance
---

# failover_instance

See the [`scalegrid_operation` data source documentation](../data-sources/operation#failover_instance) for the full reference.

**HTTP:** `POST /scalegridserviceclusters/instances/{instanceUid}/failover`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "failover_instance"
  path_params = {
    instanceUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
