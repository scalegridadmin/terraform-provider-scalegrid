---
page_title: "delete_organization"
description: |-
  Delete Organization
---

# delete_organization

See the [`scalegrid_operation` data source documentation](../data-sources/operation#delete_organization) for the full reference.

**HTTP:** `DELETE /organization`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "delete_organization"
}
```
