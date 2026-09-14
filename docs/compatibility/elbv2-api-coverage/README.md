---
title: "ELBv2 API Coverage"
seoTitle: "Elastic Load Balancing v2 API Coverage — Spinifex Docs"
description: "Every operation in the Elastic Load Balancing v2 API model and whether Spinifex implements it, for both the Application and Network Load Balancers it serves."
category: "Coverage"
sections:
  - overview
tags:
  - aws
  - compatibility
  - coverage
  - elbv2
  - load balancing
  - networking
---

# ELBv2 API Coverage

## Overview

Spinifex implements **33 of the 46** operations in the ELBv2 `2015-12-01` API model, as pinned in `aws-sdk-go v1.55.8` — **71.7%**.

### Two data planes

The data plane is a system-managed load balancer VM, launched automatically during `CreateLoadBalancer`. Application Load Balancers run HAProxy for the L7 surface — rules, fixed responses and redirects over HTTP and HTTPS. Network Load Balancers run nginx `stream` for the L4 surface — TCP, UDP, TLS and TCP_UDP — because HAProxy cannot load-balance UDP. The agent selects the engine from the configuration the control plane delivers, so the choice follows the load balancer type and is not separately configurable.

### Operations

| Status | Meaning |
|---|---|
| ✅ Implemented | A modelled operation bound to a real handler. |
| 🟡 Stub | A registered handler that answers with a fixed or empty result. |
| 🚫 Not supported | A registered handler that deliberately refuses, so a client sees "not offered" rather than an unknown action. |
| ❌ Not implemented | Modelled by AWS, not registered by Spinifex. |
| 🔒 Outside the pinned model | Registered by Spinifex but absent from the pinned model — an internal route, not a tenant-callable AWS action. |

| Operation | Status |
|---|---|
| `AddListenerCertificates` | ✅ Implemented |
| `AddTags` | ✅ Implemented |
| `AddTrustStoreRevocations` | ❌ Not implemented |
| `CreateListener` | ✅ Implemented |
| `CreateLoadBalancer` | ✅ Implemented |
| `CreateRule` | ✅ Implemented |
| `CreateTargetGroup` | ✅ Implemented |
| `CreateTrustStore` | ❌ Not implemented |
| `DeleteListener` | ✅ Implemented |
| `DeleteLoadBalancer` | ✅ Implemented |
| `DeleteRule` | ✅ Implemented |
| `DeleteSharedTrustStoreAssociation` | ❌ Not implemented |
| `DeleteTargetGroup` | ✅ Implemented |
| `DeleteTrustStore` | ❌ Not implemented |
| `DeregisterTargets` | ✅ Implemented |
| `DescribeAccountLimits` | ❌ Not implemented |
| `DescribeListenerCertificates` | ✅ Implemented |
| `DescribeListeners` | ✅ Implemented |
| `DescribeLoadBalancerAttributes` | ✅ Implemented |
| `DescribeLoadBalancers` | ✅ Implemented |
| `DescribeRules` | ✅ Implemented |
| `DescribeSSLPolicies` | ✅ Implemented |
| `DescribeTags` | ✅ Implemented |
| `DescribeTargetGroupAttributes` | ✅ Implemented |
| `DescribeTargetGroups` | ✅ Implemented |
| `DescribeTargetHealth` | ✅ Implemented |
| `DescribeTrustStoreAssociations` | ❌ Not implemented |
| `DescribeTrustStoreRevocations` | ❌ Not implemented |
| `DescribeTrustStores` | ❌ Not implemented |
| `GetResourcePolicy` | ❌ Not implemented |
| `GetTrustStoreCaCertificatesBundle` | ❌ Not implemented |
| `GetTrustStoreRevocationContent` | ❌ Not implemented |
| `ModifyListener` | ✅ Implemented |
| `ModifyLoadBalancerAttributes` | ✅ Implemented |
| `ModifyRule` | ✅ Implemented |
| `ModifyTargetGroup` | ✅ Implemented |
| `ModifyTargetGroupAttributes` | ✅ Implemented |
| `ModifyTrustStore` | ❌ Not implemented |
| `RegisterTargets` | ✅ Implemented |
| `RemoveListenerCertificates` | ✅ Implemented |
| `RemoveTags` | ✅ Implemented |
| `RemoveTrustStoreRevocations` | ❌ Not implemented |
| `SetIpAddressType` | ✅ Implemented |
| `SetRulePriorities` | ✅ Implemented |
| `SetSecurityGroups` | ✅ Implemented |
| `SetSubnets` | ✅ Implemented |
| `DescribeListenerAttributes` | 🔒 Outside the pinned model |
| `GetLBConfig` | 🔒 Outside the pinned model |
| `LBAgentHeartbeat` | 🔒 Outside the pinned model |
| `ModifyListenerAttributes` | 🔒 Outside the pinned model |
