---
page_title: "replace_instance"
description: |-
  Replace Instance
---

# replace_instance

See the [`scalegrid_operation` resource documentation](../resources/operation#replace_instance) for the full reference.

**HTTP:** `POST /scalegridserviceclusters/instances/{instanceUid}/replace`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "replace_instance"
  path_params = {
    instanceUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
