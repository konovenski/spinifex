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
| `AllocateAddress` | ✅ Implemented |
| `AssociateAddress` | ✅ Implemented |
| `AssociateIamInstanceProfile` | ✅ Implemented |
| `AssociateRouteTable` | ✅ Implemented |
| `AttachInternetGateway` | ✅ Implemented |
| `AttachNetworkInterface` | ✅ Implemented |
| `AttachVolume` | ✅ Implemented |
| `AuthorizeSecurityGroupEgress` | ✅ Implemented |
| `AuthorizeSecurityGroupIngress` | ✅ Implemented |
| `CancelCapacityReservation` | ✅ Implemented |
| `CancelSpotInstanceRequests` | ✅ Implemented |
| `CopyImage` | ✅ Implemented |
| `CopySnapshot` | ✅ Implemented |
| `CreateCapacityReservation` | ✅ Implemented |
| `CreateEgressOnlyInternetGateway` | ✅ Implemented |
| `CreateImage` | ✅ Implemented |
| `CreateInternetGateway` | ✅ Implemented |
| `CreateKeyPair` | ✅ Implemented |
| `CreateLaunchTemplate` | ✅ Implemented |
| `CreateLaunchTemplateVersion` | ✅ Implemented |
| `CreateNatGateway` | ✅ Implemented |
| `CreateNetworkInterface` | ✅ Implemented |
| `CreatePlacementGroup` | ✅ Implemented |
| `CreateRoute` | ✅ Implemented |
| `CreateRouteTable` | ✅ Implemented |
| `CreateSecurityGroup` | ✅ Implemented |
| `CreateSnapshot` | ✅ Implemented |
| `CreateSubnet` | ✅ Implemented |
| `CreateTags` | ✅ Implemented |
| `CreateVolume` | ✅ Implemented |
| `CreateVpc` | ✅ Implemented |
| `DeleteEgressOnlyInternetGateway` | ✅ Implemented |
| `DeleteInternetGateway` | ✅ Implemented |
| `DeleteKeyPair` | ✅ Implemented |
| `DeleteLaunchTemplate` | ✅ Implemented |
| `DeleteLaunchTemplateVersions` | ✅ Implemented |
| `DeleteNatGateway` | ✅ Implemented |
| `DeleteNetworkInterface` | ✅ Implemented |
| `DeletePlacementGroup` | ✅ Implemented |
| `DeleteRoute` | ✅ Implemented |
| `DeleteRouteTable` | ✅ Implemented |
| `DeleteSecurityGroup` | ✅ Implemented |
| `DeleteSnapshot` | ✅ Implemented |
| `DeleteSubnet` | ✅ Implemented |
| `DeleteTags` | ✅ Implemented |
| `DeleteVolume` | ✅ Implemented |
| `DeleteVpc` | ✅ Implemented |
| `DeregisterImage` | ✅ Implemented |
| `DescribeAccountAttributes` | ✅ Implemented |
| `DescribeAddresses` | ✅ Implemented |
| `DescribeAddressesAttribute` | ✅ Implemented |
| `DescribeAvailabilityZones` | ✅ Implemented |
| `DescribeCapacityReservations` | ✅ Implemented |
| `DescribeEgressOnlyInternetGateways` | ✅ Implemented |
| `DescribeIamInstanceProfileAssociations` | ✅ Implemented |
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
| `DescribeNatGateways` | ✅ Implemented |
| `DescribeNetworkInterfaces` | ✅ Implemented |
| `DescribePlacementGroups` | ✅ Implemented |
| `DescribeRegions` | ✅ Implemented |
| `DescribeRouteTables` | ✅ Implemented |
| `DescribeSecurityGroupRules` | ✅ Implemented |
| `DescribeSecurityGroups` | ✅ Implemented |
| `DescribeSnapshots` | ✅ Implemented |
| `DescribeSpotInstanceRequests` | ✅ Implemented |
| `DescribeSpotPriceHistory` | ⛔ Not applicable [1](#notes) |
| `DescribeSubnets` | ✅ Implemented |
| `DescribeTags` | ✅ Implemented |
| `DescribeVolumeStatus` | ✅ Implemented |
| `DescribeVolumes` | ✅ Implemented |
| `DescribeVolumesModifications` | ✅ Implemented |
| `DescribeVpcAttribute` | ✅ Implemented |
| `DescribeVpcs` | ✅ Implemented |
| `DetachInternetGateway` | ✅ Implemented |
| `DetachNetworkInterface` | ✅ Implemented |
| `DetachVolume` | ✅ Implemented |
| `DisableEbsEncryptionByDefault` | ✅ Implemented |
| `DisableSerialConsoleAccess` | ✅ Implemented |
| `DisassociateAddress` | ✅ Implemented |
| `DisassociateIamInstanceProfile` | ✅ Implemented |
| `DisassociateRouteTable` | ✅ Implemented |
| `EnableEbsEncryptionByDefault` | ✅ Implemented |
| `EnableSerialConsoleAccess` | ✅ Implemented |
| `GetConsoleOutput` | ✅ Implemented |
| `GetEbsEncryptionByDefault` | ✅ Implemented |
| `GetPasswordData` | ✅ Implemented |
| `GetSecurityGroupsForVpc` | ✅ Implemented |
| `GetSerialConsoleAccessStatus` | ✅ Implemented |
| `ImportKeyPair` | ✅ Implemented |
| `ModifyImageAttribute` | ✅ Implemented |
| `ModifyInstanceAttribute` | ✅ Implemented |
| `ModifyInstanceMetadataOptions` | ✅ Implemented |
| `ModifyLaunchTemplate` | ✅ Implemented |
| `ModifyNetworkInterfaceAttribute` | ✅ Implemented |
| `ModifySubnetAttribute` | ✅ Implemented |
| `ModifyVolume` | ✅ Implemented |
| `ModifyVpcAttribute` | ✅ Implemented |
| `MonitorInstances` | ✅ Implemented |
| `RebootInstances` | ✅ Implemented |
| `RegisterImage` | ✅ Implemented |
| `ReleaseAddress` | ✅ Implemented |
| `ReplaceIamInstanceProfileAssociation` | ✅ Implemented |
| `ReplaceRoute` | ✅ Implemented |
| `ReplaceRouteTableAssociation` | ✅ Implemented |
| `RequestSpotInstances` | ✅ Implemented |
| `ResetImageAttribute` | ✅ Implemented |
| `RevokeSecurityGroupEgress` | ✅ Implemented |
| `RevokeSecurityGroupIngress` | ✅ Implemented |
| `RunInstances` | ✅ Implemented |
| `StartInstances` | ✅ Implemented |
| `StopInstances` | ✅ Implemented |
| `TerminateInstances` | ✅ Implemented |
| `UnmonitorInstances` | ✅ Implemented |
| `UpdateSecurityGroupRuleDescriptionsEgress` | ✅ Implemented |
| `UpdateSecurityGroupRuleDescriptionsIngress` | ✅ Implemented |

### Notes

1. On owned hardware there is no spot-to-on-demand price differential, so any figure would be invented.
