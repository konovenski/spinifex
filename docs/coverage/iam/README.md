---
title: "IAM API Coverage"
seoTitle: "AWS IAM API Operation Coverage on Spinifex — Spinifex Docs"
description: "The AWS IAM API operations Spinifex implements, covering users, roles, policies, groups and instance profiles, with those the platform does not offer."
category: "Coverage"
sections:
  - overview
tags:
  - aws
  - compatibility
  - coverage
  - iam
  - identity
  - policies
---

# IAM API Coverage

## Overview

Spinifex implements **75 operations** in the IAM `2010-05-08` API model.

### Scope

All IAM operations are account-scoped. The root user of an account bypasses policy evaluation entirely, as it does on AWS.

### Operations

| Operation | Status |
|---|---|
| `AddRoleToInstanceProfile` | ✅ Implemented |
| `AddUserToGroup` | ✅ Implemented |
| `AttachGroupPolicy` | ✅ Implemented |
| `AttachRolePolicy` | ✅ Implemented |
| `AttachUserPolicy` | ✅ Implemented |
| `CreateAccessKey` | ✅ Implemented |
| `CreateGroup` | ✅ Implemented |
| `CreateInstanceProfile` | ✅ Implemented |
| `CreateOpenIDConnectProvider` | ✅ Implemented |
| `CreatePolicy` | ✅ Implemented |
| `CreateRole` | ✅ Implemented |
| `CreateServiceLinkedRole` | ⛔ Not applicable [3](#notes) |
| `CreateServiceSpecificCredential` | ⛔ Not applicable [1](#notes) |
| `CreateUser` | ✅ Implemented |
| `DeleteAccessKey` | ✅ Implemented |
| `DeleteGroup` | ✅ Implemented |
| `DeleteGroupPolicy` | ✅ Implemented |
| `DeleteInstanceProfile` | ✅ Implemented |
| `DeleteOpenIDConnectProvider` | ✅ Implemented |
| `DeletePolicy` | ✅ Implemented |
| `DeleteRole` | ✅ Implemented |
| `DeleteRolePolicy` | ✅ Implemented |
| `DeleteSSHPublicKey` | ⛔ Not applicable [1](#notes) |
| `DeleteServiceLinkedRole` | ⛔ Not applicable [3](#notes) |
| `DeleteServiceSpecificCredential` | ⛔ Not applicable [1](#notes) |
| `DeleteSigningCertificate` | ⛔ Not applicable [4](#notes) |
| `DeleteUser` | ✅ Implemented |
| `DeleteUserPolicy` | ✅ Implemented |
| `DetachGroupPolicy` | ✅ Implemented |
| `DetachRolePolicy` | ✅ Implemented |
| `DetachUserPolicy` | ✅ Implemented |
| `GenerateOrganizationsAccessReport` | ⛔ Not applicable [2](#notes) |
| `GetAccountSummary` | ✅ Implemented |
| `GetGroup` | ✅ Implemented |
| `GetGroupPolicy` | ✅ Implemented |
| `GetInstanceProfile` | ✅ Implemented |
| `GetOpenIDConnectProvider` | ✅ Implemented |
| `GetOrganizationsAccessReport` | ⛔ Not applicable [2](#notes) |
| `GetPolicy` | ✅ Implemented |
| `GetPolicyVersion` | ✅ Implemented |
| `GetRole` | ✅ Implemented |
| `GetRolePolicy` | ✅ Implemented |
| `GetSSHPublicKey` | ⛔ Not applicable [1](#notes) |
| `GetServiceLinkedRoleDeletionStatus` | ⛔ Not applicable [3](#notes) |
| `GetUser` | ✅ Implemented |
| `GetUserPolicy` | ✅ Implemented |
| `ListAccessKeys` | ✅ Implemented |
| `ListAttachedGroupPolicies` | ✅ Implemented |
| `ListAttachedRolePolicies` | ✅ Implemented |
| `ListAttachedUserPolicies` | ✅ Implemented |
| `ListGroupPolicies` | ✅ Implemented |
| `ListGroups` | ✅ Implemented |
| `ListGroupsForUser` | ✅ Implemented |
| `ListInstanceProfileTags` | ✅ Implemented |
| `ListInstanceProfiles` | ✅ Implemented |
| `ListInstanceProfilesForRole` | ✅ Implemented |
| `ListOpenIDConnectProviderTags` | ✅ Implemented |
| `ListOpenIDConnectProviders` | ✅ Implemented |
| `ListPolicies` | ✅ Implemented |
| `ListPolicyTags` | ✅ Implemented |
| `ListPolicyVersions` | ✅ Implemented |
| `ListRolePolicies` | ✅ Implemented |
| `ListRoleTags` | ✅ Implemented |
| `ListRoles` | ✅ Implemented |
| `ListSSHPublicKeys` | ⛔ Not applicable [1](#notes) |
| `ListServiceSpecificCredentials` | ⛔ Not applicable [1](#notes) |
| `ListSigningCertificates` | ⛔ Not applicable [4](#notes) |
| `ListUserPolicies` | ✅ Implemented |
| `ListUserTags` | ✅ Implemented |
| `ListUsers` | ✅ Implemented |
| `PutGroupPolicy` | ✅ Implemented |
| `PutRolePolicy` | ✅ Implemented |
| `PutUserPolicy` | ✅ Implemented |
| `RemoveRoleFromInstanceProfile` | ✅ Implemented |
| `RemoveUserFromGroup` | ✅ Implemented |
| `ResetServiceSpecificCredential` | ⛔ Not applicable [1](#notes) |
| `TagInstanceProfile` | ✅ Implemented |
| `TagOpenIDConnectProvider` | ✅ Implemented |
| `TagPolicy` | ✅ Implemented |
| `TagRole` | ✅ Implemented |
| `TagUser` | ✅ Implemented |
| `UntagInstanceProfile` | ✅ Implemented |
| `UntagOpenIDConnectProvider` | ✅ Implemented |
| `UntagPolicy` | ✅ Implemented |
| `UntagRole` | ✅ Implemented |
| `UntagUser` | ✅ Implemented |
| `UpdateAccessKey` | ✅ Implemented |
| `UpdateAssumeRolePolicy` | ✅ Implemented |
| `UpdateRole` | ✅ Implemented |
| `UpdateSSHPublicKey` | ⛔ Not applicable [1](#notes) |
| `UpdateServiceSpecificCredential` | ⛔ Not applicable [1](#notes) |
| `UpdateSigningCertificate` | ⛔ Not applicable [4](#notes) |
| `UploadSSHPublicKey` | ⛔ Not applicable [1](#notes) |
| `UploadSigningCertificate` | ⛔ Not applicable [4](#notes) |

### Notes

1. CodeCommit credentials and SSH keys; the service does not exist here.
2. AWS Organizations is not offered, so there is no organization to report on.
3. Service-linked roles exist for AWS service principals, which this platform has none of.
4. X.509 signing certificates are an EC2-Classic SOAP credential, retired by AWS.
