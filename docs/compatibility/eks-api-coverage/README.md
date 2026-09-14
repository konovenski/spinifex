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

Spinifex implements **34 of the 56** operations in the EKS `2017-11-01` API model, as pinned in `aws-sdk-go v1.55.8` — **60.7%**.

### Routing

EKS is REST-JSON, dispatched by HTTP method and path rather than by an action parameter, so the gateway's routing table is the authoritative list of what it serves. Clusters run K3s on control-plane VMs; nodegroups are EC2 instances joined to them.

### Internal routes

Four routes on this surface are internal, not AWS actions, and appear in the table below as registered outside the pinned model:

- `PublishInternal` — the control-plane VM broker, relaying bootstrap and state posts onto the message bus.
- `WebhookTokenReview` — the token webhook posting bearer tokens for host-side resolution.
- `ListInternalAddons` — the on-VM add-on sync agent fetching its staged manifests.
- `GetRecoveryDirective` — the on-VM recovery agent reading its per-member directive at boot.

None is callable with an ordinary tenant credential. Each names the target account in the path rather than taking it from the credential, so each requires a control-plane instance-role session bound to the named cluster; a broad `eks:*` grant does not reach them.

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
| `GetRecoveryDirective` | 🔒 Outside the pinned model |
| `ListInternalAddons` | 🔒 Outside the pinned model |
| `PublishInternal` | 🔒 Outside the pinned model |
| `WebhookTokenReview` | 🔒 Outside the pinned model |
