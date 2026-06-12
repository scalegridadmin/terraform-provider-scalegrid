---
page_title: "get_recommended_settings"
description: |-
  Get Recommended Settings
---

# get_recommended_settings

See the [`scalegrid_operation` resource documentation](../resources/operation#get_recommended_settings) for the full reference.

**HTTP:** `POST /scalegridserviceclusters/recommendation`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "get_recommended_settings"
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"scalegridServiceVersionUid\":\"e2e-smoke\",\"cpu\":0,\"memory\":0,\"storage\":0,\"numberOfReplicas\":0,\"recommendations\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"recommended\":\"e2e-smoke\",\"valid\":false}],\"valid\":false}"
}
```
