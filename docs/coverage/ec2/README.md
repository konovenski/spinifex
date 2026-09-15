---
title: "EC2 API Coverage"
seoTitle: "Amazon EC2 API Coverage on Spinifex — Spinifex Docs"
description: "The Amazon EC2 API operations Spinifex implements, covering instances, EBS volumes, VPC networking, tags and security groups, with those it does not offer."
category: "Coverage"
sections:
  - overview
tags:
  - aws
  - compatibility
  - coverage
  - ec2
  - compute
  - vpc
---

# EC2 API Coverage

## Overview

Spinifex implements **125 operations** in the EC2 `2016-11-15` API model.

### Spot Instances Are a Mock

Spot Instance Requests are a mock over the on-demand `RunInstances` path. A request synchronously launches real VMs on the operator's own compute and is then reported `active` and `fulfilled`. There is no spot market: no bidding, no price rejection, no interruption and no reclamation, and instances are never taken back.

### Operations

| Operation | Status |
|---|---|
| `AcceptReservedInstancesExchangeQuote` | ⛔ Not applicable [10](#notes) |
| `AdvertiseByoipCidr` | ⛔ Not applicable [5](#notes) |
| `AllocateAddress` | ✅ Implemented |
| `AssociateAddress` | ✅ Implemented |
| `AssociateIamInstanceProfile` | ✅ Implemented |
| `AssociateRouteTable` | ✅ Implemented |
| `AttachClassicLinkVpc` | ⛔ Not applicable [6](#notes) |
| `AttachInternetGateway` | ✅ Implemented |
| `AttachNetworkInterface` | ✅ Implemented |
| `AttachVolume` | ✅ Implemented |
| `AuthorizeSecurityGroupEgress` | ✅ Implemented |
| `AuthorizeSecurityGroupIngress` | ✅ Implemented |
| `CancelCapacityReservation` | ✅ Implemented |
| `CancelReservedInstancesListing` | ⛔ Not applicable [10](#notes) |
| `CancelSpotInstanceRequests` | ✅ Implemented |
| `ConfirmProductInstance` | ⛔ Not applicable [9](#notes) |
| `CopyFpgaImage` | ⛔ Not applicable [7](#notes) |
| `CopyImage` | ✅ Implemented |
| `CopySnapshot` | ✅ Implemented |
| `CreateCapacityReservation` | ✅ Implemented |
| `CreateCarrierGateway` | ⛔ Not applicable [3](#notes) |
| `CreateCoipCidr` | ⛔ Not applicable [3](#notes) |
| `CreateCoipPool` | ⛔ Not applicable [3](#notes) |
| `CreateEgressOnlyInternetGateway` | ✅ Implemented |
| `CreateFpgaImage` | ⛔ Not applicable [7](#notes) |
| `CreateImage` | ✅ Implemented |
| `CreateInternetGateway` | ✅ Implemented |
| `CreateKeyPair` | ✅ Implemented |
| `CreateLaunchTemplate` | ✅ Implemented |
| `CreateLaunchTemplateVersion` | ✅ Implemented |
| `CreateLocalGatewayRoute` | ⛔ Not applicable [3](#notes) |
| `CreateLocalGatewayRouteTable` | ⛔ Not applicable [3](#notes) |
| `CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation` | ⛔ Not applicable [3](#notes) |
| `CreateLocalGatewayRouteTableVpcAssociation` | ⛔ Not applicable [3](#notes) |
| `CreateNatGateway` | ✅ Implemented |
| `CreateNetworkInterface` | ✅ Implemented |
| `CreatePlacementGroup` | ✅ Implemented |
| `CreateReservedInstancesListing` | ⛔ Not applicable [10](#notes) |
| `CreateRoute` | ✅ Implemented |
| `CreateRouteTable` | ✅ Implemented |
| `CreateSecurityGroup` | ✅ Implemented |
| `CreateSnapshot` | ✅ Implemented |
| `CreateSubnet` | ✅ Implemented |
| `CreateTags` | ✅ Implemented |
| `CreateVolume` | ✅ Implemented |
| `CreateVpc` | ✅ Implemented |
| `DeleteCarrierGateway` | ⛔ Not applicable [3](#notes) |
| `DeleteCoipCidr` | ⛔ Not applicable [3](#notes) |
| `DeleteCoipPool` | ⛔ Not applicable [3](#notes) |
| `DeleteEgressOnlyInternetGateway` | ✅ Implemented |
| `DeleteFpgaImage` | ⛔ Not applicable [7](#notes) |
| `DeleteInternetGateway` | ✅ Implemented |
| `DeleteKeyPair` | ✅ Implemented |
| `DeleteLaunchTemplate` | ✅ Implemented |
| `DeleteLaunchTemplateVersions` | ✅ Implemented |
| `DeleteLocalGatewayRoute` | ⛔ Not applicable [3](#notes) |
| `DeleteLocalGatewayRouteTable` | ⛔ Not applicable [3](#notes) |
| `DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation` | ⛔ Not applicable [3](#notes) |
| `DeleteLocalGatewayRouteTableVpcAssociation` | ⛔ Not applicable [3](#notes) |
| `DeleteNatGateway` | ✅ Implemented |
| `DeleteNetworkInterface` | ✅ Implemented |
| `DeletePlacementGroup` | ✅ Implemented |
| `DeleteQueuedReservedInstances` | ⛔ Not applicable [10](#notes) |
| `DeleteRoute` | ✅ Implemented |
| `DeleteRouteTable` | ✅ Implemented |
| `DeleteSecurityGroup` | ✅ Implemented |
| `DeleteSnapshot` | ✅ Implemented |
| `DeleteSubnet` | ✅ Implemented |
| `DeleteTags` | ✅ Implemented |
| `DeleteVolume` | ✅ Implemented |
| `DeleteVpc` | ✅ Implemented |
| `DeprovisionByoipCidr` | ⛔ Not applicable [5](#notes) |
| `DeregisterImage` | ✅ Implemented |
| `DescribeAccountAttributes` | ✅ Implemented |
| `DescribeAddresses` | ✅ Implemented |
| `DescribeAddressesAttribute` | ✅ Implemented |
| `DescribeAggregateIdFormat` | ⛔ Not applicable [8](#notes) |
| `DescribeAvailabilityZones` | ✅ Implemented |
| `DescribeAwsNetworkPerformanceMetricSubscriptions` | ⛔ Not applicable [2](#notes) |
| `DescribeByoipCidrs` | ⛔ Not applicable [5](#notes) |
| `DescribeCapacityBlockOfferings` | ⛔ Not applicable [10](#notes) |
| `DescribeCapacityReservations` | ✅ Implemented |
| `DescribeCarrierGateways` | ⛔ Not applicable [3](#notes) |
| `DescribeClassicLinkInstances` | ⛔ Not applicable [6](#notes) |
| `DescribeCoipPools` | ⛔ Not applicable [3](#notes) |
| `DescribeEgressOnlyInternetGateways` | ✅ Implemented |
| `DescribeElasticGpus` | ⛔ Not applicable [11](#notes) |
| `DescribeFpgaImageAttribute` | ⛔ Not applicable [7](#notes) |
| `DescribeFpgaImages` | ⛔ Not applicable [7](#notes) |
| `DescribeHostReservationOfferings` | ⛔ Not applicable [10](#notes) |
| `DescribeHostReservations` | ⛔ Not applicable [10](#notes) |
| `DescribeIamInstanceProfileAssociations` | ✅ Implemented |
| `DescribeIdFormat` | ⛔ Not applicable [8](#notes) |
| `DescribeIdentityIdFormat` | ⛔ Not applicable [8](#notes) |
| `DescribeImageAttribute` | ✅ Implemented |
| `DescribeImages` | ✅ Implemented |
| `DescribeInstanceAttribute` | ✅ Implemented |
| `DescribeInstanceCreditSpecifications` | ✅ Implemented |
| `DescribeInstanceStatus` | ✅ Implemented |
| `DescribeInstanceTypeOfferings` | ✅ Implemented |
| `DescribeInstanceTypes` | ✅ Implemented |
| `DescribeInstances` | ✅ Implemented |
| `DescribeInternetGateways` | ✅ Implemented |
| `DescribeKeyPairs` | ✅ Implemented |
| `DescribeLaunchTemplateVersions` | ✅ Implemented |
| `DescribeLaunchTemplates` | ✅ Implemented |
| `DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations` | ⛔ Not applicable [3](#notes) |
| `DescribeLocalGatewayRouteTableVpcAssociations` | ⛔ Not applicable [3](#notes) |
| `DescribeLocalGatewayRouteTables` | ⛔ Not applicable [3](#notes) |
| `DescribeLocalGatewayVirtualInterfaceGroups` | ⛔ Not applicable [3](#notes) |
| `DescribeLocalGatewayVirtualInterfaces` | ⛔ Not applicable [3](#notes) |
| `DescribeLocalGateways` | ⛔ Not applicable [3](#notes) |
| `DescribeMovingAddresses` | ⛔ Not applicable [6](#notes) |
| `DescribeNatGateways` | ✅ Implemented |
| `DescribeNetworkInterfaces` | ✅ Implemented |
| `DescribePlacementGroups` | ✅ Implemented |
| `DescribePrincipalIdFormat` | ⛔ Not applicable [8](#notes) |
| `DescribeRegions` | ✅ Implemented |
| `DescribeReservedInstances` | ⛔ Not applicable [10](#notes) |
| `DescribeReservedInstancesListings` | ⛔ Not applicable [10](#notes) |
| `DescribeReservedInstancesModifications` | ⛔ Not applicable [10](#notes) |
| `DescribeReservedInstancesOfferings` | ⛔ Not applicable [10](#notes) |
| `DescribeRouteTables` | ✅ Implemented |
| `DescribeScheduledInstanceAvailability` | ⛔ Not applicable [10](#notes) |
| `DescribeScheduledInstances` | ⛔ Not applicable [10](#notes) |
| `DescribeSecurityGroupRules` | ✅ Implemented |
| `DescribeSecurityGroups` | ✅ Implemented |
| `DescribeSnapshots` | ✅ Implemented |
| `DescribeSpotInstanceRequests` | ✅ Implemented |
| `DescribeSpotPriceHistory` | ⛔ Not applicable [12](#notes) |
| `DescribeSubnets` | ✅ Implemented |
| `DescribeTags` | ✅ Implemented |
| `DescribeVolumeStatus` | ✅ Implemented |
| `DescribeVolumes` | ✅ Implemented |
| `DescribeVolumesModifications` | ✅ Implemented |
| `DescribeVpcAttribute` | ✅ Implemented |
| `DescribeVpcClassicLink` | ⛔ Not applicable [6](#notes) |
| `DescribeVpcClassicLinkDnsSupport` | ⛔ Not applicable [6](#notes) |
| `DescribeVpcs` | ✅ Implemented |
| `DetachClassicLinkVpc` | ⛔ Not applicable [6](#notes) |
| `DetachInternetGateway` | ✅ Implemented |
| `DetachNetworkInterface` | ✅ Implemented |
| `DetachVolume` | ✅ Implemented |
| `DisableAwsNetworkPerformanceMetricSubscription` | ⛔ Not applicable [2](#notes) |
| `DisableEbsEncryptionByDefault` | ✅ Implemented |
| `DisableSerialConsoleAccess` | ✅ Implemented |
| `DisableVpcClassicLink` | ⛔ Not applicable [6](#notes) |
| `DisableVpcClassicLinkDnsSupport` | ⛔ Not applicable [6](#notes) |
| `DisassociateAddress` | ✅ Implemented |
| `DisassociateIamInstanceProfile` | ✅ Implemented |
| `DisassociateRouteTable` | ✅ Implemented |
| `EnableAwsNetworkPerformanceMetricSubscription` | ⛔ Not applicable [2](#notes) |
| `EnableEbsEncryptionByDefault` | ✅ Implemented |
| `EnableSerialConsoleAccess` | ✅ Implemented |
| `EnableVpcClassicLink` | ⛔ Not applicable [6](#notes) |
| `EnableVpcClassicLinkDnsSupport` | ⛔ Not applicable [6](#notes) |
| `GetAwsNetworkPerformanceData` | ⛔ Not applicable [2](#notes) |
| `GetCoipPoolUsage` | ⛔ Not applicable [3](#notes) |
| `GetConsoleOutput` | ✅ Implemented |
| `GetEbsEncryptionByDefault` | ✅ Implemented |
| `GetFlowLogsIntegrationTemplate` | ⛔ Not applicable [1](#notes) |
| `GetHostReservationPurchasePreview` | ⛔ Not applicable [10](#notes) |
| `GetPasswordData` | ✅ Implemented |
| `GetReservedInstancesExchangeQuote` | ⛔ Not applicable [10](#notes) |
| `GetSecurityGroupsForVpc` | ✅ Implemented |
| `GetSerialConsoleAccessStatus` | ✅ Implemented |
| `ImportKeyPair` | ✅ Implemented |
| `ModifyAvailabilityZoneGroup` | ⛔ Not applicable [3](#notes) |
| `ModifyFpgaImageAttribute` | ⛔ Not applicable [7](#notes) |
| `ModifyIdFormat` | ⛔ Not applicable [8](#notes) |
| `ModifyIdentityIdFormat` | ⛔ Not applicable [8](#notes) |
| `ModifyImageAttribute` | ✅ Implemented |
| `ModifyInstanceAttribute` | ✅ Implemented |
| `ModifyInstanceMetadataOptions` | ✅ Implemented |
| `ModifyLaunchTemplate` | ✅ Implemented |
| `ModifyLocalGatewayRoute` | ⛔ Not applicable [3](#notes) |
| `ModifyNetworkInterfaceAttribute` | ✅ Implemented |
| `ModifyReservedInstances` | ⛔ Not applicable [10](#notes) |
| `ModifySubnetAttribute` | ✅ Implemented |
| `ModifyVolume` | ✅ Implemented |
| `ModifyVpcAttribute` | ✅ Implemented |
| `MonitorInstances` | ✅ Implemented |
| `MoveAddressToVpc` | ⛔ Not applicable [6](#notes) |
| `MoveByoipCidrToIpam` | ⛔ Not applicable [5](#notes) |
| `ProvisionByoipCidr` | ⛔ Not applicable [5](#notes) |
| `PurchaseCapacityBlock` | ⛔ Not applicable [10](#notes) |
| `PurchaseHostReservation` | ⛔ Not applicable [10](#notes) |
| `PurchaseReservedInstancesOffering` | ⛔ Not applicable [10](#notes) |
| `PurchaseScheduledInstances` | ⛔ Not applicable [10](#notes) |
| `RebootInstances` | ✅ Implemented |
| `RegisterImage` | ✅ Implemented |
| `ReleaseAddress` | ✅ Implemented |
| `ReplaceIamInstanceProfileAssociation` | ✅ Implemented |
| `ReplaceRoute` | ✅ Implemented |
| `ReplaceRouteTableAssociation` | ✅ Implemented |
| `ReportInstanceStatus` | ⛔ Not applicable [4](#notes) |
| `RequestSpotInstances` | ✅ Implemented |
| `ResetFpgaImageAttribute` | ⛔ Not applicable [7](#notes) |
| `ResetImageAttribute` | ✅ Implemented |
| `RestoreAddressToClassic` | ⛔ Not applicable [6](#notes) |
| `RevokeSecurityGroupEgress` | ✅ Implemented |
| `RevokeSecurityGroupIngress` | ✅ Implemented |
| `RunInstances` | ✅ Implemented |
| `RunScheduledInstances` | ⛔ Not applicable [10](#notes) |
| `SearchLocalGatewayRoutes` | ⛔ Not applicable [3](#notes) |
| `StartInstances` | ✅ Implemented |
| `StopInstances` | ✅ Implemented |
| `TerminateInstances` | ✅ Implemented |
| `UnmonitorInstances` | ✅ Implemented |
| `UpdateSecurityGroupRuleDescriptionsEgress` | ✅ Implemented |
| `UpdateSecurityGroupRuleDescriptionsIngress` | ✅ Implemented |
| `WithdrawByoipCidr` | ⛔ Not applicable [5](#notes) |

### Notes

1. The integration template wires flow logs into Athena and CloudFormation, neither of which this platform offers.
2. These report performance across AWS's own global network between its regions and zones.
3. Outposts, Local Zones and Wavelength extend an AWS region into another site; a deployment here is the site.
4. Reporting an instance as impaired files a report with AWS support about AWS hardware.
5. Bringing your own IP range means advertising it from AWS's network. The operator already owns the addresses here.
6. EC2-Classic is the pre-VPC network AWS retired in 2022. Every instance here is in a VPC.
7. FPGA images are built and ingested through AWS's own AFI pipeline for its F1 instances.
8. These switched resource IDs between the short and long forms during AWS's 2016 migration, which has long since finished.
9. Product codes belong to AWS Marketplace, which has no equivalent here.
10. Reserved Instances, Scheduled Instances, Capacity Blocks and host reservations buy AWS capacity at an AWS price. Hardware here is already bought.
11. Elastic Graphics was retired by AWS, and attached nothing but an AWS-hosted accelerator while it existed.
12. On owned hardware there is no spot-to-on-demand price differential, so any figure would be invented.
