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

Spinifex implements **26 of the 162** operations (**16.0%**) in the RDS `2014-10-31` API model.

### Engines

Spinifex offers PostgreSQL and MariaDB. Each DB instance is one dedicated system-owned VM running the engine directly, launched from a platform AMI and hidden from the customer's EC2 API. `Engine` is fixed at create: there is no in-place engine change, no cross-engine snapshot restore and no migration between the two.

`mysql` is not an accepted engine and is not an alias for `mariadb`. MariaDB is offered under its own AWS engine name, exactly as AWS RDS offers it, so a client — including Terraform's `aws_db_instance` — must set `engine = "mariadb"`.

Engine versions are pinned per engine. An `EngineVersion` naming anything but the pin is rejected, including a narrower minor version, because the AMI makes no promise about which minor it carries.

### The endpoint is private

The engine is reached over a customer-account ENI injected into a subnet of the DB subnet group, so the endpoint is reachable from inside the VPC only. The DB VM has other NICs that no customer security group governs, and the engine binds none of them — the port is not open there at all, rather than open and gated.

### TLS is required by default

Both engines enforce encrypted connections by default: `rds.force_ssl` on PostgreSQL and `require_secure_transport` on MariaDB. For MariaDB this is a deliberate divergence from AWS, which leaves it off. Both are boolean, modifiable and dynamic, so setting either to `0` in a parameter group restores plaintext without a reboot.

### Rejected parameters

A parameter whose omission would create a false safety, security or availability guarantee is rejected with `InvalidParameterValue` rather than silently dropped.

| Parameter | Why it is rejected |
|-----------|--------------------|
| `MultiAZ=true` | Single-AZ platform; a standby would not exist |
| `PubliclyAccessible=true` | The endpoint is a private VPC address |
| `StorageEncrypted=false` | Unencrypted storage is not offered |
| `EnableIAMDatabaseAuthentication` | IAM database authentication is not implemented |
| `Iops`, `StorageThroughput`, `StorageType` ≠ `gp3` | Provisioned performance classes are not implemented |
| `KmsKeyId`, `TdeCredentialArn` | Storage is encrypted with the cluster key, not a customer-managed one |
| `AvailabilityZone` | The platform exposes a single zone |
| `AvailabilityZoneGroup` (orderable options) | It selects a zone or local-zone group, and naming a zone is already refused |
| `DBSecurityGroups` | EC2-Classic security groups — use `VpcSecurityGroupIds` |
| `DBClusterIdentifier`, `DBClusterSnapshotIdentifier` | Clustered engines are not offered |
| `EnableCloudwatchLogsExports` | Log export is not implemented |
| `EngineVersion` other than the engine's pin, `Engine` on modify | No in-place engine or version change |
| `Engine=mysql` (and Aurora engines) | Oracle MySQL is not offered; `mariadb` is a distinct engine, not an alias for it |
| `NewDBInstanceIdentifier` | The identifier is the DNS label and the KV key |
| `DBPortNumber`, `DBSubnetGroupName` on modify | Both would move the endpoint |
| `MaxAllocatedStorage` | Storage autoscaling is not implemented |
| `ManageMasterUserPassword`, `RotateMasterUserPassword` | Secrets Manager integration is not offered |
| `CACertificateIdentifier` | The serving certificate is minted from the cluster CA |
| `Domain`, `DomainFqdn` | Active Directory domain join is not offered |
| `OptionGroupName` | Option groups are not offered |
| `CustomIamInstanceProfile` | The DB VM's instance profile is platform-owned |
| `EnableCustomerOwnedIp` | An Outposts feature |
| `ForceFailover` (reboot) | No standby to fail over to |
| `DBSnapshotIdentifier` (stop) | Snapshot-on-stop is not implemented |

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
| `CreateDBCluster` | ⛔ Not applicable [1](#notes) |
| `CreateDBClusterEndpoint` | ❌ Not implemented |
| `CreateDBClusterParameterGroup` | ❌ Not implemented |
| `CreateDBClusterSnapshot` | ❌ Not implemented |
| `CreateDBInstance` | ✅ Implemented |
| `CreateDBInstanceReadReplica` | ⛔ Not applicable [4](#notes) |
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
| `CreateOptionGroup` | ⛔ Not applicable [2](#notes) |
| `CreateTenantDatabase` | ❌ Not implemented |
| `DeleteBlueGreenDeployment` | ❌ Not implemented |
| `DeleteCustomDBEngineVersion` | ❌ Not implemented |
| `DeleteDBCluster` | ⛔ Not applicable [1](#notes) |
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
| `DeleteOptionGroup` | ⛔ Not applicable [2](#notes) |
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
| `DescribeDBClusters` | ⛔ Not applicable [1](#notes) |
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
| `DescribeOptionGroups` | ⛔ Not applicable [2](#notes) |
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
| `FailoverDBCluster` | ⛔ Not applicable [1](#notes) |
| `FailoverGlobalCluster` | ❌ Not implemented |
| `ListTagsForResource` | ✅ Implemented |
| `ModifyActivityStream` | ❌ Not implemented |
| `ModifyCertificates` | ❌ Not implemented |
| `ModifyCurrentDBClusterCapacity` | ❌ Not implemented |
| `ModifyCustomDBEngineVersion` | ❌ Not implemented |
| `ModifyDBCluster` | ⛔ Not applicable [1](#notes) |
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
| `ModifyOptionGroup` | ⛔ Not applicable [2](#notes) |
| `ModifyTenantDatabase` | ❌ Not implemented |
| `PromoteReadReplica` | ⛔ Not applicable [4](#notes) |
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
| `RestoreDBInstanceToPointInTime` | ⛔ Not applicable [3](#notes) |
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

### Notes

1. Aurora and Multi-AZ clusters are not offered; a DB instance here is a single VM running the engine directly.
2. Option groups configure engine add-ons for engines this platform does not offer, such as Oracle and SQL Server.
3. Point-in-time restore needs continuous transaction-log archival, which the backup path does not keep.
4. Replication between instances is not offered, so there is no replica to create or promote.
