### Predastore serves this endpoint

S3 is the one surface the AWS gateway does not answer itself. Object storage runs on [Predastore](https://github.com/mulgadc/predastore), which serves the S3 REST API directly over its own endpoint.

### Routed is not the same as conforming

This page says an operation is routed to a handler. It does not say the handler's behaviour matches S3 in every case.

That behaviour is measured separately, against the `ceph/s3-tests` suite Ceph RGW, MinIO and Garage are all validated with, and the results are published in Predastore's [S3 compatibility report](https://github.com/mulgadc/predastore/blob/dev/docs/S3-COMPATIBILITY.md).
