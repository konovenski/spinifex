---
title: "ELBv2 API Coverage"
seoTitle: "Elastic Load Balancing v2 API Coverage — Spinifex Docs"
description: "The Elastic Load Balancing v2 API operations Spinifex implements, for both the Application and Network Load Balancers it serves, listeners and target groups."
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

Spinifex implements **34 operations** in the ELBv2 `2015-12-01` API model.

### Two data planes

The data plane is a system-managed load balancer VM, launched automatically during `CreateLoadBalancer`. Application Load Balancers run HAProxy for the L7 surface — rules, fixed responses and redirects over HTTP and HTTPS. Network Load Balancers run nginx `stream` for the L4 surface — TCP, UDP, TLS and TCP_UDP — because HAProxy cannot load-balance UDP. The agent selects the engine from the configuration the control plane delivers, so the choice follows the load balancer type and is not separately configurable.

### Operations

| Operation | Status |
|---|---|
| `AddListenerCertificates` | ✅ Implemented |
| `AddTags` | ✅ Implemented |
| `CreateListener` | ✅ Implemented |
| `CreateLoadBalancer` | ✅ Implemented |
| `CreateRule` | ✅ Implemented |
| `CreateTargetGroup` | ✅ Implemented |
| `DeleteListener` | ✅ Implemented |
| `DeleteLoadBalancer` | ✅ Implemented |
| `DeleteRule` | ✅ Implemented |
| `DeleteTargetGroup` | ✅ Implemented |
| `DeregisterTargets` | ✅ Implemented |
| `DescribeAccountLimits` | ✅ Implemented |
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
| `ModifyListener` | ✅ Implemented |
| `ModifyLoadBalancerAttributes` | ✅ Implemented |
| `ModifyRule` | ✅ Implemented |
| `ModifyTargetGroup` | ✅ Implemented |
| `ModifyTargetGroupAttributes` | ✅ Implemented |
| `RegisterTargets` | ✅ Implemented |
| `RemoveListenerCertificates` | ✅ Implemented |
| `RemoveTags` | ✅ Implemented |
| `SetIpAddressType` | ✅ Implemented |
| `SetRulePriorities` | ✅ Implemented |
| `SetSecurityGroups` | ✅ Implemented |
| `SetSubnets` | ✅ Implemented |
| `DescribeListenerAttributes` | 🔒 Outside the pinned model |
| `ModifyListenerAttributes` | 🔒 Outside the pinned model |
