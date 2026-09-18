---
title: "Predastore Distributed Storage"
seoTitle: "How Predastore Distributed Storage Works — Spinifex Docs"
description: "How Spinifex stores objects across a cluster: gate, blob and meta nodes, Reed-Solomon erasure coding, placement, degraded writes, repair and encryption at rest."
category: "Admin"
tags:
  - storage
  - s3
  - predastore
  - erasure coding
  - durability
resources:
  - title: "Predastore Repository"
    url: "https://github.com/mulgadc/predastore"
  - title: "Viperblock (EBS)"
    url: "https://github.com/mulgadc/viperblock"
  - title: "Multi-Node Install"
    url: "/docs/install-multi-node"
  - title: "Network Connections"
    url: "/docs/network-connections"
---

# Predastore Distributed Storage

> How Spinifex stores objects across a cluster, and what a server failure costs.

## Table of Contents

- [Overview](#overview)
- [Hosts and nodes](#hosts-and-nodes)
- [The configuration file](#the-configuration-file)
- [Erasure coding and durability](#erasure-coding-and-durability)
- [Placement](#placement)
- [The write path](#the-write-path)
- [The read path](#the-read-path)
- [Repair](#repair)
- [Encryption at rest](#encryption-at-rest)
- [What a server failure costs](#what-a-server-failure-costs)
- [Checking it](#checking-it)
- [Troubleshooting](#troubleshooting)

---

## Overview

Predastore is the S3-compatible object store underneath a Spinifex cluster, and almost everything durable ends up in it:

- **Machine images** — the AMIs instances launch from
- **Block storage** — EBS volume data and snapshots, written through [Viperblock](https://github.com/mulgadc/viperblock), which uses Predastore as its S3 backend
- **User data** — cloud-init configuration and system artifacts
- **DNS zone files** — the zones Northstar serves
- **Anything your own workloads put in a bucket** through the S3 API

This is background reading. You do not configure any of it by hand: `spx admin init` and `spx admin join` build the topology from the servers that actually form the cluster, and every machine gets the same file with its own host ID recorded in `spinifex.toml`.

## Instructions

## Hosts and Nodes

A cluster is two levels.

A **host** is one machine: one Predastore process owning that machine's data directory, TLS identity and encryption key. A **node** is a role pinned to a host, running inside that process on its own port. Every server in a Spinifex cluster is one host carrying three nodes:

| Role | Port | Purpose |
|---|---|---|
| `gate` | TCP 8443 | Serves the S3 API — SigV4, IAM, erasure coding, placement. Every server runs one, so any of them answers an S3 request. |
| `blob` | UDP 6660 | Holds erasure-coded object shards. One per machine. |
| `meta` | UDP 7660 | Member of the Raft quorum over global state — buckets and the object index. |

Ports have to be unique within a host but not across the cluster, so every machine uses the same three. That is why the port surface does not widen as the cluster grows — see [Network requirements](/docs/install-multi-node#network-requirements).

Which transport a pair of nodes uses follows from where they sit, not from a mode flag. Nodes on the same machine talk over an in-process pipe and open no socket at all; nodes on different machines talk over QUIC, TLS 1.3, authenticated by the cluster CA. A single-server install is therefore one process listening on 8443 and nothing else.

Nothing dials a gate, so a gate binds no cluster-plane listener. Its port is the S3 port and nothing more.

## The Configuration File

The whole cluster's storage topology lives in `/etc/spinifex/predastore/predastore.toml`, identical on every server apart from which host the local process is told to be.

Each server is one `[[host]]` block carrying three `[[host.node]]` blocks:

```toml
[[host]]
id = 1
bind_addr = "10.2.0.2"                                   # cluster plane: raft and blob
addr = "10.2.0.2"                                        # what the other hosts dial
data_dir = "/var/lib/spinifex/predastore"
tls_cert = "/etc/spinifex/server.pem"
tls_key = "/etc/spinifex/server.key"
encryption_key = "/etc/spinifex/predastore/encryption.key"
admin_port = 8660                                        # /healthz, /readyz, /repair

[[host.node]]
id = 1
role = "gate"
port = 8443
bind_addr = "0.0.0.0"                                    # public plane; only a gate may set this
```

The two bind addresses are the point. `bind_addr` on the host is the cluster plane — Raft consensus and shard traffic listen there and nowhere else, at the address peers dial rather than a wildcard, so on a multi-homed machine replication never reaches the public interface. The gate binds the public plane separately, because S3 is meant to be reachable.

## Erasure Coding and Durability

The erasure code is cluster configuration, not a per-request choice, and it is chosen at formation from the number of servers — each machine contributes exactly one blob node:

| Servers | Code | What it survives |
|---|---|---|
| 1 | `RS(1,0)` | Nothing at this layer. One shard on one node; redundancy is the pool's job — ZFS raidz or a mirror under the data directory. |
| 2 | `RS(1,1)` | One server's shards, but the metadata quorum has no majority to lose. |
| 3 or more | `RS(2,1)` | Any one server's shards, with the metadata quorum intact. |

An object is encoded end to end into `K` data and `M` parity shards with no intermediate chunking. Each shard is roughly `⌈size / K⌉` bytes, and any `K` of the `K+M` shards rebuild the object.

`data + parity` must not exceed the number of blob nodes, because a stripe is spread over distinct nodes and a narrower cluster would fail every write. `RS(2,1)` surviving the loss of any one server is the reason [three servers is the recommended minimum](/docs/install-multi-node).

## Placement

Which nodes hold an object's shards comes from a consistent hash ring over the blob nodes, keyed by the object hash. `Nodes(hash, K+M)` returns the placement in shard order — shard `i` goes to the node at position `i` — so every gate in the cluster derives the same placement from the same object without consulting anything.

The gate then **records** that placement in the metadata plane rather than trusting the ring at read time. A ring that has since changed shape does not lose objects placed under the old one.

## The Write Path

The gate authenticates the request, resolves the bucket, hashes the object name, places it on the ring, encodes the body into data and parity shards, streams each shard to its blob node in parallel, and records the placement.

Three properties matter operationally:

- **The write is two-phase.** Each shard put makes the bytes durable and invisible. The placement record landing in Raft is the commit point, and the shards are published only afterwards — so an interrupted writer costs a round trip rather than the object. Shards are always durable before metadata: a crash between the two leaves shards nothing points at, rather than metadata pointing at data that was never written.
- **Every write mints an epoch**, stamped on every shard of the stripe and into the record. The record names one generation, and a shard either belongs to it or does not. That is what makes a half-completed overwrite a reconstruction rather than an object spliced from two generations.
- **One server down does not refuse writes.** The write is acknowledged once `K` shards are durable, and the response carries `X-Spx-Degraded-Write: <n>`. `K` is the floor rather than `K+1` because any `K` of the `K+M` reconstruct — at `RS(2,1)` one more buys no redundancy and would refuse three quarters of writes with a single server down.

A shard its owner will not take is written to the next node along the ring instead of being given up on — **hinted handoff** — and the response carries `X-Spx-Handoff: <n>`. The stripe is then complete and the object is as redundant as a full-width write; what is outstanding is only that the shards are not all where the record says. The holder is derived rather than recorded, so there is no hint to store, replicate or lose. Handoff is skipped when the cluster has no node to spare, since a second shard of one object on a node that already holds one is not redundancy.

If any node reported storage pressure the response also carries `X-Predastore-Pool-Pressure: nearfull`; a node that refused outright surfaces as `507 InsufficientStorage`.

## The Read Path

The gate reads the recorded placement, fetches the data shards, and joins them. Parity is never touched on the happy path. Only if the join fails does it refetch including parity and reconstruct, and a read that had to reconstruct says so in `X-Spx-Degraded: <n>` — a cost report, not an error.

Every shard read names the record's epoch, and a node holding another generation answers `epoch-mismatch`. **The gate counts that exactly as a missing shard**, so the read either rebuilds the object the record names or fails. It never returns plausible nonsense.

A shard its owner cannot serve is looked for on the handoff node before it is counted as failed, since a write that could not reach the owner will have put it there.

A `Range` request that lands entirely inside one data shard takes a fast path — a single ranged shard read — because Reed-Solomon splits data sequentially. Anything wider reconstructs the object and slices it.

## Repair

Degraded writes trade redundancy for availability, and repair is what pays that back. It restores shards a blob node owns but does not hold at the generation its object's record names, and it runs on every server with a local blob node.

Two things are worth knowing about it:

- **It carries no correctness weight.** A read compares each shard's epoch against the record it has already loaded and discards a stale one whether or not repair has ever run. What repair restores is redundancy — the number of further losses an object survives — and never the object.
- **It rebuilds onto the owner**, so it closes the gap only once the missing server returns. At `RS(2,1)` an object written while a server is down holds no parity for the duration of the outage. That is the accepted position: refusing writes meanwhile is the worse failure.

Every pass scans the placement records from the beginning, asks each holder which generation it has, and rebuilds the ones that disagree. A shard that was handed off is returned rather than recomputed — cheaper by a decode — and the holder's copy is released only after the owner's is published. If fewer than `K` peers hold the record's generation the rebuild fails and says so, rather than inventing a shard from a mixture of generations.

## Encryption at Rest

Every fragment is sealed independently under AES-256-GCM, with a deterministic nonce and additional data reconstructed at read time from the on-disk header plus per-data-directory state. There is no unencrypted code path: the process refuses to start without an encryption key.

| Property | What it means |
|---|---|
| **Confidentiality** | Disk-level access yields ciphertext and tags. Plaintext is never written. |
| **Authenticated integrity** | GCM is the sole integrity authority; there is no separate checksum. Corruption, tampering and a wrong key all surface as one integrity error. |
| **Position binding** | The additional data binds each fragment to its slot, so swapping fragments between objects fails the tag. |
| **Cross-directory defence** | A per-data-directory store ID enters the nonce, so splicing a fragment from one data directory into another fails the tag. |

The key is one 32-byte AES-256 key per cluster, and **every server must be given the same file** — fragments sealed under one key cannot be opened under another. `spx admin init` generates it and `spx admin join` distributes it, which is also why joining a server that has already been in service discards its own key and orphans everything sealed under it. The loader is fail-closed on permissions: a group- or other-readable key file is rejected outright, with no override.

## What a Server Failure Costs

| What fails | Reads | Writes |
|---|---|---|
| One gate (of three or more) | Unaffected — any gate answers for the cluster | Unaffected |
| One blob node, `RS(2,1)` | Served by reconstruction, reported in `X-Spx-Degraded` | Accepted at the `K` floor, reported in `X-Spx-Degraded-Write`, or placed complete via handoff |
| More than `M` blob nodes | Objects placed on them are lost | Refused |
| Metadata quorum | Serve whatever the surviving replicas last applied | Stopped entirely |

Reads try the cached Raft leader first, then every replica, and report not-found only once every replica has answered not-found — a replica that has not applied a key yet must not be mistaken for absence. Writes go through the leader, and a replica that cannot commit redirects the client to the one that can.

Replacing a server restores capacity, not the shards that were on the old one: repair rebuilds onto the owner it expects, so plan a replacement as a return of that node rather than as re-replication onto a new one.

## Checking It

Each host runs an unauthenticated admin listener on port 8660, bound to the cluster plane rather than the S3 address. It serves counters and readiness, never addresses or keys:

```bash
curl -s http://$(hostname -I | awk '{print $1}'):8660/readyz
curl -s http://$(hostname -I | awk '{print $1}'):8660/repair
```

`/readyz` reports the most recent probe cycle for the nodes in that process, and answers `503` until one has completed — reporting ready on no information would send traffic to a process that has never answered the question. `/healthz` answers for the process itself: reaching the handler is the whole check.

That the gate is serving is an S3 request like any other:

```bash
export AWS_PROFILE=spinifex
aws s3 ls
```

To see whether a write or read was degraded, ask for the response headers:

```bash
aws s3api put-object --bucket predastore --key probe --body /etc/hostname --debug 2>&1 | grep -i 'x-spx-\|x-predastore-'
```

And `spx get nodes` should list `predastore` in `SERVICES` on every node — a short service list on one server means something failed to start there.

## Troubleshooting

### S3 Answers but Every Write Fails

The gate is up and the metadata quorum is not. Check that a majority of servers are running `spinifex-predastore`:

```bash
spx get nodes
sudo systemctl is-active spinifex-predastore
```

Reads keep serving what the surviving replicas last applied, which is why this presents as a write-only outage.

### 507 InsufficientStorage, or `nearfull` on Every Response

A blob node is out of space, or close enough that compaction has been kicked. Free space in the data directory on each server:

```bash
df -h /var/lib/spinifex/predastore
```

Capacity here is per server, not pooled: a stripe needs its shard on a specific node, so one full server refuses writes for the objects the ring places on it.

### Reads Fail with an Integrity Error

GCM is the only integrity authority, so corruption, tampering and a wrong encryption key are indistinguishable. Confirm the key first, since it is by far the most likely of the three — every server must hold the identical file:

```bash
sudo sha256sum /etc/spinifex/predastore/encryption.key
```

Compare the digest across servers. A mismatch means that server was initialized separately rather than joined, and the fragments it sealed cannot be opened by the cluster.

### Degraded Reads That Never Clear

`X-Spx-Degraded` on reads long after an outage means repair has not restored the shards. Repair rebuilds onto the node that owns the shard, so it cannot finish while that node is absent:

```bash
curl -s http://$(hostname -I | awk '{print $1}'):8660/repair
```

`"enabled": false` means the sweep is off on that host. Otherwise bring the missing server back and let a pass complete — the sweep runs every five minutes by default.
