---
page_title: "create_scalegrid_cluster"
description: |-
  Create Cluster
---

# create_scalegrid_cluster

See the [`scalegrid_operation` data source documentation](../data-sources/operation#create_scalegrid_cluster) for the full reference.

**HTTP:** `POST /clusters`

## Terraform Example

```terraform
data "scalegrid_operation" "example" {
  operation_id = "create_scalegrid_cluster"
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"cloudProvider\":\"AWS\",\"organizationClusterName\":\"e2e-smoke\",\"clusterKubernetesNamespace\":\"e2e-smoke\",\"affinitySelector\":[\"e2e-smoke\"],\"antiAffinity\":[\"e2e-smoke\"],\"nodeSelector\":[\"e2e-smoke\"],\"storage\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"type\":\"STORAGE_CLASS\",\"name\":\"e2e-smoke\",\"allowVolumeExpansion\":false,\"valid\":false},\"organizationClusterDescription\":\"e2e-smoke\",\"nodeName\":\"e2e-smoke\",\"clusterUid\":\"e2e-smoke\",\"labels\":[\"e2e-smoke\"],\"organizationClusterRoles\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"permissions\":[{\"errors\":[{}],\"name\":\"e2e-smoke\",\"category\":\"e2e-smoke\",\"permissionName\":\"e2e-smoke\",\"resourceServerIdentifier\":\"e2e-smoke\",\"valid\":false}],\"modificationAllowed\":false,\"valid\":false}],\"os\":\"LINUX\",\"architecture\":\"AMD64\",\"helmDeploymentCommand\":\"e2e-smoke\",\"region\":\"e2e-smoke\",\"ingressControllers\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"from\":0,\"to\":0,\"name\":\"e2e-smoke\",\"ingressType\":\"PUBLICLY_EXPOSED\",\"valid\":false}],\"platformID\":\"e2e-smoke\",\"tolerations\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operation\":\"Exists\",\"effect\":\"NoSchedule\",\"valid\":false}],\"externalCidr\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"cidr\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"valid\":false},\"podCidr\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"cidr\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"valid\":false},\"tunnelCidr\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"cidr\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"valid\":false},\"kubernetesApiUrl\":\"e2e-smoke\",\"valid\":false}"
}
```
