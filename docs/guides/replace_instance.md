---
page_title: "replace_instance"
description: |-
  Replace Instance
---

# replace_instance

See the [`scalegrid_operation` data source documentation](../data-sources/operation#replace_instance) for the full reference.

**HTTP:** `POST /scalegridserviceclusters/instances/{instanceUid}/replace`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "replace_instance"
  path_params = {
    instanceUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
