### Predastore serves this endpoint

S3 is the one surface the AWS gateway does not answer itself. Object storage runs on [Predastore](https://github.com/mulgadc/predastore), which serves the S3 REST API directly over its own TLS endpoint, with Reed-Solomon erasure coding and AES-256-GCM at rest behind it.

S3 selects an operation by method, path, query parameter and header rather than by an action name, so several operations share one route. The table below is generated from the route table Predastore builds its router from: `GET /{bucket}` answers both `ListObjects` and `ListObjectsV2`, and `PUT /{bucket}/{key}` splits four ways on `?partNumber` and `x-amz-copy-source`.

### Routed is not the same as conforming

This page says an operation is routed to a handler. It does not say the handler's behaviour matches S3 in every case, and for several operations it does not: ETag is not the body MD5, `Marker` is ignored on a v1 listing, and `x-amz-meta-*` does not survive a round trip.

That behaviour is measured separately, against the `ceph/s3-tests` suite Ceph RGW, MinIO and Garage are all validated with, and the results are published in Predastore's [S3 compatibility report](https://github.com/mulgadc/predastore/blob/dev/docs/S3-COMPATIBILITY.md). Read that before porting a client that depends on exact S3 semantics.
