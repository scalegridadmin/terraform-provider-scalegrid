---
page_title: "get_certificate"
description: |-
  Get Service Certificate
---

# get_certificate

See the [`scalegrid_operation` resource documentation](../resources/operation#get_certificate) for the full reference.

**HTTP:** `GET /scalegridserviceclusters/{scalegridServiceClusterUid}/certificate`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "get_certificate"
  path_params = {
    scalegridServiceClusterUid = "00000000-0000-0000-0000-000000000001"
  }
}
```
