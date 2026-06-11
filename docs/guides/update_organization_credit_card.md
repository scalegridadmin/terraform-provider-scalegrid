---
page_title: "update_organization_credit_card"
description: |-
  Update Credit Card
---

# update_organization_credit_card

See the [`scalegrid_operation` data source documentation](../data-sources/operation#update_organization_credit_card) for the full reference.

**HTTP:** `PATCH /organization/creditCard`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "update_organization_credit_card"
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"creditCardToken\":\"e2e-smoke\",\"valid\":false}"
}
```
