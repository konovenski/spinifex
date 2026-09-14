---
title: "S3 API Coverage"
seoTitle: "Amazon S3 API Compatibility on Spinifex — Spinifex Docs"
description: "How Spinifex serves the Amazon S3 API through Predastore, and why its coverage is measured there rather than derived from a gateway dispatch table here."
category: "Coverage"
sections:
  - overview
tags:
  - aws
  - compatibility
  - coverage
  - s3
  - storage
  - predastore
---

# S3 API Coverage

## Overview

The S3 API model pins **99** operations at API version `2006-03-01`, but Spinifex's coverage of them is not mechanically enumerable: Spinifex delegates the S3 REST surface to Predastore, which has no operation-name dispatch table to compare mechanically.

### Where S3 is served

Spinifex does not dispatch S3 operations. The REST surface is delegated whole to **Predastore**, which serves it directly, so there is no operation-name dispatch table in the gateway to compare against the model — the comparison that produces every other page on this site has nothing to read here.

That is why this page reports no percentage. Claiming one from an empty dispatch table would read as 0% implemented, which is the opposite of the truth: S3 is one of the most complete surfaces in the platform.

Predastore publishes its own measured compatibility. Consult it for the operations, headers and behaviours it supports.
