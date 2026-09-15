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

S3 is the one surface the AWS gateway does not answer itself. Object storage runs on [Predastore](https://github.com/mulgadc/predastore), which serves the S3 REST API directly over its own TLS endpoint, with Reed-Solomon erasure coding and AES-256-GCM at rest behind it.

S3 selects an operation by method, path, query parameter and header rather than by an action name, so several operations share one route. The table below is generated from the route table Predastore builds its router from: `GET /{bucket}` answers both `ListObjects` and `ListObjectsV2`, and `PUT /{bucket}/{key}` splits four ways on `?partNumber` and `x-amz-copy-source`.

### Routed is not the same as conforming

This page says an operation is routed to a handler. It does not say the handler's behaviour matches S3 in every case, and for several operations it does not: ETag is not the body MD5, `Marker` is ignored on a v1 listing, and `x-amz-meta-*` does not survive a round trip.

That behaviour is measured separately, against the `ceph/s3-tests` suite Ceph RGW, MinIO and Garage are all validated with, and the results are published in Predastore's [S3 compatibility report](https://github.com/mulgadc/predastore/blob/dev/docs/S3-COMPATIBILITY.md). Read that before porting a client that depends on exact S3 semantics.

### Operations

| Operation | Status |
|---|---|
| `AbortMultipartUpload` | ✅ Implemented |
| `CompleteMultipartUpload` | ✅ Implemented |
| `CopyObject` | ✅ Implemented |
| `CreateBucket` | ✅ Implemented |
| `CreateMultipartUpload` | ✅ Implemented |
| `CreateSession` | ⛔ Not applicable [3](#notes) |
| `DeleteBucket` | ✅ Implemented |
| `DeleteBucketAnalyticsConfiguration` | ⛔ Not applicable [2](#notes) |
| `DeleteBucketIntelligentTieringConfiguration` | ⛔ Not applicable [8](#notes) |
| `DeleteBucketInventoryConfiguration` | ⛔ Not applicable [4](#notes) |
| `DeleteBucketMetricsConfiguration` | ⛔ Not applicable [2](#notes) |
| `DeleteBucketReplication` | ⛔ Not applicable [6](#notes) |
| `DeleteObject` | ✅ Implemented |
| `DeleteObjects` | ✅ Implemented |
| `GetBucketAccelerateConfiguration` | ⛔ Not applicable [1](#notes) |
| `GetBucketAnalyticsConfiguration` | ⛔ Not applicable [2](#notes) |
| `GetBucketIntelligentTieringConfiguration` | ⛔ Not applicable [8](#notes) |
| `GetBucketInventoryConfiguration` | ⛔ Not applicable [4](#notes) |
| `GetBucketMetricsConfiguration` | ⛔ Not applicable [2](#notes) |
| `GetBucketReplication` | ⛔ Not applicable [6](#notes) |
| `GetBucketRequestPayment` | ⛔ Not applicable [7](#notes) |
| `GetObject` | ✅ Implemented |
| `HeadBucket` | ✅ Implemented |
| `HeadObject` | ✅ Implemented |
| `ListBucketAnalyticsConfigurations` | ⛔ Not applicable [2](#notes) |
| `ListBucketIntelligentTieringConfigurations` | ⛔ Not applicable [8](#notes) |
| `ListBucketInventoryConfigurations` | ⛔ Not applicable [4](#notes) |
| `ListBucketMetricsConfigurations` | ⛔ Not applicable [2](#notes) |
| `ListBuckets` | ✅ Implemented |
| `ListMultipartUploads` | ✅ Implemented |
| `ListObjects` | ✅ Implemented |
| `ListObjectsV2` | ✅ Implemented |
| `ListParts` | ✅ Implemented |
| `PutBucketAccelerateConfiguration` | ⛔ Not applicable [1](#notes) |
| `PutBucketAnalyticsConfiguration` | ⛔ Not applicable [2](#notes) |
| `PutBucketIntelligentTieringConfiguration` | ⛔ Not applicable [8](#notes) |
| `PutBucketInventoryConfiguration` | ⛔ Not applicable [4](#notes) |
| `PutBucketMetricsConfiguration` | ⛔ Not applicable [2](#notes) |
| `PutBucketReplication` | ⛔ Not applicable [6](#notes) |
| `PutBucketRequestPayment` | ⛔ Not applicable [7](#notes) |
| `PutObject` | ✅ Implemented |
| `RestoreObject` | ⛔ Not applicable [8](#notes) |
| `UploadPart` | ✅ Implemented |
| `UploadPartCopy` | ✅ Implemented |
| `WriteGetObjectResponse` | ⛔ Not applicable [5](#notes) |

### Notes

1. Transfer acceleration routes uploads over AWS edge locations, which an on-premise deployment has none of.
2. Storage class analytics and CloudWatch request metrics are AWS-hosted reporting services with no equivalent here.
3. Sessions authenticate S3 Express One Zone directory buckets, a storage class this platform does not offer.
4. Inventory reports are an AWS-hosted scheduled export; a listing here is served live from the metadata store.
5. S3 Object Lambda rewrites a response from a Lambda function, and Lambda is not offered.
6. Replication copies objects to a bucket in another region, and a deployment is a single region with no peer.
7. Requester Pays shifts transfer charges to the caller, which needs AWS billing behind it.
8. Intelligent tiering, archive storage classes and the restores they need are AWS-hosted lifecycle services; objects here are held in one class on local media.
