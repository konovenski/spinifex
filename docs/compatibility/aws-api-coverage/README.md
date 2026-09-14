---
title: "AWS API Coverage"
seoTitle: "Which AWS API Operations Spinifex Implements — Spinifex Docs"
description: "Operation-level coverage of every AWS API Spinifex serves, generated from its gateway dispatch tables and the pinned AWS SDK service models on each build."
category: "Coverage"
sections:
  - overview
tags:
  - aws
  - compatibility
  - coverage
  - api
  - operations
---

# AWS API Coverage

## Overview

Spinifex serves the AWS APIs below. Every page counts the operations in the pinned `aws-sdk-go v1.55.8` `api-2.json` model for its service and reports, operation by operation, whether Spinifex implements it.

| Service | API version | Implemented | Modelled | Coverage |
|---|---|---:|---:|---:|
| [ACM](/docs/acm-api-coverage) | 2015-12-08 | 9 | 15 | 60.0% |
| [EC2](/docs/ec2-api-coverage) | 2016-11-15 | 121 | 625 | 19.4% |
| [ECR](/docs/ecr-api-coverage) | 2015-09-21 | 19 | 47 | 40.4% |
| [ECS](/docs/ecs-api-coverage) | 2014-11-13 | 31 | 56 | 55.4% |
| [EKS](/docs/eks-api-coverage) | 2017-11-01 | 34 | 56 | 60.7% |
| [ELBv2](/docs/elbv2-api-coverage) | 2015-12-01 | 33 | 46 | 71.7% |
| [IAM](/docs/iam-api-coverage) | 2010-05-08 | 75 | 159 | 47.2% |
| [RDS](/docs/rds-api-coverage) | 2014-10-31 | 26 | 162 | 16.0% |
| [S3](/docs/s3-api-coverage) | 2006-03-01 | — | 99 | — |
| [STS](/docs/sts-api-coverage) | 2011-06-15 | 4 | 8 | 50.0% |

| Status | Meaning |
|---|---|
| ✅ Implemented | A modelled operation bound to a real handler. |
| 🟡 Stub | A registered handler that answers with a fixed or empty result. |
| 🚫 Not supported | A registered handler that deliberately refuses, so a client sees "not offered" rather than an unknown action. |
| ❌ Not implemented | Modelled by AWS, not registered by Spinifex. |
| 🔒 Outside the pinned model | Registered by Spinifex but absent from the pinned model — an internal route, not a tenant-callable AWS action. |

### What "implemented" means here

An operation is implemented when the gateway binds it to a real handler. These pages are generated from those dispatch tables on every build, so they cannot go stale against the code. That is a mechanical fact, and it is the only claim they make. It does not say the handler honours every parameter AWS models, and no page reports parameter or field-level detail: a handler receives a fully-populated typed SDK struct, so nothing distinguishes a field it honours from one it ignores. Behavioural conformance is measured separately by the integration conformance suite.
