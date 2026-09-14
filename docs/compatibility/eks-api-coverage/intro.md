### Routing

EKS is REST-JSON, dispatched by HTTP method and path rather than by an action parameter, so the gateway's routing table is the authoritative list of what it serves. Clusters run K3s on control-plane VMs; nodegroups are EC2 instances joined to them.

### Internal routes

Four routes on this surface are internal, not AWS actions, and appear in the table below as registered outside the pinned model:

- `PublishInternal` — the control-plane VM broker, relaying bootstrap and state posts onto the message bus.
- `WebhookTokenReview` — the token webhook posting bearer tokens for host-side resolution.
- `ListInternalAddons` — the on-VM add-on sync agent fetching its staged manifests.
- `GetRecoveryDirective` — the on-VM recovery agent reading its per-member directive at boot.

None is callable with an ordinary tenant credential. Each names the target account in the path rather than taking it from the credential, so each requires a control-plane instance-role session bound to the named cluster; a broad `eks:*` grant does not reach them.
