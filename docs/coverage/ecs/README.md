---
title: "ECS API Coverage"
seoTitle: "Elastic Container Service API Coverage — Spinifex Docs"
description: "The Amazon ECS API operations Spinifex implements, covering clusters, services, tasks and container instances, with those the platform does not offer."
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

Spinifex implements **31 operations** in the ECS `2014-11-13` API model.

### No Fargate

Clusters, services and tasks run on EC2 container instances. There is no Fargate launch type, and no operation below provides one.

### Operations

| Operation | Status |
|---|---|
| `CreateCapacityProvider` | ✅ Implemented |
| `CreateCluster` | ✅ Implemented |
| `CreateService` | ✅ Implemented |
| `DeleteCapacityProvider` | ✅ Implemented |
| `DeleteCluster` | ✅ Implemented |
| `DeleteService` | ✅ Implemented |
| `DeregisterContainerInstance` | ✅ Implemented |
| `DeregisterTaskDefinition` | ✅ Implemented |
| `DescribeCapacityProviders` | ✅ Implemented |
| `DescribeClusters` | ✅ Implemented |
| `DescribeContainerInstances` | ✅ Implemented |
| `DescribeServices` | ✅ Implemented |
| `DescribeTaskDefinition` | ✅ Implemented |
| `DescribeTasks` | ✅ Implemented |
| `ListAccountSettings` | 🟡 Stub |
| `ListClusters` | ✅ Implemented |
| `ListContainerInstances` | ✅ Implemented |
| `ListServices` | ✅ Implemented |
| `ListServicesByNamespace` | 🟡 Stub |
| `ListTagsForResource` | ✅ Implemented |
| `ListTaskDefinitionFamilies` | 🟡 Stub |
| `ListTaskDefinitions` | ✅ Implemented |
| `ListTasks` | ✅ Implemented |
| `PutAccountSetting` | 🟡 Stub |
| `PutClusterCapacityProviders` | ✅ Implemented |
| `RegisterContainerInstance` | ✅ Implemented |
| `RegisterTaskDefinition` | ✅ Implemented |
| `RunTask` | ✅ Implemented |
| `StartTask` | ✅ Implemented |
| `StopTask` | ✅ Implemented |
| `SubmitTaskStateChange` | ✅ Implemented |
| `TagResource` | ✅ Implemented |
| `UntagResource` | ✅ Implemented |
| `UpdateCluster` | 🟡 Stub |
| `UpdateContainerInstancesState` | ✅ Implemented |
| `UpdateService` | ✅ Implemented |
