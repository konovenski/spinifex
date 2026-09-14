### Spot instances are a mock

Spot Instance Requests are a mock over the on-demand `RunInstances` path. A request synchronously launches real VMs on the operator's own compute and is then reported `active` and `fulfilled`. There is no spot market: no bidding, no price rejection, no interruption and no reclamation, and instances are never taken back.

`DescribeSpotPriceHistory` is the one operation on that surface deliberately left unimplemented. On owned hardware there is no spot-to-on-demand price differential, so any figure it returned would be invented, and an invented price is worse than an absent one.
