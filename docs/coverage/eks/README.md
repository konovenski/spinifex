---
title: "EKS API Coverage"
seoTitle: "Elastic Kubernetes Service API Coverage — Spinifex Docs"
description: "Every operation in the Amazon EKS API model and whether Spinifex implements it, covering clusters, nodegroups, add-ons and access entries on each build."
category: "Coverage"
sections:
  - overview
tags:
  - aws
  - compatibility
  - coverage
  - eks
  - kubernetes
  - containers
---

# EKS API Coverage

## Overview

Spinifex implements **34 of the 56** operations (**60.7%**) in the EKS `2017-11-01` API model.

### Operations

| Operation | Status |
|---|---|
| `AssociateAccessPolicy` | ✅ Implemented |
| `AssociateEncryptionConfig` | ❌ Not implemented |
| `AssociateIdentityProviderConfig` | ✅ Implemented |
| `CreateAccessEntry` | ✅ Implemented |
| `CreateAddon` | ✅ Implemented |
| `CreateCluster` | ✅ Implemented |
| `CreateEksAnywhereSubscription` | ❌ Not implemented |
| `CreateFargateProfile` | ❌ Not implemented |
| `CreateNodegroup` | ✅ Implemented |
| `CreatePodIdentityAssociation` | ❌ Not implemented |
| `DeleteAccessEntry` | ✅ Implemented |
| `DeleteAddon` | ✅ Implemented |
| `DeleteCluster` | ✅ Implemented |
| `DeleteEksAnywhereSubscription` | ❌ Not implemented |
| `DeleteFargateProfile` | ❌ Not implemented |
| `DeleteNodegroup` | ✅ Implemented |
| `DeletePodIdentityAssociation` | ❌ Not implemented |
| `DeregisterCluster` | ❌ Not implemented |
| `DescribeAccessEntry` | ✅ Implemented |
| `DescribeAddon` | ✅ Implemented |
| `DescribeAddonConfiguration` | ❌ Not implemented |
| `DescribeAddonVersions` | ✅ Implemented |
| `DescribeCluster` | ✅ Implemented |
| `DescribeEksAnywhereSubscription` | ❌ Not implemented |
| `DescribeFargateProfile` | ❌ Not implemented |
| `DescribeIdentityProviderConfig` | ✅ Implemented |
| `DescribeInsight` | ❌ Not implemented |
| `DescribeNodegroup` | ✅ Implemented |
| `DescribePodIdentityAssociation` | ❌ Not implemented |
| `DescribeUpdate` | ❌ Not implemented |
| `DisassociateAccessPolicy` | ✅ Implemented |
| `DisassociateIdentityProviderConfig` | ✅ Implemented |
| `ListAccessEntries` | ✅ Implemented |
| `ListAccessPolicies` | ✅ Implemented |
| `ListAddons` | ✅ Implemented |
| `ListAssociatedAccessPolicies` | ✅ Implemented |
| `ListClusters` | ✅ Implemented |
| `ListEksAnywhereSubscriptions` | ❌ Not implemented |
| `ListFargateProfiles` | ❌ Not implemented |
| `ListIdentityProviderConfigs` | ✅ Implemented |
| `ListInsights` | ❌ Not implemented |
| `ListNodegroups` | ✅ Implemented |
| `ListPodIdentityAssociations` | ❌ Not implemented |
| `ListTagsForResource` | ✅ Implemented |
| `ListUpdates` | ❌ Not implemented |
| `RegisterCluster` | ❌ Not implemented |
| `TagResource` | ✅ Implemented |
| `UntagResource` | ✅ Implemented |
| `UpdateAccessEntry` | ✅ Implemented |
| `UpdateAddon` | ✅ Implemented |
| `UpdateClusterConfig` | ✅ Implemented |
| `UpdateClusterVersion` | ✅ Implemented |
| `UpdateEksAnywhereSubscription` | ❌ Not implemented |
| `UpdateNodegroupConfig` | ✅ Implemented |
| `UpdateNodegroupVersion` | ✅ Implemented |
| `UpdatePodIdentityAssociation` | ❌ Not implemented |
