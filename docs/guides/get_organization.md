---
page_title: "get_organization"
description: |-
  Get Organization
---

# get_organization

See the [`scalegrid_operation` data source documentation](../data-sources/operation#get_organization) for the full reference.

**HTTP:** `GET /organization`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "get_organization"
}
```
