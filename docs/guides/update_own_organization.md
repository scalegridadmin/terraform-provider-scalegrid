---
page_title: "update_own_organization"
description: |-
  Update Organization
---

# update_own_organization

See the [`scalegrid_operation` resource documentation](../resources/operation#update_own_organization) for the full reference.

**HTTP:** `PUT /organization`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "update_own_organization"
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"email\":\"e2e-smoke\",\"status\":\"TRIAL\",\"trialExpirationDate\":\"e2e-smoke\",\"invoicingType\":\"MANUAL\",\"paymentMethod\":\"BANK_TRANSFER\",\"defaultInviteRole\":\"e2e-smoke\",\"statusHistory\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"newStatus\":\"TRIAL\",\"oldStatus\":\"TRIAL\",\"timestamp\":\"e2e-smoke\",\"valid\":false}],\"creationTime\":\"e2e-smoke\",\"creditCardBrand\":\"e2e-smoke\",\"creditCardLast4\":\"e2e-smoke\",\"maxNumberOfRoles\":0,\"valid\":false}"
}
```
