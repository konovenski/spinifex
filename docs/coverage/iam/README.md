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

| Operation |
|---|
| `AddRoleToInstanceProfile` |
| `AddUserToGroup` |
| `AttachGroupPolicy` |
| `AttachRolePolicy` |
| `AttachUserPolicy` |
| `CreateAccessKey` |
| `CreateGroup` |
| `CreateInstanceProfile` |
| `CreateOpenIDConnectProvider` |
| `CreatePolicy` |
| `CreateRole` |
| `CreateUser` |
| `DeleteAccessKey` |
| `DeleteGroup` |
| `DeleteGroupPolicy` |
| `DeleteInstanceProfile` |
| `DeleteOpenIDConnectProvider` |
| `DeletePolicy` |
| `DeleteRole` |
| `DeleteRolePolicy` |
| `DeleteUser` |
| `DeleteUserPolicy` |
| `DetachGroupPolicy` |
| `DetachRolePolicy` |
| `DetachUserPolicy` |
| `GetAccountSummary` |
| `GetGroup` |
| `GetGroupPolicy` |
| `GetInstanceProfile` |
| `GetOpenIDConnectProvider` |
| `GetPolicy` |
| `GetPolicyVersion` |
| `GetRole` |
| `GetRolePolicy` |
| `GetUser` |
| `GetUserPolicy` |
| `ListAccessKeys` |
| `ListAttachedGroupPolicies` |
| `ListAttachedRolePolicies` |
| `ListAttachedUserPolicies` |
| `ListGroupPolicies` |
| `ListGroups` |
| `ListGroupsForUser` |
| `ListInstanceProfileTags` |
| `ListInstanceProfiles` |
| `ListInstanceProfilesForRole` |
| `ListOpenIDConnectProviderTags` |
| `ListOpenIDConnectProviders` |
| `ListPolicies` |
| `ListPolicyTags` |
| `ListPolicyVersions` |
| `ListRolePolicies` |
| `ListRoleTags` |
| `ListRoles` |
| `ListUserPolicies` |
| `ListUserTags` |
| `ListUsers` |
| `PutGroupPolicy` |
| `PutRolePolicy` |
| `PutUserPolicy` |
| `RemoveRoleFromInstanceProfile` |
| `RemoveUserFromGroup` |
| `TagInstanceProfile` |
| `TagOpenIDConnectProvider` |
| `TagPolicy` |
| `TagRole` |
| `TagUser` |
| `UntagInstanceProfile` |
| `UntagOpenIDConnectProvider` |
| `UntagPolicy` |
| `UntagRole` |
| `UntagUser` |
| `UpdateAccessKey` |
| `UpdateAssumeRolePolicy` |
| `UpdateRole` |

### Not applicable

These operations describe AWS-hosted features this platform does not offer.

| Operation | Reason |
|---|---|
| `CreateServiceLinkedRole` | Service-linked roles exist for AWS service principals, which this platform has none of. |
| `CreateServiceSpecificCredential` | CodeCommit credentials and SSH keys; the service does not exist here. |
| `DeleteSSHPublicKey` | CodeCommit credentials and SSH keys; the service does not exist here. |
| `DeleteServiceLinkedRole` | Service-linked roles exist for AWS service principals, which this platform has none of. |
| `DeleteServiceSpecificCredential` | CodeCommit credentials and SSH keys; the service does not exist here. |
| `DeleteSigningCertificate` | X.509 signing certificates are an EC2-Classic SOAP credential, retired by AWS. |
| `GenerateOrganizationsAccessReport` | AWS Organizations is not offered, so there is no organization to report on. |
| `GetOrganizationsAccessReport` | AWS Organizations is not offered, so there is no organization to report on. |
| `GetSSHPublicKey` | CodeCommit credentials and SSH keys; the service does not exist here. |
| `GetServiceLinkedRoleDeletionStatus` | Service-linked roles exist for AWS service principals, which this platform has none of. |
| `ListSSHPublicKeys` | CodeCommit credentials and SSH keys; the service does not exist here. |
| `ListServiceSpecificCredentials` | CodeCommit credentials and SSH keys; the service does not exist here. |
| `ListSigningCertificates` | X.509 signing certificates are an EC2-Classic SOAP credential, retired by AWS. |
| `ResetServiceSpecificCredential` | CodeCommit credentials and SSH keys; the service does not exist here. |
| `UpdateSSHPublicKey` | CodeCommit credentials and SSH keys; the service does not exist here. |
| `UpdateServiceSpecificCredential` | CodeCommit credentials and SSH keys; the service does not exist here. |
| `UpdateSigningCertificate` | X.509 signing certificates are an EC2-Classic SOAP credential, retired by AWS. |
| `UploadSSHPublicKey` | CodeCommit credentials and SSH keys; the service does not exist here. |
| `UploadSigningCertificate` | X.509 signing certificates are an EC2-Classic SOAP credential, retired by AWS. |
