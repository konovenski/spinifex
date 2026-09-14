### Engines

Spinifex offers PostgreSQL and MariaDB. Each DB instance is one dedicated system-owned VM running the engine directly, launched from a platform AMI and hidden from the customer's EC2 API. `Engine` is fixed at create: there is no in-place engine change, no cross-engine snapshot restore and no migration between the two.

`mysql` is not an accepted engine and is not an alias for `mariadb`. MariaDB is offered under its own AWS engine name, exactly as AWS RDS offers it, so a client — including Terraform's `aws_db_instance` — must set `engine = "mariadb"`. Aliasing would report an engine and a version the instance is not running, and the discrepancy would propagate into `DescribeDBInstances`, parameter-group families and snapshot metadata.

Engine versions are pinned per engine. An `EngineVersion` naming anything but the pin is rejected, including a narrower minor version, because the AMI makes no promise about which minor it carries.

### The endpoint is private

The engine is reached over a customer-account ENI injected into a subnet of the DB subnet group, so the endpoint is reachable from inside the VPC only. The DB VM has other NICs that no customer security group governs, and the engine binds none of them — the port is not open there at all, rather than open and gated.

### TLS is required by default

Both engines enforce encrypted connections by default: `rds.force_ssl` on PostgreSQL and `require_secure_transport` on MariaDB. For MariaDB this is a deliberate divergence from AWS, which leaves it off. Both are boolean, modifiable and dynamic, so setting either to `0` in a parameter group restores plaintext without a reboot.

Enforcement requires that the connection is encrypted, not that the client validates the certificate. The engine serves a per-instance certificate signed by the cluster CA carrying both the ENI address and the DNS name, so a client that wants full verification can have it. A deployment holding no cluster CA cannot serve TLS at all, and a parameter group asking for enforcement there is refused rather than quietly ignored.

### Rejected parameters

A parameter whose omission would create a false safety, security or availability guarantee is rejected rather than silently dropped — `MultiAZ=true` on a single-AZ platform, `PubliclyAccessible=true` against a private endpoint, `StorageEncrypted=false` where unencrypted storage is not offered. Parameters that are merely inert are accepted as no-ops.
