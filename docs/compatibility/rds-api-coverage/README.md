---
title: "RDS API Coverage"
seoTitle: "Relational Database Service API Coverage — Spinifex Docs"
description: "Every operation in the Amazon RDS API model and whether Spinifex implements it, for the managed PostgreSQL and MariaDB engines it offers on each build."
category: "Coverage"
sections:
  - overview
tags:
  - aws
  - compatibility
  - coverage
  - rds
  - databases
  - postgresql
  - mariadb
---

# RDS API Coverage

## Overview

Spinifex implements **26 of the 162** operations in the RDS `2014-10-31` API model, as pinned in `aws-sdk-go v1.55.8` — **16.0%**.

### Engines

Spinifex offers PostgreSQL and MariaDB. Each DB instance is one dedicated system-owned VM running the engine directly, launched from a platform AMI and hidden from the customer's EC2 API. `Engine` is fixed at create: there is no in-place engine change, no cross-engine snapshot restore and no migration between the two.

`mysql` is not an accepted engine and is not an alias for `mariadb`. MariaDB is offered under its own AWS engine name, exactly as AWS RDS offers it, so a client — including Terraform's `aws_db_instance` — must set `engine = "mariadb"`. Aliasing would report an engine and a version the instance is not running, and the discrepancy would propagate into `DescribeDBInstances`, parameter-group families and snapshot metadata.

Engine versions are pinned per engine. An `EngineVersion` naming anything but the pin is rejected, including a narrower minor version, because the AMI makes no promise about which minor it carries.

### The endpoint is private

The engine is reached over a customer-account ENI injected into a subnet of the DB subnet group, so the endpoint is reachable from inside the VPC only. The DB VM has other NICs that no customer security group governs, and the engine binds none of them — the port is not open there at all, rather than open and gated.

### TLS is required by default

Both engines enforce encrypted connections by default: `rds.force_ssl` on PostgreSQL and `require_secure_transport` on MariaDB. For MariaDB this is a deliberate divergence from AWS, which leaves it off. Both are boolean, modifiable and dynamic, so setting either to `0` in a parameter group restores plaintext without a reboot.

Enforcement requires that the connection is encrypted, not that the client validates the certificate. The engine serves a per-instance certificate signed by the cluster CA carrying both the ENI address and the DNS name, so a client that wants full verification can have it. A deployment holding no cluster CA cannot serve TLS at all, and a parameter group asking for enforcement there is refused rather than quietly ignored.

### Rejected parameters

A parameter whose omission would create a false safety, security or availability guarantee is rejected rather than silently dropped — `MultiAZ=true` on a single-AZ platform, `PubliclyAccessible=true` against a private endpoint, `StorageEncrypted=false` where unencrypted storage is not offered. Parameters that are merely inert are accepted as no-ops.

### Operations

| Operation | Status |
|---|---|
| `AddRoleToDBCluster` | ❌ Not implemented |
| `AddRoleToDBInstance` | ❌ Not implemented |
| `AddSourceIdentifierToSubscription` | ❌ Not implemented |
| `AddTagsToResource` | ✅ Implemented |
| `ApplyPendingMaintenanceAction` | ❌ Not implemented |
| `AuthorizeDBSecurityGroupIngress` | ❌ Not implemented |
| `BacktrackDBCluster` | ❌ Not implemented |
| `CancelExportTask` | ❌ Not implemented |
| `CopyDBClusterParameterGroup` | ❌ Not implemented |
| `CopyDBClusterSnapshot` | ❌ Not implemented |
| `CopyDBParameterGroup` | ❌ Not implemented |
| `CopyDBSnapshot` | ❌ Not implemented |
| `CopyOptionGroup` | ❌ Not implemented |
| `CreateBlueGreenDeployment` | ❌ Not implemented |
| `CreateCustomDBEngineVersion` | ❌ Not implemented |
| `CreateDBCluster` | 🚫 Not supported |
| `CreateDBClusterEndpoint` | ❌ Not implemented |
| `CreateDBClusterParameterGroup` | ❌ Not implemented |
| `CreateDBClusterSnapshot` | ❌ Not implemented |
| `CreateDBInstance` | ✅ Implemented |
| `CreateDBInstanceReadReplica` | 🚫 Not supported |
| `CreateDBParameterGroup` | ✅ Implemented |
| `CreateDBProxy` | ❌ Not implemented |
| `CreateDBProxyEndpoint` | ❌ Not implemented |
| `CreateDBSecurityGroup` | ❌ Not implemented |
| `CreateDBShardGroup` | ❌ Not implemented |
| `CreateDBSnapshot` | ✅ Implemented |
| `CreateDBSubnetGroup` | ✅ Implemented |
| `CreateEventSubscription` | ❌ Not implemented |
| `CreateGlobalCluster` | ❌ Not implemented |
| `CreateIntegration` | ❌ Not implemented |
| `CreateOptionGroup` | 🚫 Not supported |
| `CreateTenantDatabase` | ❌ Not implemented |
| `DeleteBlueGreenDeployment` | ❌ Not implemented |
| `DeleteCustomDBEngineVersion` | ❌ Not implemented |
| `DeleteDBCluster` | 🚫 Not supported |
| `DeleteDBClusterAutomatedBackup` | ❌ Not implemented |
| `DeleteDBClusterEndpoint` | ❌ Not implemented |
| `DeleteDBClusterParameterGroup` | ❌ Not implemented |
| `DeleteDBClusterSnapshot` | ❌ Not implemented |
| `DeleteDBInstance` | ✅ Implemented |
| `DeleteDBInstanceAutomatedBackup` | ❌ Not implemented |
| `DeleteDBParameterGroup` | ✅ Implemented |
| `DeleteDBProxy` | ❌ Not implemented |
| `DeleteDBProxyEndpoint` | ❌ Not implemented |
| `DeleteDBSecurityGroup` | ❌ Not implemented |
| `DeleteDBShardGroup` | ❌ Not implemented |
| `DeleteDBSnapshot` | ✅ Implemented |
| `DeleteDBSubnetGroup` | ✅ Implemented |
| `DeleteEventSubscription` | ❌ Not implemented |
| `DeleteGlobalCluster` | ❌ Not implemented |
| `DeleteIntegration` | ❌ Not implemented |
| `DeleteOptionGroup` | 🚫 Not supported |
| `DeleteTenantDatabase` | ❌ Not implemented |
| `DeregisterDBProxyTargets` | ❌ Not implemented |
| `DescribeAccountAttributes` | ❌ Not implemented |
| `DescribeBlueGreenDeployments` | ❌ Not implemented |
| `DescribeCertificates` | ❌ Not implemented |
| `DescribeDBClusterAutomatedBackups` | ❌ Not implemented |
| `DescribeDBClusterBacktracks` | ❌ Not implemented |
| `DescribeDBClusterEndpoints` | ❌ Not implemented |
| `DescribeDBClusterParameterGroups` | ❌ Not implemented |
| `DescribeDBClusterParameters` | ❌ Not implemented |
| `DescribeDBClusterSnapshotAttributes` | ❌ Not implemented |
| `DescribeDBClusterSnapshots` | ❌ Not implemented |
| `DescribeDBClusters` | 🚫 Not supported |
| `DescribeDBEngineVersions` | ✅ Implemented |
| `DescribeDBInstanceAutomatedBackups` | ✅ Implemented |
| `DescribeDBInstances` | ✅ Implemented |
| `DescribeDBLogFiles` | ❌ Not implemented |
| `DescribeDBParameterGroups` | ✅ Implemented |
| `DescribeDBParameters` | ✅ Implemented |
| `DescribeDBProxies` | ❌ Not implemented |
| `DescribeDBProxyEndpoints` | ❌ Not implemented |
| `DescribeDBProxyTargetGroups` | ❌ Not implemented |
| `DescribeDBProxyTargets` | ❌ Not implemented |
| `DescribeDBRecommendations` | ❌ Not implemented |
| `DescribeDBSecurityGroups` | ❌ Not implemented |
| `DescribeDBShardGroups` | ❌ Not implemented |
| `DescribeDBSnapshotAttributes` | ❌ Not implemented |
| `DescribeDBSnapshotTenantDatabases` | ❌ Not implemented |
| `DescribeDBSnapshots` | ✅ Implemented |
| `DescribeDBSubnetGroups` | ✅ Implemented |
| `DescribeEngineDefaultClusterParameters` | ❌ Not implemented |
| `DescribeEngineDefaultParameters` | ❌ Not implemented |
| `DescribeEventCategories` | ❌ Not implemented |
| `DescribeEventSubscriptions` | ❌ Not implemented |
| `DescribeEvents` | ✅ Implemented |
| `DescribeExportTasks` | ❌ Not implemented |
| `DescribeGlobalClusters` | ❌ Not implemented |
| `DescribeIntegrations` | ❌ Not implemented |
| `DescribeOptionGroupOptions` | ❌ Not implemented |
| `DescribeOptionGroups` | 🚫 Not supported |
| `DescribeOrderableDBInstanceOptions` | ✅ Implemented |
| `DescribePendingMaintenanceActions` | ❌ Not implemented |
| `DescribeReservedDBInstances` | ❌ Not implemented |
| `DescribeReservedDBInstancesOfferings` | ❌ Not implemented |
| `DescribeSourceRegions` | ❌ Not implemented |
| `DescribeTenantDatabases` | ❌ Not implemented |
| `DescribeValidDBInstanceModifications` | ❌ Not implemented |
| `DisableHttpEndpoint` | ❌ Not implemented |
| `DownloadDBLogFilePortion` | ❌ Not implemented |
| `EnableHttpEndpoint` | ❌ Not implemented |
| `FailoverDBCluster` | 🚫 Not supported |
| `FailoverGlobalCluster` | ❌ Not implemented |
| `ListTagsForResource` | ✅ Implemented |
| `ModifyActivityStream` | ❌ Not implemented |
| `ModifyCertificates` | ❌ Not implemented |
| `ModifyCurrentDBClusterCapacity` | ❌ Not implemented |
| `ModifyCustomDBEngineVersion` | ❌ Not implemented |
| `ModifyDBCluster` | 🚫 Not supported |
| `ModifyDBClusterEndpoint` | ❌ Not implemented |
| `ModifyDBClusterParameterGroup` | ❌ Not implemented |
| `ModifyDBClusterSnapshotAttribute` | ❌ Not implemented |
| `ModifyDBInstance` | ✅ Implemented |
| `ModifyDBParameterGroup` | ✅ Implemented |
| `ModifyDBProxy` | ❌ Not implemented |
| `ModifyDBProxyEndpoint` | ❌ Not implemented |
| `ModifyDBProxyTargetGroup` | ❌ Not implemented |
| `ModifyDBRecommendation` | ❌ Not implemented |
| `ModifyDBShardGroup` | ❌ Not implemented |
| `ModifyDBSnapshot` | ❌ Not implemented |
| `ModifyDBSnapshotAttribute` | ❌ Not implemented |
| `ModifyDBSubnetGroup` | ❌ Not implemented |
| `ModifyEventSubscription` | ❌ Not implemented |
| `ModifyGlobalCluster` | ❌ Not implemented |
| `ModifyIntegration` | ❌ Not implemented |
| `ModifyOptionGroup` | 🚫 Not supported |
| `ModifyTenantDatabase` | ❌ Not implemented |
| `PromoteReadReplica` | 🚫 Not supported |
| `PromoteReadReplicaDBCluster` | ❌ Not implemented |
| `PurchaseReservedDBInstancesOffering` | ❌ Not implemented |
| `RebootDBCluster` | ❌ Not implemented |
| `RebootDBInstance` | ✅ Implemented |
| `RebootDBShardGroup` | ❌ Not implemented |
| `RegisterDBProxyTargets` | ❌ Not implemented |
| `RemoveFromGlobalCluster` | ❌ Not implemented |
| `RemoveRoleFromDBCluster` | ❌ Not implemented |
| `RemoveRoleFromDBInstance` | ❌ Not implemented |
| `RemoveSourceIdentifierFromSubscription` | ❌ Not implemented |
| `RemoveTagsFromResource` | ✅ Implemented |
| `ResetDBClusterParameterGroup` | ❌ Not implemented |
| `ResetDBParameterGroup` | ❌ Not implemented |
| `RestoreDBClusterFromS3` | ❌ Not implemented |
| `RestoreDBClusterFromSnapshot` | ❌ Not implemented |
| `RestoreDBClusterToPointInTime` | ❌ Not implemented |
| `RestoreDBInstanceFromDBSnapshot` | ✅ Implemented |
| `RestoreDBInstanceFromS3` | ❌ Not implemented |
| `RestoreDBInstanceToPointInTime` | 🚫 Not supported |
| `RevokeDBSecurityGroupIngress` | ❌ Not implemented |
| `StartActivityStream` | ❌ Not implemented |
| `StartDBCluster` | ❌ Not implemented |
| `StartDBInstance` | ✅ Implemented |
| `StartDBInstanceAutomatedBackupsReplication` | ❌ Not implemented |
| `StartExportTask` | ❌ Not implemented |
| `StopActivityStream` | ❌ Not implemented |
| `StopDBCluster` | ❌ Not implemented |
| `StopDBInstance` | ✅ Implemented |
| `StopDBInstanceAutomatedBackupsReplication` | ❌ Not implemented |
| `SwitchoverBlueGreenDeployment` | ❌ Not implemented |
| `SwitchoverGlobalCluster` | ❌ Not implemented |
| `SwitchoverReadReplica` | ❌ Not implemented |
| `AcknowledgeDBBootstrap` | 🔒 Outside the pinned model |
| `GetDBBootstrapConfig` | 🔒 Outside the pinned model |
| `PollDBCommands` | 🔒 Outside the pinned model |
| `RegisterDBInstance` | 🔒 Outside the pinned model |
| `SubmitDBStateChange` | 🔒 Outside the pinned model |
