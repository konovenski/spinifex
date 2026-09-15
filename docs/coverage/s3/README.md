---
title: "S3 API Coverage"
seoTitle: "Amazon S3 API Coverage on Spinifex — Spinifex Docs"
description: "The Amazon S3 API operations Predastore serves on the platform's S3 endpoint, alongside the AWS-hosted features it does not offer and why each is absent."
category: "Coverage"
sections:
  - overview
tags:
  - aws
  - compatibility
  - coverage
  - s3
  - storage
  - objects
---

# S3 API Coverage

## Overview

Predastore implements **19 operations** in the S3 `2006-03-01` API model.

### Predastore serves this endpoint

S3 is the one surface the AWS gateway does not answer itself. Object storage runs on [Predastore](https://github.com/mulgadc/predastore), which serves the S3 REST API directly over its own endpoint.

### Routed is not the same as conforming

This page says an operation is routed to a handler. It does not say the handler's behaviour matches S3 in every case.

That behaviour is measured separately, against the `ceph/s3-tests` suite Ceph RGW, MinIO and Garage are all validated with, and the results are published in Predastore's [S3 compatibility report](https://github.com/mulgadc/predastore/blob/dev/docs/S3-COMPATIBILITY.md).

### Operations

| Operation | Status |
|---|---|
| `AbortMultipartUpload` | ✅ Implemented |
| `CompleteMultipartUpload` | ✅ Implemented |
| `CopyObject` | ✅ Implemented |
| `CreateBucket` | ✅ Implemented |
| `CreateMultipartUpload` | ✅ Implemented |
| `CreateSession` | ⛔ Not applicable [2](#notes) |
| `DeleteBucket` | ✅ Implemented |
| `DeleteObject` | ✅ Implemented |
| `DeleteObjects` | ✅ Implemented |
| `GetBucketAccelerateConfiguration` | ⛔ Not applicable [1](#notes) |
| `GetBucketRequestPayment` | ⛔ Not applicable [4](#notes) |
| `GetObject` | ✅ Implemented |
| `HeadBucket` | ✅ Implemented |
| `HeadObject` | ✅ Implemented |
| `ListBuckets` | ✅ Implemented |
| `ListMultipartUploads` | ✅ Implemented |
| `ListObjects` | ✅ Implemented |
| `ListObjectsV2` | ✅ Implemented |
| `ListParts` | ✅ Implemented |
| `PutBucketAccelerateConfiguration` | ⛔ Not applicable [1](#notes) |
| `PutBucketRequestPayment` | ⛔ Not applicable [4](#notes) |
| `PutObject` | ✅ Implemented |
| `UploadPart` | ✅ Implemented |
| `UploadPartCopy` | ✅ Implemented |
| `WriteGetObjectResponse` | ⛔ Not applicable [3](#notes) |

### Notes

1. Transfer acceleration routes uploads over AWS edge locations, which an on-premise deployment has none of.
2. Sessions authenticate S3 Express One Zone directory buckets, a storage class this platform does not offer.
3. S3 Object Lambda rewrites a response from a Lambda function, and Lambda is not offered.
4. Requester Pays shifts transfer charges to the caller, which needs AWS billing behind it.
