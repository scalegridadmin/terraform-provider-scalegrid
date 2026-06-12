---
page_title: "Operations"
description: |-
  Index of ScaleGrid provider operations.
---

# Operations

The complete documentation for every operation lives on the [`scalegrid_operation` resource page](../resources/operation). This index links to individual operation summaries.

- [`clone_scalegrid_service_cluster`](./clone_scalegrid_service_cluster) — `POST /scalegridserviceclusters/clone`
- [`create_backup`](./create_backup) — `POST /backupconfigurations/{backupConfigurationUid}/backups/run`
- [`create_backup_locale`](./create_backup_locale) — `POST /backuplocales`
- [`create_restore`](./create_restore) — `POST /backups/restore`
- [`create_role`](./create_role) — `POST /identity/organizations/roles`
- [`create_scalegrid_cluster`](./create_scalegrid_cluster) — `POST /clusters`
- [`create_scalegrid_service_cluster`](./create_scalegrid_service_cluster) — `POST /scalegridserviceclusters`
- [`delete_backup`](./delete_backup) — `DELETE /backups/{backupConfigurationsUid}/backups/{backupUid}`
- [`delete_backup_locale`](./delete_backup_locale) — `DELETE /backuplocales/{backupLocaleUid}`
- [`delete_cluster`](./delete_cluster) — `DELETE /clusters/{organizationClusterUid}`
- [`delete_invite`](./delete_invite) — `DELETE /identity/organizations/users/invite/{inviteUid}`
- [`delete_organization`](./delete_organization) — `DELETE /organization`
- [`delete_role`](./delete_role) — `DELETE /identity/organizations/roles/{roleId}`
- [`delete_scalegrid_service_cluster`](./delete_scalegrid_service_cluster) — `DELETE /scalegridserviceclusters/{scalegridServiceClusterUid}`
- [`delete_user_from_current_organization`](./delete_user_from_current_organization) — `DELETE /identity/organizations/users/{userId}`
- [`failover_instance`](./failover_instance) — `POST /scalegridserviceclusters/instances/{instanceUid}/failover`
- [`get_audits`](./get_audits) — `POST /audits/list`
- [`get_backup`](./get_backup) — `GET /backups/{backupUid}`
- [`get_backup_list`](./get_backup_list) — `POST /backups/list`
- [`get_backup_locale`](./get_backup_locale) — `GET /backuplocales/{backupLocaleUid}`
- [`get_backup_locale_list`](./get_backup_locale_list) — `POST /backuplocales/list`
- [`get_backup_locale_state`](./get_backup_locale_state) — `GET /backuplocales/{backupLocaleUid}/state`
- [`get_backup_types`](./get_backup_types) — `POST /backuptypes/list`
- [`get_certificate`](./get_certificate) — `GET /scalegridserviceclusters/{scalegridServiceClusterUid}/certificate`
- [`get_cloud_providers`](./get_cloud_providers) — `POST /clusters/cloudproviders/list`
- [`get_cluster_list`](./get_cluster_list) — `POST /clusters/list`
- [`get_cluster_metadata`](./get_cluster_metadata) — `GET /clusters/{organizationClusterUid}/metadata`
- [`get_cluster_state`](./get_cluster_state) — `GET /scalegridserviceclusters/{scalegridServiceClusterUid}/state`
- [`get_cluster_state_1`](./get_cluster_state_1) — `GET /clusters/{organizationClusterUid}/state`
- [`get_cluster_storages`](./get_cluster_storages) — `GET /clusters/{organizationClusterUid}/storages`
- [`get_current_organization_user_by_id`](./get_current_organization_user_by_id) — `GET /identity/organizations/users/{userId}`
- [`get_current_organization_user_by_token`](./get_current_organization_user_by_token) — `GET /identity/organizations/users/user`
- [`get_dashboards`](./get_dashboards) — `POST /dashboards/scalegridservices/dashboards/list`
- [`get_event_list`](./get_event_list) — `POST /scalegridserviceclusters/{scalegridServiceClusterUid}/instances/{instanceUid}/events/list`
- [`get_ingresses`](./get_ingresses) — `POST /scalegridservices/ingresstypes/list`
- [`get_instance_state`](./get_instance_state) — `GET /scalegridserviceclusters/{scalegridServiceClusterUid}/instances/{instanceUid}/state`
- [`get_invites`](./get_invites) — `POST /identity/organizations/users/invite/list`
- [`get_latest_cluster_version`](./get_latest_cluster_version) — `GET /clusters/clusterversion`
- [`get_log_list`](./get_log_list) — `POST /scalegridserviceclusters/{scalegridServiceClusterUid}/instances/{instanceUid}/logs/list`
- [`get_organization`](./get_organization) — `GET /organization`
- [`get_organization_cluster`](./get_organization_cluster) — `GET /clusters/organizationClusters/{organizationClusterUid}`
- [`get_organization_clusters`](./get_organization_clusters) — `POST /clusters/clusters/list`
- [`get_permissions`](./get_permissions) — `POST /identity/roles/permissions/list`
- [`get_recommended_settings`](./get_recommended_settings) — `POST /scalegridserviceclusters/recommendation`
- [`get_restore_list`](./get_restore_list) — `POST /backups/restores/list`
- [`get_role`](./get_role) — `GET /identity/organizations/roles/{roleId}`
- [`get_roles`](./get_roles) — `POST /identity/organizations/roles/list`
- [`get_scalegrid_service_cluster`](./get_scalegrid_service_cluster) — `GET /scalegridserviceclusters/{scalegridServiceClusterUid}`
- [`get_scalegrid_service_clusters`](./get_scalegrid_service_clusters) — `POST /scalegridserviceclusters/list`
- [`get_scalegrid_service_types`](./get_scalegrid_service_types) — `POST /scalegridservices/scalegridservicetypes/list`
- [`get_scalegrid_service_versions`](./get_scalegrid_service_versions) — `POST /scalegridservices/scalegridserviceversions/list`
- [`get_services_alert_templates`](./get_services_alert_templates) — `POST /alerts/templates/list`
- [`get_users`](./get_users) — `POST /identity/organizations/users/list`
- [`invite_user_for_current_organization`](./invite_user_for_current_organization) — `POST /identity/organizations/users/invite/{email}`
- [`onboard_current_user`](./onboard_current_user) — `POST /identity/users/onboard`
- [`replace_instance`](./replace_instance) — `POST /scalegridserviceclusters/instances/{instanceUid}/replace`
- [`update_backup_locale`](./update_backup_locale) — `PUT /backuplocales`
- [`update_cluster`](./update_cluster) — `PUT /clusters/{organizationClusterUid}`
- [`update_current_organization_user`](./update_current_organization_user) — `PUT /identity/organizations/users/{userId}`
- [`update_organization_credit_card`](./update_organization_credit_card) — `PATCH /organization/creditCard`
- [`update_own_organization`](./update_own_organization) — `PUT /organization`
- [`update_role`](./update_role) — `PUT /identity/organizations/roles/{roleId}`
- [`update_scalegrid_service_cluster`](./update_scalegrid_service_cluster) — `PUT /scalegridserviceclusters/{scalegridServiceClusterUid}`
