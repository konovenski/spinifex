---
title: "ECR API Coverage"
seoTitle: "Elastic Container Registry API Coverage — Spinifex Docs"
description: "Every operation in the Amazon ECR API model and whether Spinifex implements it, alongside the OCI distribution endpoint that carries the image layers."
category: "Coverage"
sections:
  - overview
tags:
  - aws
  - compatibility
  - coverage
  - ecr
  - containers
  - registry
---

# ECR API Coverage

## Overview

Spinifex implements **19 of the 47** operations (**40.4%**) in the ECR `2015-09-21` API model.

### Two endpoints, one registry

Repository metadata is served over the AWS API on the gateway endpoint. Image data moves over the OCI Distribution `/v2/` endpoint on that same host, authenticated by the bearer token `GetAuthorizationToken` mints for `docker login`.

The split explains the largest group of unimplemented operations below. The layer-transfer operations — `BatchCheckLayerAvailability`, `InitiateLayerUpload`, `UploadLayerPart`, `CompleteLayerUpload` and `GetDownloadUrlForLayer` — are registered stubs because the `/v2/` endpoint carries that traffic instead. A client using `docker` or any OCI-compatible tool never calls them.

Registry replication is absent as well: a deployment is a single registry, with no cross-region peer to replicate to.

### Operations

| Operation | Status |
|---|---|
| `BatchCheckLayerAvailability` | 🟡 Stub |
| `BatchDeleteImage` | ✅ Implemented |
| `BatchGetImage` | ✅ Implemented |
| `BatchGetRepositoryScanningConfiguration` | ⛔ Not applicable [1](#notes) |
| `CompleteLayerUpload` | 🟡 Stub |
| `CreatePullThroughCacheRule` | ❌ Not implemented |
| `CreateRepository` | ✅ Implemented |
| `CreateRepositoryCreationTemplate` | ❌ Not implemented |
| `DeleteLifecyclePolicy` | ✅ Implemented |
| `DeletePullThroughCacheRule` | ❌ Not implemented |
| `DeleteRegistryPolicy` | ❌ Not implemented |
| `DeleteRepository` | ✅ Implemented |
| `DeleteRepositoryCreationTemplate` | ❌ Not implemented |
| `DeleteRepositoryPolicy` | ✅ Implemented |
| `DescribeImageReplicationStatus` | ❌ Not implemented |
| `DescribeImageScanFindings` | ⛔ Not applicable [1](#notes) |
| `DescribeImages` | ✅ Implemented |
| `DescribePullThroughCacheRules` | ❌ Not implemented |
| `DescribeRegistry` | 🟡 Stub |
| `DescribeRepositories` | ✅ Implemented |
| `DescribeRepositoryCreationTemplates` | ❌ Not implemented |
| `GetAuthorizationToken` | ✅ Implemented |
| `GetDownloadUrlForLayer` | 🟡 Stub |
| `GetLifecyclePolicy` | ✅ Implemented |
| `GetLifecyclePolicyPreview` | ✅ Implemented |
| `GetRegistryPolicy` | 🟡 Stub |
| `GetRegistryScanningConfiguration` | ⛔ Not applicable [1](#notes) |
| `GetRepositoryPolicy` | ✅ Implemented |
| `InitiateLayerUpload` | 🟡 Stub |
| `ListImages` | ✅ Implemented |
| `ListTagsForResource` | ✅ Implemented |
| `PutImage` | ✅ Implemented |
| `PutImageScanningConfiguration` | ⛔ Not applicable [1](#notes) |
| `PutImageTagMutability` | ✅ Implemented |
| `PutLifecyclePolicy` | ✅ Implemented |
| `PutRegistryPolicy` | 🟡 Stub |
| `PutRegistryScanningConfiguration` | ⛔ Not applicable [1](#notes) |
| `PutReplicationConfiguration` | 🟡 Stub |
| `SetRepositoryPolicy` | ✅ Implemented |
| `StartImageScan` | ⛔ Not applicable [1](#notes) |
| `StartLifecyclePolicyPreview` | ✅ Implemented |
| `TagResource` | 🟡 Stub |
| `UntagResource` | 🟡 Stub |
| `UpdatePullThroughCacheRule` | ❌ Not implemented |
| `UpdateRepositoryCreationTemplate` | ❌ Not implemented |
| `UploadLayerPart` | 🟡 Stub |
| `ValidatePullThroughCacheRule` | ❌ Not implemented |
| `GetImageScanningConfiguration` | 🔒 Outside the pinned model |
| `ListRepositories` | 🔒 Outside the pinned model |
| `ReplicateImage` | 🔒 Outside the pinned model |

### Notes

1. Image scanning is an AWS-hosted vulnerability service with no equivalent here, so a finding set would always be empty.
