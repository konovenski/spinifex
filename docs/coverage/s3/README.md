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

| Operation |
|---|
| `AbortMultipartUpload` |
| `CompleteMultipartUpload` |
| `CopyObject` |
| `CreateBucket` |
| `CreateMultipartUpload` |
| `DeleteBucket` |
| `DeleteObject` |
| `DeleteObjects` |
| `GetObject` |
| `HeadBucket` |
| `HeadObject` |
| `ListBuckets` |
| `ListMultipartUploads` |
| `ListObjects` |
| `ListObjectsV2` |
| `ListParts` |
| `PutObject` |
| `UploadPart` |
| `UploadPartCopy` |

### Not applicable

These operations describe AWS-hosted features this platform does not offer.

| Operation | Reason |
|---|---|
| `CreateSession` | Sessions authenticate S3 Express One Zone directory buckets, a storage class this platform does not offer. |
| `DeleteBucketAnalyticsConfiguration` | Storage class analytics and CloudWatch request metrics are AWS-hosted reporting services with no equivalent here. |
| `DeleteBucketIntelligentTieringConfiguration` | Intelligent tiering, archive storage classes and the restores they need are AWS-hosted lifecycle services; objects here are held in one class on local media. |
| `DeleteBucketInventoryConfiguration` | Inventory reports are an AWS-hosted scheduled export; a listing here is served live from the metadata store. |
| `DeleteBucketMetricsConfiguration` | Storage class analytics and CloudWatch request metrics are AWS-hosted reporting services with no equivalent here. |
| `DeleteBucketReplication` | Replication copies objects to a bucket in another region, and a deployment is a single region with no peer. |
| `GetBucketAccelerateConfiguration` | Transfer acceleration routes uploads over AWS edge locations, which an on-premise deployment has none of. |
| `GetBucketAnalyticsConfiguration` | Storage class analytics and CloudWatch request metrics are AWS-hosted reporting services with no equivalent here. |
| `GetBucketIntelligentTieringConfiguration` | Intelligent tiering, archive storage classes and the restores they need are AWS-hosted lifecycle services; objects here are held in one class on local media. |
| `GetBucketInventoryConfiguration` | Inventory reports are an AWS-hosted scheduled export; a listing here is served live from the metadata store. |
| `GetBucketMetricsConfiguration` | Storage class analytics and CloudWatch request metrics are AWS-hosted reporting services with no equivalent here. |
| `GetBucketReplication` | Replication copies objects to a bucket in another region, and a deployment is a single region with no peer. |
| `GetBucketRequestPayment` | Requester Pays shifts transfer charges to the caller, which needs AWS billing behind it. |
| `ListBucketAnalyticsConfigurations` | Storage class analytics and CloudWatch request metrics are AWS-hosted reporting services with no equivalent here. |
| `ListBucketIntelligentTieringConfigurations` | Intelligent tiering, archive storage classes and the restores they need are AWS-hosted lifecycle services; objects here are held in one class on local media. |
| `ListBucketInventoryConfigurations` | Inventory reports are an AWS-hosted scheduled export; a listing here is served live from the metadata store. |
| `ListBucketMetricsConfigurations` | Storage class analytics and CloudWatch request metrics are AWS-hosted reporting services with no equivalent here. |
| `PutBucketAccelerateConfiguration` | Transfer acceleration routes uploads over AWS edge locations, which an on-premise deployment has none of. |
| `PutBucketAnalyticsConfiguration` | Storage class analytics and CloudWatch request metrics are AWS-hosted reporting services with no equivalent here. |
| `PutBucketIntelligentTieringConfiguration` | Intelligent tiering, archive storage classes and the restores they need are AWS-hosted lifecycle services; objects here are held in one class on local media. |
| `PutBucketInventoryConfiguration` | Inventory reports are an AWS-hosted scheduled export; a listing here is served live from the metadata store. |
| `PutBucketMetricsConfiguration` | Storage class analytics and CloudWatch request metrics are AWS-hosted reporting services with no equivalent here. |
| `PutBucketReplication` | Replication copies objects to a bucket in another region, and a deployment is a single region with no peer. |
| `PutBucketRequestPayment` | Requester Pays shifts transfer charges to the caller, which needs AWS billing behind it. |
| `RestoreObject` | Intelligent tiering, archive storage classes and the restores they need are AWS-hosted lifecycle services; objects here are held in one class on local media. |
| `WriteGetObjectResponse` | S3 Object Lambda rewrites a response from a Lambda function, and Lambda is not offered. |
