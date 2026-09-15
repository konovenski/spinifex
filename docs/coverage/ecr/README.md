---
title: "ECR API Coverage"
seoTitle: "Elastic Container Registry API Coverage — Spinifex Docs"
description: "The Amazon ECR API operations Spinifex implements, alongside the OCI distribution endpoint that carries the image layers, and those the platform will not."
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

Spinifex implements **19 operations** in the ECR `2015-09-21` API model.

### Two endpoints, one registry

Repository metadata is served over the AWS API on the gateway endpoint. Image data moves over the OCI Distribution `/v2/` endpoint on that same host, authenticated by the bearer token `GetAuthorizationToken` mints for `docker login`.

The split explains the stubs below. The layer-transfer operations — `BatchCheckLayerAvailability`, `InitiateLayerUpload`, `UploadLayerPart`, `CompleteLayerUpload` and `GetDownloadUrlForLayer` — are registered stubs because the `/v2/` endpoint carries that traffic instead. A client using `docker` or any OCI-compatible tool never calls them.

Registry replication is a stub for the same kind of reason: a deployment is a single registry, with no cross-region peer to replicate to.

### Operations

| Operation | Status |
|---|---|
| `BatchCheckLayerAvailability` | 🟡 Stub |
| `BatchDeleteImage` | ✅ Implemented |
| `BatchGetImage` | ✅ Implemented |
| `CompleteLayerUpload` | 🟡 Stub |
| `CreateRepository` | ✅ Implemented |
| `DeleteLifecyclePolicy` | ✅ Implemented |
| `DeleteRepository` | ✅ Implemented |
| `DeleteRepositoryPolicy` | ✅ Implemented |
| `DescribeImages` | ✅ Implemented |
| `DescribeRegistry` | 🟡 Stub |
| `DescribeRepositories` | ✅ Implemented |
| `GetAuthorizationToken` | ✅ Implemented |
| `GetDownloadUrlForLayer` | 🟡 Stub |
| `GetLifecyclePolicy` | ✅ Implemented |
| `GetLifecyclePolicyPreview` | ✅ Implemented |
| `GetRegistryPolicy` | 🟡 Stub |
| `GetRepositoryPolicy` | ✅ Implemented |
| `InitiateLayerUpload` | 🟡 Stub |
| `ListImages` | ✅ Implemented |
| `ListTagsForResource` | ✅ Implemented |
| `PutImage` | ✅ Implemented |
| `PutImageTagMutability` | ✅ Implemented |
| `PutLifecyclePolicy` | ✅ Implemented |
| `PutRegistryPolicy` | 🟡 Stub |
| `PutReplicationConfiguration` | 🟡 Stub |
| `SetRepositoryPolicy` | ✅ Implemented |
| `StartLifecyclePolicyPreview` | ✅ Implemented |
| `TagResource` | 🟡 Stub |
| `UntagResource` | 🟡 Stub |
| `UploadLayerPart` | 🟡 Stub |
| `GetImageScanningConfiguration` | 🔒 Outside the pinned model |
| `ListRepositories` | 🔒 Outside the pinned model |
| `ReplicateImage` | 🔒 Outside the pinned model |

### Not applicable

These operations describe AWS-hosted features this platform does not offer.

| Operation | Reason |
|---|---|
| `BatchGetRepositoryScanningConfiguration` | Image scanning is an AWS-hosted vulnerability service with no equivalent here, so a finding set would always be empty. |
| `DescribeImageScanFindings` | Image scanning is an AWS-hosted vulnerability service with no equivalent here, so a finding set would always be empty. |
| `GetRegistryScanningConfiguration` | Image scanning is an AWS-hosted vulnerability service with no equivalent here, so a finding set would always be empty. |
| `PutImageScanningConfiguration` | Image scanning is an AWS-hosted vulnerability service with no equivalent here, so a finding set would always be empty. |
| `PutRegistryScanningConfiguration` | Image scanning is an AWS-hosted vulnerability service with no equivalent here, so a finding set would always be empty. |
| `StartImageScan` | Image scanning is an AWS-hosted vulnerability service with no equivalent here, so a finding set would always be empty. |
