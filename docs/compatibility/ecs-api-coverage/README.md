---
title: "ECS API Coverage"
seoTitle: "Elastic Container Service API Coverage — Spinifex Docs"
description: "Every operation in the Amazon ECS API model and whether Spinifex implements it, covering clusters, services, tasks and container instances on each build."
category: "Coverage"
sections:
  - overview
tags:
  - aws
  - compatibility
  - coverage
  - ecs
  - containers
  - orchestration
---

# ECS API Coverage

## Overview

Spinifex implements **31 of the 56** operations (**55.4%**) in the ECS `2014-11-13` API model.

### No Fargate

Clusters, services and tasks run on EC2 container instances. There is no Fargate launch type, and no operation below provides one.

### Operations

| Operation | Status |
|---|---|
| `CreateCapacityProvider` | ✅ Implemented |
| `CreateCluster` | ✅ Implemented |
| `CreateService` | ✅ Implemented |
| `CreateTaskSet` | ❌ Not implemented |
| `DeleteAccountSetting` | ❌ Not implemented |
| `DeleteAttributes` | ❌ Not implemented |
| `DeleteCapacityProvider` | ✅ Implemented |
| `DeleteCluster` | ✅ Implemented |
| `DeleteService` | ✅ Implemented |
| `DeleteTaskDefinitions` | ❌ Not implemented |
| `DeleteTaskSet` | ❌ Not implemented |
| `DeregisterContainerInstance` | ✅ Implemented |
| `DeregisterTaskDefinition` | ✅ Implemented |
| `DescribeCapacityProviders` | ✅ Implemented |
| `DescribeClusters` | ✅ Implemented |
| `DescribeContainerInstances` | ✅ Implemented |
| `DescribeServices` | ✅ Implemented |
| `DescribeTaskDefinition` | ✅ Implemented |
| `DescribeTaskSets` | ❌ Not implemented |
| `DescribeTasks` | ✅ Implemented |
| `DiscoverPollEndpoint` | ❌ Not implemented |
| `ExecuteCommand` | ❌ Not implemented |
| `GetTaskProtection` | ❌ Not implemented |
| `ListAccountSettings` | 🟡 Stub |
| `ListAttributes` | ❌ Not implemented |
| `ListClusters` | ✅ Implemented |
| `ListContainerInstances` | ✅ Implemented |
| `ListServices` | ✅ Implemented |
| `ListServicesByNamespace` | 🟡 Stub |
| `ListTagsForResource` | ✅ Implemented |
| `ListTaskDefinitionFamilies` | 🟡 Stub |
| `ListTaskDefinitions` | ✅ Implemented |
| `ListTasks` | ✅ Implemented |
| `PutAccountSetting` | 🟡 Stub |
| `PutAccountSettingDefault` | ❌ Not implemented |
| `PutAttributes` | ❌ Not implemented |
| `PutClusterCapacityProviders` | ✅ Implemented |
| `RegisterContainerInstance` | ✅ Implemented |
| `RegisterTaskDefinition` | ✅ Implemented |
| `RunTask` | ✅ Implemented |
| `StartTask` | ✅ Implemented |
| `StopTask` | ✅ Implemented |
| `SubmitAttachmentStateChanges` | ❌ Not implemented |
| `SubmitContainerStateChange` | ❌ Not implemented |
| `SubmitTaskStateChange` | ✅ Implemented |
| `TagResource` | ✅ Implemented |
| `UntagResource` | ✅ Implemented |
| `UpdateCapacityProvider` | ❌ Not implemented |
| `UpdateCluster` | 🟡 Stub |
| `UpdateClusterSettings` | ❌ Not implemented |
| `UpdateContainerAgent` | ❌ Not implemented |
| `UpdateContainerInstancesState` | ✅ Implemented |
| `UpdateService` | ✅ Implemented |
| `UpdateServicePrimaryTaskSet` | ❌ Not implemented |
| `UpdateTaskProtection` | ❌ Not implemented |
| `UpdateTaskSet` | ❌ Not implemented |
