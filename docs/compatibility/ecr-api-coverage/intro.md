### Two endpoints, one registry

Repository metadata is served over the AWS API on the gateway endpoint. Image data moves over the OCI Distribution `/v2/` endpoint on that same host, authenticated by the bearer token `GetAuthorizationToken` mints for `docker login`.

The split explains the largest group of unimplemented operations below. The layer-transfer operations — `BatchCheckLayerAvailability`, `InitiateLayerUpload`, `UploadLayerPart`, `CompleteLayerUpload` and `GetDownloadUrlForLayer` — are registered stubs because the `/v2/` endpoint carries that traffic instead. A client using `docker` or any OCI-compatible tool never calls them.

Vulnerability scanning is not offered, so the image-scanning operations refuse deliberately rather than pretending to have found nothing. Registry replication is likewise absent: a deployment is a single registry, with no cross-region peer to replicate to.
