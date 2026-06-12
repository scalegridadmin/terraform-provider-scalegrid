---
page_title: "clone_scalegrid_service_cluster"
description: |-
  Clone Service
---

# clone_scalegrid_service_cluster

See the [`scalegrid_operation` resource documentation](../resources/operation#clone_scalegrid_service_cluster) for the full reference.

**HTTP:** `POST /scalegridserviceclusters/clone`

## Terraform Example

```terraform
resource "scalegrid_operation" "example" {
  operation_id = "clone_scalegrid_service_cluster"
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"scalegridServiceClusterUid\":\"e2e-smoke\",\"targetScalegridServiceClusterUid\":\"e2e-smoke\",\"backupUid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"valid\":false}"
}
```
