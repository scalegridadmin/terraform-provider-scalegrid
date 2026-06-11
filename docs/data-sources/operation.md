---
page_title: "scalegrid_operation Data Source"
description: |-
  Calls any ScaleGrid API operation by operation_id.
---

# scalegrid_operation (Data Source)

Calls ScaleGrid API endpoints from Terraform. Set `operation_id` to the operation you want to invoke, then optionally supply `path_params` and `body_json`.

```terraform
data "scalegrid_operation" "example" {
  operation_id = "<operation_id>"
}
```

## Schema

### Required

- `operation_id` (String) Operation key listed below (for example, `get_organization`).

### Optional

- `path_params` (Map of String) Path parameter values for templated URLs.
- `body_json` (String) JSON request body for POST, PUT, and PATCH operations.
- `allow_non_success_status` (Boolean) When `true`, non-2xx HTTP responses do not fail the read. Defaults to `false`.

### Read-Only

- `id` (String) Synthetic identifier for Terraform state.
- `status_code` (Number) HTTP status code from the API response.
- `response` (String) JSON response body from the API.

## Operations

The table below lists every supported `operation_id`. Each entry links to its Terraform example further down this page.

| operation_id | Method | Path | Summary |
| --- | --- | --- | --- |
| [clone_scalegrid_service_cluster](#clone_scalegrid_service_cluster) | `POST` | `/scalegridserviceclusters/clone` | Clone Service |
| [create_backup](#create_backup) | `POST` | `/backupconfigurations/{backupConfigurationUid}/backups/run` | Create Backup |
| [create_backup_locale](#create_backup_locale) | `POST` | `/backuplocales` | Create Backup Locale |
| [create_restore](#create_restore) | `POST` | `/backups/restore` | Restore Backup |
| [create_role](#create_role) | `POST` | `/identity/organizations/roles` | Create Role |
| [create_scalegrid_cluster](#create_scalegrid_cluster) | `POST` | `/clusters` | Create Cluster |
| [create_scalegrid_service_cluster](#create_scalegrid_service_cluster) | `POST` | `/scalegridserviceclusters` | Create Service |
| [delete_backup](#delete_backup) | `DELETE` | `/backups/{backupConfigurationsUid}/backups/{backupUid}` | Delete Backup |
| [delete_backup_locale](#delete_backup_locale) | `DELETE` | `/backuplocales/{backupLocaleUid}` | Delete Backup Locale |
| [delete_cluster](#delete_cluster) | `DELETE` | `/clusters/{organizationClusterUid}` | Delete Cluster |
| [delete_invite](#delete_invite) | `DELETE` | `/identity/organizations/users/invite/{inviteUid}` | Delete User Invite |
| [delete_organization](#delete_organization) | `DELETE` | `/organization` | Delete Organization |
| [delete_role](#delete_role) | `DELETE` | `/identity/organizations/roles/{roleId}` | Delete Role |
| [delete_scalegrid_service_cluster](#delete_scalegrid_service_cluster) | `DELETE` | `/scalegridserviceclusters/{scalegridServiceClusterUid}` | Delete Service |
| [delete_user_from_current_organization](#delete_user_from_current_organization) | `DELETE` | `/identity/organizations/users/{userId}` | Delete User |
| [failover_instance](#failover_instance) | `POST` | `/scalegridserviceclusters/instances/{instanceUid}/failover` | Failover Instance |
| [get_audits](#get_audits) | `POST` | `/audits/list` | Get Audit Logs |
| [get_backup](#get_backup) | `GET` | `/backups/{backupUid}` | Get Backup |
| [get_backup_list](#get_backup_list) | `POST` | `/backups/list` | Get Backups |
| [get_backup_locale](#get_backup_locale) | `GET` | `/backuplocales/{backupLocaleUid}` | Get Backup Locale |
| [get_backup_locale_list](#get_backup_locale_list) | `POST` | `/backuplocales/list` | Get Backup Locale |
| [get_backup_locale_state](#get_backup_locale_state) | `GET` | `/backuplocales/{backupLocaleUid}/state` | Get Backup Locale State |
| [get_backup_types](#get_backup_types) | `POST` | `/backuptypes/list` | Get Backup Types |
| [get_certificate](#get_certificate) | `GET` | `/scalegridserviceclusters/{scalegridServiceClusterUid}/certificate` | Get Service Certificate |
| [get_cloud_providers](#get_cloud_providers) | `POST` | `/clusters/cloudproviders/list` | Get Cloud Providers |
| [get_cluster_list](#get_cluster_list) | `POST` | `/clusters/list` | Get Dedicated Hosting Clusters |
| [get_cluster_metadata](#get_cluster_metadata) | `GET` | `/clusters/{organizationClusterUid}/metadata` | Get Metadata |
| [get_cluster_state](#get_cluster_state) | `GET` | `/scalegridserviceclusters/{scalegridServiceClusterUid}/state` | Get Service State |
| [get_cluster_state_1](#get_cluster_state_1) | `GET` | `/clusters/{organizationClusterUid}/state` | Get Cluster State |
| [get_cluster_storages](#get_cluster_storages) | `GET` | `/clusters/{organizationClusterUid}/storages` | Get Storages |
| [get_current_organization_user_by_id](#get_current_organization_user_by_id) | `GET` | `/identity/organizations/users/{userId}` | Get User |
| [get_current_organization_user_by_token](#get_current_organization_user_by_token) | `GET` | `/identity/organizations/users/user` | Get Current User |
| [get_dashboards](#get_dashboards) | `POST` | `/dashboards/scalegridservices/dashboards/list` | Get Dashboards |
| [get_event_list](#get_event_list) | `POST` | `/scalegridserviceclusters/{scalegridServiceClusterUid}/instances/{instanceUid}/events/list` | Get Instance Events |
| [get_ingresses](#get_ingresses) | `POST` | `/scalegridservices/ingresstypes/list` | Get Ingresses |
| [get_instance_state](#get_instance_state) | `GET` | `/scalegridserviceclusters/{scalegridServiceClusterUid}/instances/{instanceUid}/state` | Get Instance State |
| [get_invites](#get_invites) | `POST` | `/identity/organizations/users/invite/list` | Get User Invites |
| [get_latest_cluster_version](#get_latest_cluster_version) | `GET` | `/clusters/clusterversion` | Get Latest Cluster Version |
| [get_log_list](#get_log_list) | `POST` | `/scalegridserviceclusters/{scalegridServiceClusterUid}/instances/{instanceUid}/logs/list` | Get Instance Logs |
| [get_organization](#get_organization) | `GET` | `/organization` | Get Organization |
| [get_organization_cluster](#get_organization_cluster) | `GET` | `/clusters/organizationClusters/{organizationClusterUid}` | Get Cluster |
| [get_organization_clusters](#get_organization_clusters) | `POST` | `/clusters/clusters/list` | Get Clusters |
| [get_permissions](#get_permissions) | `POST` | `/identity/roles/permissions/list` | Get Permissions |
| [get_recommended_settings](#get_recommended_settings) | `POST` | `/scalegridserviceclusters/recommendation` | Get Recommended Settings |
| [get_restore_list](#get_restore_list) | `POST` | `/backups/restores/list` | Get Restores |
| [get_role](#get_role) | `GET` | `/identity/organizations/roles/{roleId}` | Get Role |
| [get_roles](#get_roles) | `POST` | `/identity/organizations/roles/list` | Get Roles |
| [get_scalegrid_service_cluster](#get_scalegrid_service_cluster) | `GET` | `/scalegridserviceclusters/{scalegridServiceClusterUid}` | Get Service |
| [get_scalegrid_service_clusters](#get_scalegrid_service_clusters) | `POST` | `/scalegridserviceclusters/list` | Get Services |
| [get_scalegrid_service_types](#get_scalegrid_service_types) | `POST` | `/scalegridservices/scalegridservicetypes/list` | Get Service Types |
| [get_scalegrid_service_versions](#get_scalegrid_service_versions) | `POST` | `/scalegridservices/scalegridserviceversions/list` | Get Service Versions |
| [get_services_alert_templates](#get_services_alert_templates) | `POST` | `/alerts/templates/list` | Get Alert Templates |
| [get_users](#get_users) | `POST` | `/identity/organizations/users/list` | Get Users |
| [invite_user_for_current_organization](#invite_user_for_current_organization) | `POST` | `/identity/organizations/users/invite/{email}` | Invite User |
| [onboard_current_user](#onboard_current_user) | `POST` | `/identity/users/onboard` | Onboard User |
| [replace_instance](#replace_instance) | `POST` | `/scalegridserviceclusters/instances/{instanceUid}/replace` | Replace Instance |
| [update_backup_locale](#update_backup_locale) | `PUT` | `/backuplocales` | Update Backup Locale |
| [update_cluster](#update_cluster) | `PUT` | `/clusters/{organizationClusterUid}` | Update Clusters |
| [update_current_organization_user](#update_current_organization_user) | `PUT` | `/identity/organizations/users/{userId}` | Update User |
| [update_organization_credit_card](#update_organization_credit_card) | `PATCH` | `/organization/creditCard` | Update Credit Card |
| [update_own_organization](#update_own_organization) | `PUT` | `/organization` | Update Organization |
| [update_role](#update_role) | `PUT` | `/identity/organizations/roles/{roleId}` | Update Role |
| [update_scalegrid_service_cluster](#update_scalegrid_service_cluster) | `PUT` | `/scalegridserviceclusters/{scalegridServiceClusterUid}` | Update Service |

### clone_scalegrid_service_cluster

**HTTP:** `POST /scalegridserviceclusters/clone`  
**Category:** Scalegrid Services

Clones a service based on an existing backup.

**Request body example:**

```json
{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"scalegridServiceClusterUid":"e2e-smoke","targetScalegridServiceClusterUid":"e2e-smoke","backupUid":"e2e-smoke","name":"e2e-smoke","valid":false}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "clone_scalegrid_service_cluster" {
  operation_id = "clone_scalegrid_service_cluster"
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"scalegridServiceClusterUid\":\"e2e-smoke\",\"targetScalegridServiceClusterUid\":\"e2e-smoke\",\"backupUid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"valid\":false}"
}
```

### create_backup

**HTTP:** `POST /backupconfigurations/{backupConfigurationUid}/backups/run`  
**Category:** Backups

Creates a backup based on a specific backup configuration on demand.

**Path parameters:**
- `backupConfigurationUid`

**Terraform example:**

```terraform
data "scalegrid_operation" "create_backup" {
  operation_id = "create_backup"
  path_params = {
    backupConfigurationUid = "00000000-0000-0000-0000-000000000001"
  }
}
```

### create_backup_locale

**HTTP:** `POST /backuplocales`  
**Category:** Backup Locales

Creates a new backup locale.

**Request body example:**

```json
{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","creationTime":"e2e-smoke","labels":["e2e-smoke"],"backupLocaleRoles":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","permissions":[{"errors":[{}],"name":"e2e-smoke","category":"e2e-smoke","permissionName":"e2e-smoke","resourceServerIdentifier":"e2e-smoke","valid":false}],"modificationAllowed":false,"valid":false}],"organizationUid":"e2e-smoke","cloudProvider":"AWS","platformId":"e2e-smoke","backupStorages":["STORAGE_CLASS"],"objectStorageName":"e2e-smoke","valid":false}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "create_backup_locale" {
  operation_id = "create_backup_locale"
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"creationTime\":\"e2e-smoke\",\"labels\":[\"e2e-smoke\"],\"backupLocaleRoles\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"permissions\":[{\"errors\":[{}],\"name\":\"e2e-smoke\",\"category\":\"e2e-smoke\",\"permissionName\":\"e2e-smoke\",\"resourceServerIdentifier\":\"e2e-smoke\",\"valid\":false}],\"modificationAllowed\":false,\"valid\":false}],\"organizationUid\":\"e2e-smoke\",\"cloudProvider\":\"AWS\",\"platformId\":\"e2e-smoke\",\"backupStorages\":[\"STORAGE_CLASS\"],\"objectStorageName\":\"e2e-smoke\",\"valid\":false}"
}
```

### create_restore

**HTTP:** `POST /backups/restore`  
**Category:** Backups

Restores a specific backup.

**Request body example:**

```json
{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","targetInstances":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","backupUid":"e2e-smoke","configurationDetails":"e2e-smoke","startTime":"e2e-smoke","endTime":"e2e-smoke","status":"WIP","statusDetails":"e2e-smoke","targetInstanceUid":"e2e-smoke","restoreClusterUid":"e2e-smoke","targetInstanceName":"e2e-smoke","valid":false}],"valid":false}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "create_restore" {
  operation_id = "create_restore"
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"targetInstances\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"backupUid\":\"e2e-smoke\",\"configurationDetails\":\"e2e-smoke\",\"startTime\":\"e2e-smoke\",\"endTime\":\"e2e-smoke\",\"status\":\"WIP\",\"statusDetails\":\"e2e-smoke\",\"targetInstanceUid\":\"e2e-smoke\",\"restoreClusterUid\":\"e2e-smoke\",\"targetInstanceName\":\"e2e-smoke\",\"valid\":false}],\"valid\":false}"
}
```

### create_role

**HTTP:** `POST /identity/organizations/roles`  
**Category:** Roles

Creates a role.

**Request body example:**

```json
{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","permissions":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"name":"e2e-smoke","category":"e2e-smoke","permissionName":"e2e-smoke","resourceServerIdentifier":"e2e-smoke","valid":false}],"modificationAllowed":false,"valid":false}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "create_role" {
  operation_id = "create_role"
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"permissions\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"name\":\"e2e-smoke\",\"category\":\"e2e-smoke\",\"permissionName\":\"e2e-smoke\",\"resourceServerIdentifier\":\"e2e-smoke\",\"valid\":false}],\"modificationAllowed\":false,\"valid\":false}"
}
```

### create_scalegrid_cluster

**HTTP:** `POST /clusters`  
**Category:** Clusters

Creates a new cluster which represents the connection between Scalegrid* and the infrastructure where your managed services are provisioned and operated.

**Request body example:**

```json
{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","cloudProvider":"AWS","organizationClusterName":"e2e-smoke","clusterKubernetesNamespace":"e2e-smoke","affinitySelector":["e2e-smoke"],"antiAffinity":["e2e-smoke"],"nodeSelector":["e2e-smoke"],"storage":{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"type":"STORAGE_CLASS","name":"e2e-smoke","allowVolumeExpansion":false,"valid":false},"organizationClusterDescription":"e2e-smoke","nodeName":"e2e-smoke","clusterUid":"e2e-smoke","labels":["e2e-smoke"],"organizationClusterRoles":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","permissions":[{"errors":[{}],"name":"e2e-smoke","category":"e2e-smoke","permissionName":"e2e-smoke","resourceServerIdentifier":"e2e-smoke","valid":false}],"modificationAllowed":false,"valid":false}],"os":"LINUX","architecture":"AMD64","helmDeploymentCommand":"e2e-smoke","region":"e2e-smoke","ingressControllers":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"from":0,"to":0,"name":"e2e-smoke","ingressType":"PUBLICLY_EXPOSED","valid":false}],"platformID":"e2e-smoke","tolerations":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operation":"Exists","effect":"NoSchedule","valid":false}],"externalCidr":{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"cidr":"e2e-smoke","name":"e2e-smoke","valid":false},"podCidr":{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"cidr":"e2e-smoke","name":"e2e-smoke","valid":false},"tunnelCidr":{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"cidr":"e2e-smoke","name":"e2e-smoke","valid":false},"kubernetesApiUrl":"e2e-smoke","valid":false}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "create_scalegrid_cluster" {
  operation_id = "create_scalegrid_cluster"
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"cloudProvider\":\"AWS\",\"organizationClusterName\":\"e2e-smoke\",\"clusterKubernetesNamespace\":\"e2e-smoke\",\"affinitySelector\":[\"e2e-smoke\"],\"antiAffinity\":[\"e2e-smoke\"],\"nodeSelector\":[\"e2e-smoke\"],\"storage\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"type\":\"STORAGE_CLASS\",\"name\":\"e2e-smoke\",\"allowVolumeExpansion\":false,\"valid\":false},\"organizationClusterDescription\":\"e2e-smoke\",\"nodeName\":\"e2e-smoke\",\"clusterUid\":\"e2e-smoke\",\"labels\":[\"e2e-smoke\"],\"organizationClusterRoles\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"permissions\":[{\"errors\":[{}],\"name\":\"e2e-smoke\",\"category\":\"e2e-smoke\",\"permissionName\":\"e2e-smoke\",\"resourceServerIdentifier\":\"e2e-smoke\",\"valid\":false}],\"modificationAllowed\":false,\"valid\":false}],\"os\":\"LINUX\",\"architecture\":\"AMD64\",\"helmDeploymentCommand\":\"e2e-smoke\",\"region\":\"e2e-smoke\",\"ingressControllers\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"from\":0,\"to\":0,\"name\":\"e2e-smoke\",\"ingressType\":\"PUBLICLY_EXPOSED\",\"valid\":false}],\"platformID\":\"e2e-smoke\",\"tolerations\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operation\":\"Exists\",\"effect\":\"NoSchedule\",\"valid\":false}],\"externalCidr\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"cidr\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"valid\":false},\"podCidr\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"cidr\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"valid\":false},\"tunnelCidr\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"cidr\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"valid\":false},\"kubernetesApiUrl\":\"e2e-smoke\",\"valid\":false}"
}
```

### create_scalegrid_service_cluster

**HTTP:** `POST /scalegridserviceclusters`  
**Category:** Scalegrid Services

Creates a new Scalegrid managed service.

**Request body example:**

```json
{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","creationTime":"e2e-smoke","hasTelemetry":false,"ingress":"PUBLICLY_EXPOSED","instances":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","organizationClusterUid":"e2e-smoke","scalegridServiceClusterUid":"e2e-smoke","name":"e2e-smoke","kubernetesNamespace":"e2e-smoke","affinitySelector":["e2e-smoke"],"antiAffinitySelector":["e2e-smoke"],"nodeSelector":["e2e-smoke"],"nodeName":"e2e-smoke","autoScalingType":"NONE","autoScalingCpuPercentage":0,"autoScalingMemoryPercentage":0,"storageType":"STORAGE_CLASS","storageName":"e2e-smoke","inboundCIDRs":[{"errors":[{}],"cidr":"e2e-smoke","name":"e2e-smoke","valid":false}],"failoverPriority":0,"additionalSettings":"e2e-smoke","resources":[{"name":"e2e-smoke","resource":"e2e-smoke","certificateExpirationDate":"e2e-smoke"}],"creationTime":"e2e-smoke","cpuSizeMRequest":0,"memorySizeMbRequest":0,"storageSizeMb":0,"cpuSizeMMaxLimit":0,"cpuSizeMMinLimit":0,"memorySizeMbMaxLimit":0,"memorySizeMbMinLimit":0,"alerts":[{"errors":[{}],"uid":"e2e-smoke","organizationClusterUid":"e2e-smoke","scalegridServiceClusterUid":"e2e-smoke","scalegridServiceClusterName":"e2e-smoke","scalegridServiceUid":"e2e-smoke","instanceUid":"e2e-smoke","instanceName":"e2e-smoke","organizationUid":"e2e-smoke","serviceAlertUid":"e2e-smoke","sendResolved":false,"emailsToNotify":["e2e-smoke"],"parameters":[{}],"alertType":"POD","valid":false}],"labels":["e2e-smoke"],"ruleBasedVPAs":[{"errors":[{}],"uid":"e2e-smoke","cpuSizeMRequest":0,"memorySizeMbRequest":0,"cpuSizeMMaxLimit":0,"cpuSizeMMinLimit":0,"memorySizeMbMaxLimit":0,"memorySizeMbMinLimit":0,"cronJob":"e2e-smoke","autoScalingType":"NONE","valid":false}],"architecture":"AMD64","os":"LINUX","tolerations":[{"errors":[{}],"key":"e2e-smoke","value":"e2e-smoke","operation":"Exists","effect":"NoSchedule","valid":false}],"valid":false}],"scalegridServiceVersion":{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","scalegridServiceEnum":"RabbitMQ","versionName":"e2e-smoke","preRequiredScalegridVersionsForUpgrade":"e2e-smoke","availableForNewDeployments":false,"description":"e2e-smoke","allowedAutoScalingTypes":["NONE"],"replicationTypes":[{"errors":[{"key":"e2e-smoke","error":{}}],"uid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","numberOfReplicas":0,"valid":false}],"backupInstallInstructions":"e2e-smoke","clusterType":"ACTIVE_ACTIVE","architectures":["AMD64"],"oss":["LINUX"],"settings":{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"settings":[{"errors":[{}],"uid":"e2e-smoke","name":"e2e-smoke","value":"e2e-smoke","valueType":"NUMBER","options":"e2e-smoke","enabled":false,"description":"e2e-smoke","parentSettingUid":"e2e-smoke","valid":false}],"extensions":[{"errors":[{}],"uid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","enabled":false,"settings":[{}],"requiredExtensionUids":["e2e-smoke"],"valid":false}],"serviceType":"RabbitMQ","valid":false},"instanceSettings":{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"settings":[{"errors":[{}],"uid":"e2e-smoke","name":"e2e-smoke","value":"e2e-smoke","valueType":"NUMBER","options":"e2e-smoke","enabled":false,"description":"e2e-smoke","parentSettingUid":"e2e-smoke","valid":false}],"extensions":[{"errors":[{}],"uid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","enabled":false,"settings":[{}],"requiredExtensionUids":["e2e-smoke"],"valid":false}],"serviceType":"RabbitMQ","valid":false},"compatibleScalegridVersionsBackups":"e2e-smoke","valid":false},"replicationTypeUid":"e2e-smoke","additionalSettings":"e2e-smoke","resources":[{"name":"e2e-smoke","resource":"e2e-smoke","certificateExpirationDate":"e2e-smoke"}],"autoPatch":false,"hasLatestPatch":false,"alerts":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","organizationClusterUid":"e2e-smoke","scalegridServiceClusterUid":"e2e-smoke","scalegridServiceClusterName":"e2e-smoke","scalegridServiceUid":"e2e-smoke","instanceUid":"e2e-smoke","instanceName":"e2e-smoke","organizationUid":"e2e-smoke","serviceAlertUid":"e2e-smoke","sendResolved":false,"emailsToNotify":["e2e-smoke"],"parameters":[{"errors":[{}],"key":"e2e-smoke","value":"e2e-smoke","isCustomerExposed":false,"unit":"e2e-smoke","minValue":0,"maxValue":0,"valid":false}],"alertType":"POD","valid":false}],"serviceClusterRoles":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","roleUid":"e2e-smoke","valid":false}],"autoPatchCronJob":"e2e-smoke","ruleBasedVPAs":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","cpuSizeMRequest":0,"memorySizeMbRequest":0,"cpuSizeMMaxLimit":0,"cpuSizeMMinLimit":0,"memorySizeMbMaxLimit":0,"memorySizeMbMinLimit":0,"cronJob":"e2e-smoke","autoScalingType":"NONE","valid":false}],"backupConfigurations":{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","scalegridServiceClusterUid":"e2e-smoke","backupConfigurations":[{"errors":[{"key":"e2e-smoke","error":{}}],"uid":"e2e-smoke","backupLocaleUid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","backupCronJob":"e2e-smoke","timeToLiveDays":0,"maxSizeInBackupsMb":0,"backupType":{"errors":[{}],"uid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","scalegridServiceVersionList":["00000000-0000-0000-0000-000000000001"],"cloudList":["AWS"],"restoreOptions":"SAME_PLATFORM","backupMethod":"DISC_SNAPSHOT","backupType":"PITR","restoreType":"STANDALONE","backupStorageTypes":["STORAGE_CLASS"],"scalegridServiceEnum":"RabbitMQ","backupOperatorDeployment":"TARGET_INSTANCE","isTargetOptionRequiredForBackup":false,"supportsMultiRegion":false,"valid":false},"maxNumberOfBackups":0,"targetName":"e2e-smoke","compressed":false,"scalegridServiceClusterUid":"e2e-smoke","valid":false}],"timeToLiveDays":0,"maxSizeInBackupsMb":0,"maxNumberOfBackups":0,"valid":false},"cpuSizeMRequest":0,"memorySizeMbRequest":0,"cpuSizeMMaxLimit":0,"cpuSizeMMinLimit":0,"memorySizeMbMaxLimit":0,"memorySizeMbMinLimit":0,"autoScalingType":"NONE","autoScalingCpuPercentage":0,"autoScalingMemoryPercentage":0,"valid":false}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "create_scalegrid_service_cluster" {
  operation_id = "create_scalegrid_service_cluster"
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"creationTime\":\"e2e-smoke\",\"hasTelemetry\":false,\"ingress\":\"PUBLICLY_EXPOSED\",\"instances\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"organizationClusterUid\":\"e2e-smoke\",\"scalegridServiceClusterUid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"kubernetesNamespace\":\"e2e-smoke\",\"affinitySelector\":[\"e2e-smoke\"],\"antiAffinitySelector\":[\"e2e-smoke\"],\"nodeSelector\":[\"e2e-smoke\"],\"nodeName\":\"e2e-smoke\",\"autoScalingType\":\"NONE\",\"autoScalingCpuPercentage\":0,\"autoScalingMemoryPercentage\":0,\"storageType\":\"STORAGE_CLASS\",\"storageName\":\"e2e-smoke\",\"inboundCIDRs\":[{\"errors\":[{}],\"cidr\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"valid\":false}],\"failoverPriority\":0,\"additionalSettings\":\"e2e-smoke\",\"resources\":[{\"name\":\"e2e-smoke\",\"resource\":\"e2e-smoke\",\"certificateExpirationDate\":\"e2e-smoke\"}],\"creationTime\":\"e2e-smoke\",\"cpuSizeMRequest\":0,\"memorySizeMbRequest\":0,\"storageSizeMb\":0,\"cpuSizeMMaxLimit\":0,\"cpuSizeMMinLimit\":0,\"memorySizeMbMaxLimit\":0,\"memorySizeMbMinLimit\":0,\"alerts\":[{\"errors\":[{}],\"uid\":\"e2e-smoke\",\"organizationClusterUid\":\"e2e-smoke\",\"scalegridServiceClusterUid\":\"e2e-smoke\",\"scalegridServiceClusterName\":\"e2e-smoke\",\"scalegridServiceUid\":\"e2e-smoke\",\"instanceUid\":\"e2e-smoke\",\"instanceName\":\"e2e-smoke\",\"organizationUid\":\"e2e-smoke\",\"serviceAlertUid\":\"e2e-smoke\",\"sendResolved\":false,\"emailsToNotify\":[\"e2e-smoke\"],\"parameters\":[{}],\"alertType\":\"POD\",\"valid\":false}],\"labels\":[\"e2e-smoke\"],\"ruleBasedVPAs\":[{\"errors\":[{}],\"uid\":\"e2e-smoke\",\"cpuSizeMRequest\":0,\"memorySizeMbRequest\":0,\"cpuSizeMMaxLimit\":0,\"cpuSizeMMinLimit\":0,\"memorySizeMbMaxLimit\":0,\"memorySizeMbMinLimit\":0,\"cronJob\":\"e2e-smoke\",\"autoScalingType\":\"NONE\",\"valid\":false}],\"architecture\":\"AMD64\",\"os\":\"LINUX\",\"tolerations\":[{\"errors\":[{}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operation\":\"Exists\",\"effect\":\"NoSchedule\",\"valid\":false}],\"valid\":false}],\"scalegridServiceVersion\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"scalegridServiceEnum\":\"RabbitMQ\",\"versionName\":\"e2e-smoke\",\"preRequiredScalegridVersionsForUpgrade\":\"e2e-smoke\",\"availableForNewDeployments\":false,\"description\":\"e2e-smoke\",\"allowedAutoScalingTypes\":[\"NONE\"],\"replicationTypes\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"numberOfReplicas\":0,\"valid\":false}],\"backupInstallInstructions\":\"e2e-smoke\",\"clusterType\":\"ACTIVE_ACTIVE\",\"architectures\":[\"AMD64\"],\"oss\":[\"LINUX\"],\"settings\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"settings\":[{\"errors\":[{}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"valueType\":\"NUMBER\",\"options\":\"e2e-smoke\",\"enabled\":false,\"description\":\"e2e-smoke\",\"parentSettingUid\":\"e2e-smoke\",\"valid\":false}],\"extensions\":[{\"errors\":[{}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"enabled\":false,\"settings\":[{}],\"requiredExtensionUids\":[\"e2e-smoke\"],\"valid\":false}],\"serviceType\":\"RabbitMQ\",\"valid\":false},\"instanceSettings\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"settings\":[{\"errors\":[{}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"valueType\":\"NUMBER\",\"options\":\"e2e-smoke\",\"enabled\":false,\"description\":\"e2e-smoke\",\"parentSettingUid\":\"e2e-smoke\",\"valid\":false}],\"extensions\":[{\"errors\":[{}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"enabled\":false,\"settings\":[{}],\"requiredExtensionUids\":[\"e2e-smoke\"],\"valid\":false}],\"serviceType\":\"RabbitMQ\",\"valid\":false},\"compatibleScalegridVersionsBackups\":\"e2e-smoke\",\"valid\":false},\"replicationTypeUid\":\"e2e-smoke\",\"additionalSettings\":\"e2e-smoke\",\"resources\":[{\"name\":\"e2e-smoke\",\"resource\":\"e2e-smoke\",\"certificateExpirationDate\":\"e2e-smoke\"}],\"autoPatch\":false,\"hasLatestPatch\":false,\"alerts\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"organizationClusterUid\":\"e2e-smoke\",\"scalegridServiceClusterUid\":\"e2e-smoke\",\"scalegridServiceClusterName\":\"e2e-smoke\",\"scalegridServiceUid\":\"e2e-smoke\",\"instanceUid\":\"e2e-smoke\",\"instanceName\":\"e2e-smoke\",\"organizationUid\":\"e2e-smoke\",\"serviceAlertUid\":\"e2e-smoke\",\"sendResolved\":false,\"emailsToNotify\":[\"e2e-smoke\"],\"parameters\":[{\"errors\":[{}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"isCustomerExposed\":false,\"unit\":\"e2e-smoke\",\"minValue\":0,\"maxValue\":0,\"valid\":false}],\"alertType\":\"POD\",\"valid\":false}],\"serviceClusterRoles\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"roleUid\":\"e2e-smoke\",\"valid\":false}],\"autoPatchCronJob\":\"e2e-smoke\",\"ruleBasedVPAs\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"cpuSizeMRequest\":0,\"memorySizeMbRequest\":0,\"cpuSizeMMaxLimit\":0,\"cpuSizeMMinLimit\":0,\"memorySizeMbMaxLimit\":0,\"memorySizeMbMinLimit\":0,\"cronJob\":\"e2e-smoke\",\"autoScalingType\":\"NONE\",\"valid\":false}],\"backupConfigurations\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"scalegridServiceClusterUid\":\"e2e-smoke\",\"backupConfigurations\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{}}],\"uid\":\"e2e-smoke\",\"backupLocaleUid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"backupCronJob\":\"e2e-smoke\",\"timeToLiveDays\":0,\"maxSizeInBackupsMb\":0,\"backupType\":{\"errors\":[{}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"scalegridServiceVersionList\":[\"00000000-0000-0000-0000-000000000001\"],\"cloudList\":[\"AWS\"],\"restoreOptions\":\"SAME_PLATFORM\",\"backupMethod\":\"DISC_SNAPSHOT\",\"backupType\":\"PITR\",\"restoreType\":\"STANDALONE\",\"backupStorageTypes\":[\"STORAGE_CLASS\"],\"scalegridServiceEnum\":\"RabbitMQ\",\"backupOperatorDeployment\":\"TARGET_INSTANCE\",\"isTargetOptionRequiredForBackup\":false,\"supportsMultiRegion\":false,\"valid\":false},\"maxNumberOfBackups\":0,\"targetName\":\"e2e-smoke\",\"compressed\":false,\"scalegridServiceClusterUid\":\"e2e-smoke\",\"valid\":false}],\"timeToLiveDays\":0,\"maxSizeInBackupsMb\":0,\"maxNumberOfBackups\":0,\"valid\":false},\"cpuSizeMRequest\":0,\"memorySizeMbRequest\":0,\"cpuSizeMMaxLimit\":0,\"cpuSizeMMinLimit\":0,\"memorySizeMbMaxLimit\":0,\"memorySizeMbMinLimit\":0,\"autoScalingType\":\"NONE\",\"autoScalingCpuPercentage\":0,\"autoScalingMemoryPercentage\":0,\"valid\":false}"
}
```

### delete_backup

**HTTP:** `DELETE /backups/{backupConfigurationsUid}/backups/{backupUid}`  
**Category:** Backups

Deletes a specific backup.

**Path parameters:**
- `backupConfigurationsUid`
- `backupUid`

**Terraform example:**

```terraform
data "scalegrid_operation" "delete_backup" {
  operation_id = "delete_backup"
  path_params = {
    backupConfigurationsUid = "00000000-0000-0000-0000-000000000001"
    backupUid = "00000000-0000-0000-0000-000000000001"
  }
}
```

### delete_backup_locale

**HTTP:** `DELETE /backuplocales/{backupLocaleUid}`  
**Category:** Backup Locales

Deletes an existing backup locale.

**Path parameters:**
- `backupLocaleUid`

**Terraform example:**

```terraform
data "scalegrid_operation" "delete_backup_locale" {
  operation_id = "delete_backup_locale"
  path_params = {
    backupLocaleUid = "00000000-0000-0000-0000-000000000001"
  }
}
```

### delete_cluster

**HTTP:** `DELETE /clusters/{organizationClusterUid}`  
**Category:** Clusters

Deletes an existing cluster.

**Path parameters:**
- `organizationClusterUid`

**Terraform example:**

```terraform
data "scalegrid_operation" "delete_cluster" {
  operation_id = "delete_cluster"
  path_params = {
    organizationClusterUid = "00000000-0000-0000-0000-000000000001"
  }
}
```

### delete_invite

**HTTP:** `DELETE /identity/organizations/users/invite/{inviteUid}`  
**Category:** Users

Deletes an existing user invite.

**Path parameters:**
- `inviteUid`

**Terraform example:**

```terraform
data "scalegrid_operation" "delete_invite" {
  operation_id = "delete_invite"
  path_params = {
    inviteUid = "00000000-0000-0000-0000-000000000001"
  }
}
```

### delete_organization

**HTTP:** `DELETE /organization`  
**Category:** Organization

Deletes an existing organization.

**Terraform example:**

```terraform
data "scalegrid_operation" "delete_organization" {
  operation_id = "delete_organization"
}
```

### delete_role

**HTTP:** `DELETE /identity/organizations/roles/{roleId}`  
**Category:** Roles

Deletes an existing role.

**Path parameters:**
- `roleId`

**Terraform example:**

```terraform
data "scalegrid_operation" "delete_role" {
  operation_id = "delete_role"
  path_params = {
    roleId = "00000000-0000-0000-0000-000000000001"
  }
}
```

### delete_scalegrid_service_cluster

**HTTP:** `DELETE /scalegridserviceclusters/{scalegridServiceClusterUid}`  
**Category:** Scalegrid Services

Deletes an existing Scalegrid managed service.

**Path parameters:**
- `scalegridServiceClusterUid`

**Terraform example:**

```terraform
data "scalegrid_operation" "delete_scalegrid_service_cluster" {
  operation_id = "delete_scalegrid_service_cluster"
  path_params = {
    scalegridServiceClusterUid = "00000000-0000-0000-0000-000000000001"
  }
}
```

### delete_user_from_current_organization

**HTTP:** `DELETE /identity/organizations/users/{userId}`  
**Category:** Users

Deletes an existing user.

**Path parameters:**
- `userId`

**Terraform example:**

```terraform
data "scalegrid_operation" "delete_user_from_current_organization" {
  operation_id = "delete_user_from_current_organization"
  path_params = {
    userId = "00000000-0000-0000-0000-000000000001"
  }
}
```

### failover_instance

**HTTP:** `POST /scalegridserviceclusters/instances/{instanceUid}/failover`  
**Category:** Scalegrid Services

Forces instance to failover.

**Path parameters:**
- `instanceUid`

**Terraform example:**

```terraform
data "scalegrid_operation" "failover_instance" {
  operation_id = "failover_instance"
  path_params = {
    instanceUid = "00000000-0000-0000-0000-000000000001"
  }
}
```

### get_audits

**HTTP:** `POST /audits/list`  
**Category:** Audit

Returns a list of audit logs.

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_audits" {
  operation_id = "get_audits"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### get_backup

**HTTP:** `GET /backups/{backupUid}`  
**Category:** Backups

Returns an existing backup.

**Path parameters:**
- `backupUid`

**Terraform example:**

```terraform
data "scalegrid_operation" "get_backup" {
  operation_id = "get_backup"
  path_params = {
    backupUid = "00000000-0000-0000-0000-000000000001"
  }
}
```

### get_backup_list

**HTTP:** `POST /backups/list`  
**Category:** Backups

Returns a list of existing backups.

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_backup_list" {
  operation_id = "get_backup_list"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### get_backup_locale

**HTTP:** `GET /backuplocales/{backupLocaleUid}`  
**Category:** Backup Locales

Returns a specific backup locale.

**Path parameters:**
- `backupLocaleUid`

**Terraform example:**

```terraform
data "scalegrid_operation" "get_backup_locale" {
  operation_id = "get_backup_locale"
  path_params = {
    backupLocaleUid = "00000000-0000-0000-0000-000000000001"
  }
}
```

### get_backup_locale_list

**HTTP:** `POST /backuplocales/list`  
**Category:** Backup Locales

Returns a list of existing backup locales.

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_backup_locale_list" {
  operation_id = "get_backup_locale_list"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### get_backup_locale_state

**HTTP:** `GET /backuplocales/{backupLocaleUid}/state`  
**Category:** Backup Locales

Returns the state of an existing backup locale.

**Path parameters:**
- `backupLocaleUid`

**Terraform example:**

```terraform
data "scalegrid_operation" "get_backup_locale_state" {
  operation_id = "get_backup_locale_state"
  path_params = {
    backupLocaleUid = "00000000-0000-0000-0000-000000000001"
  }
}
```

### get_backup_types

**HTTP:** `POST /backuptypes/list`  
**Category:** Backups

Returns a list of backup types for a specific service version.

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_backup_types" {
  operation_id = "get_backup_types"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### get_certificate

**HTTP:** `GET /scalegridserviceclusters/{scalegridServiceClusterUid}/certificate`  
**Category:** Scalegrid Services

Returns the TLS certificate and expiration date for an existing Scalegrid managed service.

**Path parameters:**
- `scalegridServiceClusterUid`

**Terraform example:**

```terraform
data "scalegrid_operation" "get_certificate" {
  operation_id = "get_certificate"
  path_params = {
    scalegridServiceClusterUid = "00000000-0000-0000-0000-000000000001"
  }
}
```

### get_cloud_providers

**HTTP:** `POST /clusters/cloudproviders/list`  
**Category:** Clusters

Returns a list of cloud providers.

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_cloud_providers" {
  operation_id = "get_cloud_providers"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### get_cluster_list

**HTTP:** `POST /clusters/list`  
**Category:** Clusters

Returns a list of available dedicated hosting clusters.

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_cluster_list" {
  operation_id = "get_cluster_list"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### get_cluster_metadata

**HTTP:** `GET /clusters/{organizationClusterUid}/metadata`  
**Category:** Clusters

Returns the available metadata of a specific cluster.

**Path parameters:**
- `organizationClusterUid`

**Terraform example:**

```terraform
data "scalegrid_operation" "get_cluster_metadata" {
  operation_id = "get_cluster_metadata"
  path_params = {
    organizationClusterUid = "00000000-0000-0000-0000-000000000001"
  }
}
```

### get_cluster_state

**HTTP:** `GET /scalegridserviceclusters/{scalegridServiceClusterUid}/state`  
**Category:** Scalegrid Services

Returns the state of an existing service.

**Path parameters:**
- `scalegridServiceClusterUid`

**Terraform example:**

```terraform
data "scalegrid_operation" "get_cluster_state" {
  operation_id = "get_cluster_state"
  path_params = {
    scalegridServiceClusterUid = "00000000-0000-0000-0000-000000000001"
  }
}
```

### get_cluster_state_1

**HTTP:** `GET /clusters/{organizationClusterUid}/state`  
**Category:** Clusters

Returns the state of an existing cluster.

**Path parameters:**
- `organizationClusterUid`

**Terraform example:**

```terraform
data "scalegrid_operation" "get_cluster_state_1" {
  operation_id = "get_cluster_state_1"
  path_params = {
    organizationClusterUid = "00000000-0000-0000-0000-000000000001"
  }
}
```

### get_cluster_storages

**HTTP:** `GET /clusters/{organizationClusterUid}/storages`  
**Category:** Clusters

Returns a list of available storage types in a specific cluster.

**Path parameters:**
- `organizationClusterUid`

**Terraform example:**

```terraform
data "scalegrid_operation" "get_cluster_storages" {
  operation_id = "get_cluster_storages"
  path_params = {
    organizationClusterUid = "00000000-0000-0000-0000-000000000001"
  }
}
```

### get_current_organization_user_by_id

**HTTP:** `GET /identity/organizations/users/{userId}`  
**Category:** Users

Returns an existing user.

**Path parameters:**
- `userId`

**Terraform example:**

```terraform
data "scalegrid_operation" "get_current_organization_user_by_id" {
  operation_id = "get_current_organization_user_by_id"
  path_params = {
    userId = "00000000-0000-0000-0000-000000000001"
  }
}
```

### get_current_organization_user_by_token

**HTTP:** `GET /identity/organizations/users/user`  
**Category:** Users

Returns an existing user.

**Terraform example:**

```terraform
data "scalegrid_operation" "get_current_organization_user_by_token" {
  operation_id = "get_current_organization_user_by_token"
}
```

### get_dashboards

**HTTP:** `POST /dashboards/scalegridservices/dashboards/list`  
**Category:** Dashboards

Returns a list of available dashboards for your service.

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_dashboards" {
  operation_id = "get_dashboards"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### get_event_list

**HTTP:** `POST /scalegridserviceclusters/{scalegridServiceClusterUid}/instances/{instanceUid}/events/list`  
**Category:** Scalegrid Services

Returns a list of events from an existing instance.

**Path parameters:**
- `scalegridServiceClusterUid`
- `instanceUid`

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_event_list" {
  operation_id = "get_event_list"
  path_params = {
    scalegridServiceClusterUid = "00000000-0000-0000-0000-000000000001"
    instanceUid = "00000000-0000-0000-0000-000000000001"
  }
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### get_ingresses

**HTTP:** `POST /scalegridservices/ingresstypes/list`  
**Category:** Scalegrid Services

Returns a list of Scalegrid managed services ingress options.

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_ingresses" {
  operation_id = "get_ingresses"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### get_instance_state

**HTTP:** `GET /scalegridserviceclusters/{scalegridServiceClusterUid}/instances/{instanceUid}/state`  
**Category:** Scalegrid Services

Returns the state of an existing instance.

**Path parameters:**
- `scalegridServiceClusterUid`
- `instanceUid`

**Terraform example:**

```terraform
data "scalegrid_operation" "get_instance_state" {
  operation_id = "get_instance_state"
  path_params = {
    scalegridServiceClusterUid = "00000000-0000-0000-0000-000000000001"
    instanceUid = "00000000-0000-0000-0000-000000000001"
  }
}
```

### get_invites

**HTTP:** `POST /identity/organizations/users/invite/list`  
**Category:** Users

Returns a list of user invites.

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_invites" {
  operation_id = "get_invites"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### get_latest_cluster_version

**HTTP:** `GET /clusters/clusterversion`  
**Category:** Clusters

Returns the latest cluster version.

**Terraform example:**

```terraform
data "scalegrid_operation" "get_latest_cluster_version" {
  operation_id = "get_latest_cluster_version"
}
```

### get_log_list

**HTTP:** `POST /scalegridserviceclusters/{scalegridServiceClusterUid}/instances/{instanceUid}/logs/list`  
**Category:** Scalegrid Services

Returns a list of logs from an existing instance.

**Path parameters:**
- `scalegridServiceClusterUid`
- `instanceUid`

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_log_list" {
  operation_id = "get_log_list"
  path_params = {
    scalegridServiceClusterUid = "00000000-0000-0000-0000-000000000001"
    instanceUid = "00000000-0000-0000-0000-000000000001"
  }
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### get_organization

**HTTP:** `GET /organization`  
**Category:** Organization

Returns an existing organization.

**Terraform example:**

```terraform
data "scalegrid_operation" "get_organization" {
  operation_id = "get_organization"
}
```

### get_organization_cluster

**HTTP:** `GET /clusters/organizationClusters/{organizationClusterUid}`  
**Category:** Clusters

Returns a specific cluster.

**Path parameters:**
- `organizationClusterUid`

**Terraform example:**

```terraform
data "scalegrid_operation" "get_organization_cluster" {
  operation_id = "get_organization_cluster"
  path_params = {
    organizationClusterUid = "00000000-0000-0000-0000-000000000001"
  }
}
```

### get_organization_clusters

**HTTP:** `POST /clusters/clusters/list`  
**Category:** Clusters

Returns a list of clusters.

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_organization_clusters" {
  operation_id = "get_organization_clusters"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### get_permissions

**HTTP:** `POST /identity/roles/permissions/list`  
**Category:** Roles

Returns available permissions for roles.

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_permissions" {
  operation_id = "get_permissions"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### get_recommended_settings

**HTTP:** `POST /scalegridserviceclusters/recommendation`  
**Category:** Scalegrid Services

Returns recommended settings for a service based on the provided parameters.

**Request body example:**

```json
{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"scalegridServiceVersionUid":"e2e-smoke","cpu":0,"memory":0,"storage":0,"numberOfReplicas":0,"recommendations":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","name":"e2e-smoke","recommended":"e2e-smoke","valid":false}],"valid":false}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_recommended_settings" {
  operation_id = "get_recommended_settings"
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"scalegridServiceVersionUid\":\"e2e-smoke\",\"cpu\":0,\"memory\":0,\"storage\":0,\"numberOfReplicas\":0,\"recommendations\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"recommended\":\"e2e-smoke\",\"valid\":false}],\"valid\":false}"
}
```

### get_restore_list

**HTTP:** `POST /backups/restores/list`  
**Category:** Backups

Returns a list of existing restores.

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_restore_list" {
  operation_id = "get_restore_list"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### get_role

**HTTP:** `GET /identity/organizations/roles/{roleId}`  
**Category:** Roles

Returns an existing role.

**Path parameters:**
- `roleId`

**Terraform example:**

```terraform
data "scalegrid_operation" "get_role" {
  operation_id = "get_role"
  path_params = {
    roleId = "00000000-0000-0000-0000-000000000001"
  }
}
```

### get_roles

**HTTP:** `POST /identity/organizations/roles/list`  
**Category:** Roles

Returns a list of roles.

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_roles" {
  operation_id = "get_roles"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### get_scalegrid_service_cluster

**HTTP:** `GET /scalegridserviceclusters/{scalegridServiceClusterUid}`  
**Category:** Scalegrid Services

Returns an existing Scalegrid managed service.

**Path parameters:**
- `scalegridServiceClusterUid`

**Terraform example:**

```terraform
data "scalegrid_operation" "get_scalegrid_service_cluster" {
  operation_id = "get_scalegrid_service_cluster"
  path_params = {
    scalegridServiceClusterUid = "00000000-0000-0000-0000-000000000001"
  }
}
```

### get_scalegrid_service_clusters

**HTTP:** `POST /scalegridserviceclusters/list`  
**Category:** Scalegrid Services

Returns a list of Scalegrid managed services.

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_scalegrid_service_clusters" {
  operation_id = "get_scalegrid_service_clusters"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### get_scalegrid_service_types

**HTTP:** `POST /scalegridservices/scalegridservicetypes/list`  
**Category:** Scalegrid Services

Returns a list of Scalegrid managed service types.

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_scalegrid_service_types" {
  operation_id = "get_scalegrid_service_types"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### get_scalegrid_service_versions

**HTTP:** `POST /scalegridservices/scalegridserviceversions/list`  
**Category:** Scalegrid Services

Returns a list of Scalegrid managed service versions.

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_scalegrid_service_versions" {
  operation_id = "get_scalegrid_service_versions"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### get_services_alert_templates

**HTTP:** `POST /alerts/templates/list`  
**Category:** Alerts

Returns a list of available alerts for your service.

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_services_alert_templates" {
  operation_id = "get_services_alert_templates"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### get_users

**HTTP:** `POST /identity/organizations/users/list`  
**Category:** Users

Returns a list of users.

**Request body example:**

```json
{"pageIndex":0,"pageSize":0,"sortParameters":[{"attributeName":"e2e-smoke","direction":"ASC"}],"filters":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operator":"LIKE","valid":false}],"conjunction":"AND"}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "get_users" {
  operation_id = "get_users"
  body_json = "{\"pageIndex\":0,\"pageSize\":0,\"sortParameters\":[{\"attributeName\":\"e2e-smoke\",\"direction\":\"ASC\"}],\"filters\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operator\":\"LIKE\",\"valid\":false}],\"conjunction\":\"AND\"}"
}
```

### invite_user_for_current_organization

**HTTP:** `POST /identity/organizations/users/invite/{email}`  
**Category:** Users

Invites a user via email, to join the organization.

**Path parameters:**
- `email`

**Terraform example:**

```terraform
data "scalegrid_operation" "invite_user_for_current_organization" {
  operation_id = "invite_user_for_current_organization"
  path_params = {
    email = "e2e-smoke@example.com"
  }
}
```

### onboard_current_user

**HTTP:** `POST /identity/users/onboard`  
**Category:** Users

Onboards an existing user.

**Terraform example:**

```terraform
data "scalegrid_operation" "onboard_current_user" {
  operation_id = "onboard_current_user"
}
```

### replace_instance

**HTTP:** `POST /scalegridserviceclusters/instances/{instanceUid}/replace`  
**Category:** Scalegrid Services

Replaces an existing instance.

**Path parameters:**
- `instanceUid`

**Terraform example:**

```terraform
data "scalegrid_operation" "replace_instance" {
  operation_id = "replace_instance"
  path_params = {
    instanceUid = "00000000-0000-0000-0000-000000000001"
  }
}
```

### update_backup_locale

**HTTP:** `PUT /backuplocales`  
**Category:** Backup Locales

Updates an existing backup locale.

**Request body example:**

```json
{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","creationTime":"e2e-smoke","labels":["e2e-smoke"],"backupLocaleRoles":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","permissions":[{"errors":[{}],"name":"e2e-smoke","category":"e2e-smoke","permissionName":"e2e-smoke","resourceServerIdentifier":"e2e-smoke","valid":false}],"modificationAllowed":false,"valid":false}],"organizationUid":"e2e-smoke","cloudProvider":"AWS","platformId":"e2e-smoke","backupStorages":["STORAGE_CLASS"],"objectStorageName":"e2e-smoke","valid":false}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "update_backup_locale" {
  operation_id = "update_backup_locale"
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"creationTime\":\"e2e-smoke\",\"labels\":[\"e2e-smoke\"],\"backupLocaleRoles\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"permissions\":[{\"errors\":[{}],\"name\":\"e2e-smoke\",\"category\":\"e2e-smoke\",\"permissionName\":\"e2e-smoke\",\"resourceServerIdentifier\":\"e2e-smoke\",\"valid\":false}],\"modificationAllowed\":false,\"valid\":false}],\"organizationUid\":\"e2e-smoke\",\"cloudProvider\":\"AWS\",\"platformId\":\"e2e-smoke\",\"backupStorages\":[\"STORAGE_CLASS\"],\"objectStorageName\":\"e2e-smoke\",\"valid\":false}"
}
```

### update_cluster

**HTTP:** `PUT /clusters/{organizationClusterUid}`  
**Category:** Clusters

Updates an existing cluster.

**Path parameters:**
- `organizationClusterUid`

**Request body example:**

```json
{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","cloudProvider":"AWS","organizationClusterName":"e2e-smoke","clusterKubernetesNamespace":"e2e-smoke","affinitySelector":["e2e-smoke"],"antiAffinity":["e2e-smoke"],"nodeSelector":["e2e-smoke"],"storage":{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"type":"STORAGE_CLASS","name":"e2e-smoke","allowVolumeExpansion":false,"valid":false},"organizationClusterDescription":"e2e-smoke","nodeName":"e2e-smoke","clusterUid":"e2e-smoke","labels":["e2e-smoke"],"organizationClusterRoles":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","permissions":[{"errors":[{}],"name":"e2e-smoke","category":"e2e-smoke","permissionName":"e2e-smoke","resourceServerIdentifier":"e2e-smoke","valid":false}],"modificationAllowed":false,"valid":false}],"os":"LINUX","architecture":"AMD64","helmDeploymentCommand":"e2e-smoke","region":"e2e-smoke","ingressControllers":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"from":0,"to":0,"name":"e2e-smoke","ingressType":"PUBLICLY_EXPOSED","valid":false}],"platformID":"e2e-smoke","tolerations":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"key":"e2e-smoke","value":"e2e-smoke","operation":"Exists","effect":"NoSchedule","valid":false}],"externalCidr":{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"cidr":"e2e-smoke","name":"e2e-smoke","valid":false},"podCidr":{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"cidr":"e2e-smoke","name":"e2e-smoke","valid":false},"tunnelCidr":{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"cidr":"e2e-smoke","name":"e2e-smoke","valid":false},"kubernetesApiUrl":"e2e-smoke","valid":false}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "update_cluster" {
  operation_id = "update_cluster"
  path_params = {
    organizationClusterUid = "00000000-0000-0000-0000-000000000001"
  }
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"cloudProvider\":\"AWS\",\"organizationClusterName\":\"e2e-smoke\",\"clusterKubernetesNamespace\":\"e2e-smoke\",\"affinitySelector\":[\"e2e-smoke\"],\"antiAffinity\":[\"e2e-smoke\"],\"nodeSelector\":[\"e2e-smoke\"],\"storage\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"type\":\"STORAGE_CLASS\",\"name\":\"e2e-smoke\",\"allowVolumeExpansion\":false,\"valid\":false},\"organizationClusterDescription\":\"e2e-smoke\",\"nodeName\":\"e2e-smoke\",\"clusterUid\":\"e2e-smoke\",\"labels\":[\"e2e-smoke\"],\"organizationClusterRoles\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"permissions\":[{\"errors\":[{}],\"name\":\"e2e-smoke\",\"category\":\"e2e-smoke\",\"permissionName\":\"e2e-smoke\",\"resourceServerIdentifier\":\"e2e-smoke\",\"valid\":false}],\"modificationAllowed\":false,\"valid\":false}],\"os\":\"LINUX\",\"architecture\":\"AMD64\",\"helmDeploymentCommand\":\"e2e-smoke\",\"region\":\"e2e-smoke\",\"ingressControllers\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"from\":0,\"to\":0,\"name\":\"e2e-smoke\",\"ingressType\":\"PUBLICLY_EXPOSED\",\"valid\":false}],\"platformID\":\"e2e-smoke\",\"tolerations\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operation\":\"Exists\",\"effect\":\"NoSchedule\",\"valid\":false}],\"externalCidr\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"cidr\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"valid\":false},\"podCidr\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"cidr\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"valid\":false},\"tunnelCidr\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"cidr\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"valid\":false},\"kubernetesApiUrl\":\"e2e-smoke\",\"valid\":false}"
}
```

### update_current_organization_user

**HTTP:** `PUT /identity/organizations/users/{userId}`  
**Category:** Users

Updates an existing user.

**Path parameters:**
- `userId`

**Request body example:**

```json
{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","email":"e2e-smoke","password":"e2e-smoke","organizationId":"e2e-smoke","connection":"e2e-smoke","roles":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","permissions":[{"errors":[{}],"name":"e2e-smoke","category":"e2e-smoke","permissionName":"e2e-smoke","resourceServerIdentifier":"e2e-smoke","valid":false}],"modificationAllowed":false,"valid":false}],"creationDate":"e2e-smoke","mfaEnabled":false,"mfaKey":"e2e-smoke","valid":false}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "update_current_organization_user" {
  operation_id = "update_current_organization_user"
  path_params = {
    userId = "00000000-0000-0000-0000-000000000001"
  }
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"email\":\"e2e-smoke\",\"password\":\"e2e-smoke\",\"organizationId\":\"e2e-smoke\",\"connection\":\"e2e-smoke\",\"roles\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"permissions\":[{\"errors\":[{}],\"name\":\"e2e-smoke\",\"category\":\"e2e-smoke\",\"permissionName\":\"e2e-smoke\",\"resourceServerIdentifier\":\"e2e-smoke\",\"valid\":false}],\"modificationAllowed\":false,\"valid\":false}],\"creationDate\":\"e2e-smoke\",\"mfaEnabled\":false,\"mfaKey\":\"e2e-smoke\",\"valid\":false}"
}
```

### update_organization_credit_card

**HTTP:** `PATCH /organization/creditCard`  
**Category:** Organization

Updates an existing organization credit card details.

**Request body example:**

```json
{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"creditCardToken":"e2e-smoke","valid":false}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "update_organization_credit_card" {
  operation_id = "update_organization_credit_card"
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"creditCardToken\":\"e2e-smoke\",\"valid\":false}"
}
```

### update_own_organization

**HTTP:** `PUT /organization`  
**Category:** Organization

Updates an existing organization.

**Request body example:**

```json
{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","name":"e2e-smoke","email":"e2e-smoke","status":"TRIAL","trialExpirationDate":"e2e-smoke","invoicingType":"MANUAL","paymentMethod":"BANK_TRANSFER","defaultInviteRole":"e2e-smoke","statusHistory":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","newStatus":"TRIAL","oldStatus":"TRIAL","timestamp":"e2e-smoke","valid":false}],"creationTime":"e2e-smoke","creditCardBrand":"e2e-smoke","creditCardLast4":"e2e-smoke","maxNumberOfRoles":0,"valid":false}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "update_own_organization" {
  operation_id = "update_own_organization"
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"email\":\"e2e-smoke\",\"status\":\"TRIAL\",\"trialExpirationDate\":\"e2e-smoke\",\"invoicingType\":\"MANUAL\",\"paymentMethod\":\"BANK_TRANSFER\",\"defaultInviteRole\":\"e2e-smoke\",\"statusHistory\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"newStatus\":\"TRIAL\",\"oldStatus\":\"TRIAL\",\"timestamp\":\"e2e-smoke\",\"valid\":false}],\"creationTime\":\"e2e-smoke\",\"creditCardBrand\":\"e2e-smoke\",\"creditCardLast4\":\"e2e-smoke\",\"maxNumberOfRoles\":0,\"valid\":false}"
}
```

### update_role

**HTTP:** `PUT /identity/organizations/roles/{roleId}`  
**Category:** Roles

Updates an existing role.

**Path parameters:**
- `roleId`

**Request body example:**

```json
{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","permissions":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"name":"e2e-smoke","category":"e2e-smoke","permissionName":"e2e-smoke","resourceServerIdentifier":"e2e-smoke","valid":false}],"modificationAllowed":false,"valid":false}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "update_role" {
  operation_id = "update_role"
  path_params = {
    roleId = "00000000-0000-0000-0000-000000000001"
  }
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"permissions\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"name\":\"e2e-smoke\",\"category\":\"e2e-smoke\",\"permissionName\":\"e2e-smoke\",\"resourceServerIdentifier\":\"e2e-smoke\",\"valid\":false}],\"modificationAllowed\":false,\"valid\":false}"
}
```

### update_scalegrid_service_cluster

**HTTP:** `PUT /scalegridserviceclusters/{scalegridServiceClusterUid}`  
**Category:** Scalegrid Services

Updates an existing Scalegrid managed service.

**Path parameters:**
- `scalegridServiceClusterUid`

**Request body example:**

```json
{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","creationTime":"e2e-smoke","hasTelemetry":false,"ingress":"PUBLICLY_EXPOSED","instances":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","organizationClusterUid":"e2e-smoke","scalegridServiceClusterUid":"e2e-smoke","name":"e2e-smoke","kubernetesNamespace":"e2e-smoke","affinitySelector":["e2e-smoke"],"antiAffinitySelector":["e2e-smoke"],"nodeSelector":["e2e-smoke"],"nodeName":"e2e-smoke","autoScalingType":"NONE","autoScalingCpuPercentage":0,"autoScalingMemoryPercentage":0,"storageType":"STORAGE_CLASS","storageName":"e2e-smoke","inboundCIDRs":[{"errors":[{}],"cidr":"e2e-smoke","name":"e2e-smoke","valid":false}],"failoverPriority":0,"additionalSettings":"e2e-smoke","resources":[{"name":"e2e-smoke","resource":"e2e-smoke","certificateExpirationDate":"e2e-smoke"}],"creationTime":"e2e-smoke","cpuSizeMRequest":0,"memorySizeMbRequest":0,"storageSizeMb":0,"cpuSizeMMaxLimit":0,"cpuSizeMMinLimit":0,"memorySizeMbMaxLimit":0,"memorySizeMbMinLimit":0,"alerts":[{"errors":[{}],"uid":"e2e-smoke","organizationClusterUid":"e2e-smoke","scalegridServiceClusterUid":"e2e-smoke","scalegridServiceClusterName":"e2e-smoke","scalegridServiceUid":"e2e-smoke","instanceUid":"e2e-smoke","instanceName":"e2e-smoke","organizationUid":"e2e-smoke","serviceAlertUid":"e2e-smoke","sendResolved":false,"emailsToNotify":["e2e-smoke"],"parameters":[{}],"alertType":"POD","valid":false}],"labels":["e2e-smoke"],"ruleBasedVPAs":[{"errors":[{}],"uid":"e2e-smoke","cpuSizeMRequest":0,"memorySizeMbRequest":0,"cpuSizeMMaxLimit":0,"cpuSizeMMinLimit":0,"memorySizeMbMaxLimit":0,"memorySizeMbMinLimit":0,"cronJob":"e2e-smoke","autoScalingType":"NONE","valid":false}],"architecture":"AMD64","os":"LINUX","tolerations":[{"errors":[{}],"key":"e2e-smoke","value":"e2e-smoke","operation":"Exists","effect":"NoSchedule","valid":false}],"valid":false}],"scalegridServiceVersion":{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","scalegridServiceEnum":"RabbitMQ","versionName":"e2e-smoke","preRequiredScalegridVersionsForUpgrade":"e2e-smoke","availableForNewDeployments":false,"description":"e2e-smoke","allowedAutoScalingTypes":["NONE"],"replicationTypes":[{"errors":[{"key":"e2e-smoke","error":{}}],"uid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","numberOfReplicas":0,"valid":false}],"backupInstallInstructions":"e2e-smoke","clusterType":"ACTIVE_ACTIVE","architectures":["AMD64"],"oss":["LINUX"],"settings":{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"settings":[{"errors":[{}],"uid":"e2e-smoke","name":"e2e-smoke","value":"e2e-smoke","valueType":"NUMBER","options":"e2e-smoke","enabled":false,"description":"e2e-smoke","parentSettingUid":"e2e-smoke","valid":false}],"extensions":[{"errors":[{}],"uid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","enabled":false,"settings":[{}],"requiredExtensionUids":["e2e-smoke"],"valid":false}],"serviceType":"RabbitMQ","valid":false},"instanceSettings":{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"settings":[{"errors":[{}],"uid":"e2e-smoke","name":"e2e-smoke","value":"e2e-smoke","valueType":"NUMBER","options":"e2e-smoke","enabled":false,"description":"e2e-smoke","parentSettingUid":"e2e-smoke","valid":false}],"extensions":[{"errors":[{}],"uid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","enabled":false,"settings":[{}],"requiredExtensionUids":["e2e-smoke"],"valid":false}],"serviceType":"RabbitMQ","valid":false},"compatibleScalegridVersionsBackups":"e2e-smoke","valid":false},"replicationTypeUid":"e2e-smoke","additionalSettings":"e2e-smoke","resources":[{"name":"e2e-smoke","resource":"e2e-smoke","certificateExpirationDate":"e2e-smoke"}],"autoPatch":false,"hasLatestPatch":false,"alerts":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","organizationClusterUid":"e2e-smoke","scalegridServiceClusterUid":"e2e-smoke","scalegridServiceClusterName":"e2e-smoke","scalegridServiceUid":"e2e-smoke","instanceUid":"e2e-smoke","instanceName":"e2e-smoke","organizationUid":"e2e-smoke","serviceAlertUid":"e2e-smoke","sendResolved":false,"emailsToNotify":["e2e-smoke"],"parameters":[{"errors":[{}],"key":"e2e-smoke","value":"e2e-smoke","isCustomerExposed":false,"unit":"e2e-smoke","minValue":0,"maxValue":0,"valid":false}],"alertType":"POD","valid":false}],"serviceClusterRoles":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","roleUid":"e2e-smoke","valid":false}],"autoPatchCronJob":"e2e-smoke","ruleBasedVPAs":[{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","cpuSizeMRequest":0,"memorySizeMbRequest":0,"cpuSizeMMaxLimit":0,"cpuSizeMMinLimit":0,"memorySizeMbMaxLimit":0,"memorySizeMbMinLimit":0,"cronJob":"e2e-smoke","autoScalingType":"NONE","valid":false}],"backupConfigurations":{"errors":[{"key":"e2e-smoke","error":{"message":"e2e-smoke","code":"e2e-smoke"}}],"uid":"e2e-smoke","scalegridServiceClusterUid":"e2e-smoke","backupConfigurations":[{"errors":[{"key":"e2e-smoke","error":{}}],"uid":"e2e-smoke","backupLocaleUid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","backupCronJob":"e2e-smoke","timeToLiveDays":0,"maxSizeInBackupsMb":0,"backupType":{"errors":[{}],"uid":"e2e-smoke","name":"e2e-smoke","description":"e2e-smoke","scalegridServiceVersionList":["00000000-0000-0000-0000-000000000001"],"cloudList":["AWS"],"restoreOptions":"SAME_PLATFORM","backupMethod":"DISC_SNAPSHOT","backupType":"PITR","restoreType":"STANDALONE","backupStorageTypes":["STORAGE_CLASS"],"scalegridServiceEnum":"RabbitMQ","backupOperatorDeployment":"TARGET_INSTANCE","isTargetOptionRequiredForBackup":false,"supportsMultiRegion":false,"valid":false},"maxNumberOfBackups":0,"targetName":"e2e-smoke","compressed":false,"scalegridServiceClusterUid":"e2e-smoke","valid":false}],"timeToLiveDays":0,"maxSizeInBackupsMb":0,"maxNumberOfBackups":0,"valid":false},"cpuSizeMRequest":0,"memorySizeMbRequest":0,"cpuSizeMMaxLimit":0,"cpuSizeMMinLimit":0,"memorySizeMbMaxLimit":0,"memorySizeMbMinLimit":0,"autoScalingType":"NONE","autoScalingCpuPercentage":0,"autoScalingMemoryPercentage":0,"valid":false}
```

**Terraform example:**

```terraform
data "scalegrid_operation" "update_scalegrid_service_cluster" {
  operation_id = "update_scalegrid_service_cluster"
  path_params = {
    scalegridServiceClusterUid = "00000000-0000-0000-0000-000000000001"
  }
  body_json = "{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"creationTime\":\"e2e-smoke\",\"hasTelemetry\":false,\"ingress\":\"PUBLICLY_EXPOSED\",\"instances\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"organizationClusterUid\":\"e2e-smoke\",\"scalegridServiceClusterUid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"kubernetesNamespace\":\"e2e-smoke\",\"affinitySelector\":[\"e2e-smoke\"],\"antiAffinitySelector\":[\"e2e-smoke\"],\"nodeSelector\":[\"e2e-smoke\"],\"nodeName\":\"e2e-smoke\",\"autoScalingType\":\"NONE\",\"autoScalingCpuPercentage\":0,\"autoScalingMemoryPercentage\":0,\"storageType\":\"STORAGE_CLASS\",\"storageName\":\"e2e-smoke\",\"inboundCIDRs\":[{\"errors\":[{}],\"cidr\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"valid\":false}],\"failoverPriority\":0,\"additionalSettings\":\"e2e-smoke\",\"resources\":[{\"name\":\"e2e-smoke\",\"resource\":\"e2e-smoke\",\"certificateExpirationDate\":\"e2e-smoke\"}],\"creationTime\":\"e2e-smoke\",\"cpuSizeMRequest\":0,\"memorySizeMbRequest\":0,\"storageSizeMb\":0,\"cpuSizeMMaxLimit\":0,\"cpuSizeMMinLimit\":0,\"memorySizeMbMaxLimit\":0,\"memorySizeMbMinLimit\":0,\"alerts\":[{\"errors\":[{}],\"uid\":\"e2e-smoke\",\"organizationClusterUid\":\"e2e-smoke\",\"scalegridServiceClusterUid\":\"e2e-smoke\",\"scalegridServiceClusterName\":\"e2e-smoke\",\"scalegridServiceUid\":\"e2e-smoke\",\"instanceUid\":\"e2e-smoke\",\"instanceName\":\"e2e-smoke\",\"organizationUid\":\"e2e-smoke\",\"serviceAlertUid\":\"e2e-smoke\",\"sendResolved\":false,\"emailsToNotify\":[\"e2e-smoke\"],\"parameters\":[{}],\"alertType\":\"POD\",\"valid\":false}],\"labels\":[\"e2e-smoke\"],\"ruleBasedVPAs\":[{\"errors\":[{}],\"uid\":\"e2e-smoke\",\"cpuSizeMRequest\":0,\"memorySizeMbRequest\":0,\"cpuSizeMMaxLimit\":0,\"cpuSizeMMinLimit\":0,\"memorySizeMbMaxLimit\":0,\"memorySizeMbMinLimit\":0,\"cronJob\":\"e2e-smoke\",\"autoScalingType\":\"NONE\",\"valid\":false}],\"architecture\":\"AMD64\",\"os\":\"LINUX\",\"tolerations\":[{\"errors\":[{}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"operation\":\"Exists\",\"effect\":\"NoSchedule\",\"valid\":false}],\"valid\":false}],\"scalegridServiceVersion\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"scalegridServiceEnum\":\"RabbitMQ\",\"versionName\":\"e2e-smoke\",\"preRequiredScalegridVersionsForUpgrade\":\"e2e-smoke\",\"availableForNewDeployments\":false,\"description\":\"e2e-smoke\",\"allowedAutoScalingTypes\":[\"NONE\"],\"replicationTypes\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{}}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"numberOfReplicas\":0,\"valid\":false}],\"backupInstallInstructions\":\"e2e-smoke\",\"clusterType\":\"ACTIVE_ACTIVE\",\"architectures\":[\"AMD64\"],\"oss\":[\"LINUX\"],\"settings\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"settings\":[{\"errors\":[{}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"valueType\":\"NUMBER\",\"options\":\"e2e-smoke\",\"enabled\":false,\"description\":\"e2e-smoke\",\"parentSettingUid\":\"e2e-smoke\",\"valid\":false}],\"extensions\":[{\"errors\":[{}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"enabled\":false,\"settings\":[{}],\"requiredExtensionUids\":[\"e2e-smoke\"],\"valid\":false}],\"serviceType\":\"RabbitMQ\",\"valid\":false},\"instanceSettings\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"settings\":[{\"errors\":[{}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"valueType\":\"NUMBER\",\"options\":\"e2e-smoke\",\"enabled\":false,\"description\":\"e2e-smoke\",\"parentSettingUid\":\"e2e-smoke\",\"valid\":false}],\"extensions\":[{\"errors\":[{}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"enabled\":false,\"settings\":[{}],\"requiredExtensionUids\":[\"e2e-smoke\"],\"valid\":false}],\"serviceType\":\"RabbitMQ\",\"valid\":false},\"compatibleScalegridVersionsBackups\":\"e2e-smoke\",\"valid\":false},\"replicationTypeUid\":\"e2e-smoke\",\"additionalSettings\":\"e2e-smoke\",\"resources\":[{\"name\":\"e2e-smoke\",\"resource\":\"e2e-smoke\",\"certificateExpirationDate\":\"e2e-smoke\"}],\"autoPatch\":false,\"hasLatestPatch\":false,\"alerts\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"organizationClusterUid\":\"e2e-smoke\",\"scalegridServiceClusterUid\":\"e2e-smoke\",\"scalegridServiceClusterName\":\"e2e-smoke\",\"scalegridServiceUid\":\"e2e-smoke\",\"instanceUid\":\"e2e-smoke\",\"instanceName\":\"e2e-smoke\",\"organizationUid\":\"e2e-smoke\",\"serviceAlertUid\":\"e2e-smoke\",\"sendResolved\":false,\"emailsToNotify\":[\"e2e-smoke\"],\"parameters\":[{\"errors\":[{}],\"key\":\"e2e-smoke\",\"value\":\"e2e-smoke\",\"isCustomerExposed\":false,\"unit\":\"e2e-smoke\",\"minValue\":0,\"maxValue\":0,\"valid\":false}],\"alertType\":\"POD\",\"valid\":false}],\"serviceClusterRoles\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"roleUid\":\"e2e-smoke\",\"valid\":false}],\"autoPatchCronJob\":\"e2e-smoke\",\"ruleBasedVPAs\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"cpuSizeMRequest\":0,\"memorySizeMbRequest\":0,\"cpuSizeMMaxLimit\":0,\"cpuSizeMMinLimit\":0,\"memorySizeMbMaxLimit\":0,\"memorySizeMbMinLimit\":0,\"cronJob\":\"e2e-smoke\",\"autoScalingType\":\"NONE\",\"valid\":false}],\"backupConfigurations\":{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{\"message\":\"e2e-smoke\",\"code\":\"e2e-smoke\"}}],\"uid\":\"e2e-smoke\",\"scalegridServiceClusterUid\":\"e2e-smoke\",\"backupConfigurations\":[{\"errors\":[{\"key\":\"e2e-smoke\",\"error\":{}}],\"uid\":\"e2e-smoke\",\"backupLocaleUid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"backupCronJob\":\"e2e-smoke\",\"timeToLiveDays\":0,\"maxSizeInBackupsMb\":0,\"backupType\":{\"errors\":[{}],\"uid\":\"e2e-smoke\",\"name\":\"e2e-smoke\",\"description\":\"e2e-smoke\",\"scalegridServiceVersionList\":[\"00000000-0000-0000-0000-000000000001\"],\"cloudList\":[\"AWS\"],\"restoreOptions\":\"SAME_PLATFORM\",\"backupMethod\":\"DISC_SNAPSHOT\",\"backupType\":\"PITR\",\"restoreType\":\"STANDALONE\",\"backupStorageTypes\":[\"STORAGE_CLASS\"],\"scalegridServiceEnum\":\"RabbitMQ\",\"backupOperatorDeployment\":\"TARGET_INSTANCE\",\"isTargetOptionRequiredForBackup\":false,\"supportsMultiRegion\":false,\"valid\":false},\"maxNumberOfBackups\":0,\"targetName\":\"e2e-smoke\",\"compressed\":false,\"scalegridServiceClusterUid\":\"e2e-smoke\",\"valid\":false}],\"timeToLiveDays\":0,\"maxSizeInBackupsMb\":0,\"maxNumberOfBackups\":0,\"valid\":false},\"cpuSizeMRequest\":0,\"memorySizeMbRequest\":0,\"cpuSizeMMaxLimit\":0,\"cpuSizeMMinLimit\":0,\"memorySizeMbMaxLimit\":0,\"memorySizeMbMinLimit\":0,\"autoScalingType\":\"NONE\",\"autoScalingCpuPercentage\":0,\"autoScalingMemoryPercentage\":0,\"valid\":false}"
}
```
