---
title: "IAM API Coverage"
seoTitle: "AWS IAM API Operation Coverage on Spinifex — Spinifex Docs"
description: "Every operation in the AWS IAM API model and whether Spinifex implements it, covering users, roles, policies, groups and instance profiles on each build."
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

Spinifex implements **75 of the 159** operations (47.2%) in the IAM `2010-05-08` API model.

### Scope

All IAM operations are account-scoped. The root user of an account bypasses policy evaluation entirely, as it does on AWS.

### Operations

| Operation | Status | Notes |
|---|---|---|
| `AddClientIDToOpenIDConnectProvider` | ❌ Not implemented |  |
| `AddRoleToInstanceProfile` | ✅ Implemented |  |
| `AddUserToGroup` | ✅ Implemented |  |
| `AttachGroupPolicy` | ✅ Implemented |  |
| `AttachRolePolicy` | ✅ Implemented |  |
| `AttachUserPolicy` | ✅ Implemented |  |
| `ChangePassword` | ❌ Not implemented |  |
| `CreateAccessKey` | ✅ Implemented |  |
| `CreateAccountAlias` | ❌ Not implemented |  |
| `CreateGroup` | ✅ Implemented |  |
| `CreateInstanceProfile` | ✅ Implemented |  |
| `CreateLoginProfile` | ❌ Not implemented |  |
| `CreateOpenIDConnectProvider` | ✅ Implemented |  |
| `CreatePolicy` | ✅ Implemented |  |
| `CreatePolicyVersion` | ❌ Not implemented |  |
| `CreateRole` | ✅ Implemented |  |
| `CreateSAMLProvider` | ❌ Not implemented |  |
| `CreateServiceLinkedRole` | ⛔ Not applicable | Service-linked roles exist for AWS service principals, which this platform has none of. |
| `CreateServiceSpecificCredential` | ⛔ Not applicable | CodeCommit credentials; the service does not exist here. |
| `CreateUser` | ✅ Implemented |  |
| `CreateVirtualMFADevice` | ❌ Not implemented |  |
| `DeactivateMFADevice` | ❌ Not implemented |  |
| `DeleteAccessKey` | ✅ Implemented |  |
| `DeleteAccountAlias` | ❌ Not implemented |  |
| `DeleteAccountPasswordPolicy` | ❌ Not implemented |  |
| `DeleteGroup` | ✅ Implemented |  |
| `DeleteGroupPolicy` | ✅ Implemented |  |
| `DeleteInstanceProfile` | ✅ Implemented |  |
| `DeleteLoginProfile` | ❌ Not implemented |  |
| `DeleteOpenIDConnectProvider` | ✅ Implemented |  |
| `DeletePolicy` | ✅ Implemented |  |
| `DeletePolicyVersion` | ❌ Not implemented |  |
| `DeleteRole` | ✅ Implemented |  |
| `DeleteRolePermissionsBoundary` | ❌ Not implemented |  |
| `DeleteRolePolicy` | ✅ Implemented |  |
| `DeleteSAMLProvider` | ❌ Not implemented |  |
| `DeleteSSHPublicKey` | ⛔ Not applicable | CodeCommit SSH keys; the service does not exist here. |
| `DeleteServerCertificate` | ❌ Not implemented |  |
| `DeleteServiceLinkedRole` | ⛔ Not applicable | Service-linked roles exist for AWS service principals, which this platform has none of. |
| `DeleteServiceSpecificCredential` | ⛔ Not applicable | CodeCommit credentials; the service does not exist here. |
| `DeleteSigningCertificate` | ⛔ Not applicable | X.509 signing certificates are an EC2-Classic SOAP credential, retired by AWS. |
| `DeleteUser` | ✅ Implemented |  |
| `DeleteUserPermissionsBoundary` | ❌ Not implemented |  |
| `DeleteUserPolicy` | ✅ Implemented |  |
| `DeleteVirtualMFADevice` | ❌ Not implemented |  |
| `DetachGroupPolicy` | ✅ Implemented |  |
| `DetachRolePolicy` | ✅ Implemented |  |
| `DetachUserPolicy` | ✅ Implemented |  |
| `EnableMFADevice` | ❌ Not implemented |  |
| `GenerateCredentialReport` | ❌ Not implemented |  |
| `GenerateOrganizationsAccessReport` | ⛔ Not applicable | AWS Organizations is not offered, so there is no organization to report on. |
| `GenerateServiceLastAccessedDetails` | ❌ Not implemented |  |
| `GetAccessKeyLastUsed` | ❌ Not implemented |  |
| `GetAccountAuthorizationDetails` | ❌ Not implemented |  |
| `GetAccountPasswordPolicy` | ❌ Not implemented |  |
| `GetAccountSummary` | ✅ Implemented |  |
| `GetContextKeysForCustomPolicy` | ❌ Not implemented |  |
| `GetContextKeysForPrincipalPolicy` | ❌ Not implemented |  |
| `GetCredentialReport` | ❌ Not implemented |  |
| `GetGroup` | ✅ Implemented |  |
| `GetGroupPolicy` | ✅ Implemented |  |
| `GetInstanceProfile` | ✅ Implemented |  |
| `GetLoginProfile` | ❌ Not implemented |  |
| `GetMFADevice` | ❌ Not implemented |  |
| `GetOpenIDConnectProvider` | ✅ Implemented |  |
| `GetOrganizationsAccessReport` | ⛔ Not applicable | AWS Organizations is not offered, so there is no organization to report on. |
| `GetPolicy` | ✅ Implemented |  |
| `GetPolicyVersion` | ✅ Implemented |  |
| `GetRole` | ✅ Implemented |  |
| `GetRolePolicy` | ✅ Implemented |  |
| `GetSAMLProvider` | ❌ Not implemented |  |
| `GetSSHPublicKey` | ⛔ Not applicable | CodeCommit SSH keys; the service does not exist here. |
| `GetServerCertificate` | ❌ Not implemented |  |
| `GetServiceLastAccessedDetails` | ❌ Not implemented |  |
| `GetServiceLastAccessedDetailsWithEntities` | ❌ Not implemented |  |
| `GetServiceLinkedRoleDeletionStatus` | ⛔ Not applicable | Service-linked roles exist for AWS service principals, which this platform has none of. |
| `GetUser` | ✅ Implemented |  |
| `GetUserPolicy` | ✅ Implemented |  |
| `ListAccessKeys` | ✅ Implemented |  |
| `ListAccountAliases` | ❌ Not implemented |  |
| `ListAttachedGroupPolicies` | ✅ Implemented |  |
| `ListAttachedRolePolicies` | ✅ Implemented |  |
| `ListAttachedUserPolicies` | ✅ Implemented |  |
| `ListEntitiesForPolicy` | ❌ Not implemented |  |
| `ListGroupPolicies` | ✅ Implemented |  |
| `ListGroups` | ✅ Implemented |  |
| `ListGroupsForUser` | ✅ Implemented |  |
| `ListInstanceProfileTags` | ✅ Implemented |  |
| `ListInstanceProfiles` | ✅ Implemented |  |
| `ListInstanceProfilesForRole` | ✅ Implemented |  |
| `ListMFADeviceTags` | ❌ Not implemented |  |
| `ListMFADevices` | ❌ Not implemented |  |
| `ListOpenIDConnectProviderTags` | ✅ Implemented |  |
| `ListOpenIDConnectProviders` | ✅ Implemented |  |
| `ListPolicies` | ✅ Implemented |  |
| `ListPoliciesGrantingServiceAccess` | ❌ Not implemented |  |
| `ListPolicyTags` | ✅ Implemented |  |
| `ListPolicyVersions` | ✅ Implemented |  |
| `ListRolePolicies` | ✅ Implemented |  |
| `ListRoleTags` | ✅ Implemented |  |
| `ListRoles` | ✅ Implemented |  |
| `ListSAMLProviderTags` | ❌ Not implemented |  |
| `ListSAMLProviders` | ❌ Not implemented |  |
| `ListSSHPublicKeys` | ⛔ Not applicable | CodeCommit SSH keys; the service does not exist here. |
| `ListServerCertificateTags` | ❌ Not implemented |  |
| `ListServerCertificates` | ❌ Not implemented |  |
| `ListServiceSpecificCredentials` | ⛔ Not applicable | CodeCommit credentials; the service does not exist here. |
| `ListSigningCertificates` | ⛔ Not applicable | X.509 signing certificates are an EC2-Classic SOAP credential, retired by AWS. |
| `ListUserPolicies` | ✅ Implemented |  |
| `ListUserTags` | ✅ Implemented |  |
| `ListUsers` | ✅ Implemented |  |
| `ListVirtualMFADevices` | ❌ Not implemented |  |
| `PutGroupPolicy` | ✅ Implemented |  |
| `PutRolePermissionsBoundary` | ❌ Not implemented |  |
| `PutRolePolicy` | ✅ Implemented |  |
| `PutUserPermissionsBoundary` | ❌ Not implemented |  |
| `PutUserPolicy` | ✅ Implemented |  |
| `RemoveClientIDFromOpenIDConnectProvider` | ❌ Not implemented |  |
| `RemoveRoleFromInstanceProfile` | ✅ Implemented |  |
| `RemoveUserFromGroup` | ✅ Implemented |  |
| `ResetServiceSpecificCredential` | ⛔ Not applicable | CodeCommit credentials; the service does not exist here. |
| `ResyncMFADevice` | ❌ Not implemented |  |
| `SetDefaultPolicyVersion` | ❌ Not implemented |  |
| `SetSecurityTokenServicePreferences` | ❌ Not implemented |  |
| `SimulateCustomPolicy` | ❌ Not implemented |  |
| `SimulatePrincipalPolicy` | ❌ Not implemented |  |
| `TagInstanceProfile` | ✅ Implemented |  |
| `TagMFADevice` | ❌ Not implemented |  |
| `TagOpenIDConnectProvider` | ✅ Implemented |  |
| `TagPolicy` | ✅ Implemented |  |
| `TagRole` | ✅ Implemented |  |
| `TagSAMLProvider` | ❌ Not implemented |  |
| `TagServerCertificate` | ❌ Not implemented |  |
| `TagUser` | ✅ Implemented |  |
| `UntagInstanceProfile` | ✅ Implemented |  |
| `UntagMFADevice` | ❌ Not implemented |  |
| `UntagOpenIDConnectProvider` | ✅ Implemented |  |
| `UntagPolicy` | ✅ Implemented |  |
| `UntagRole` | ✅ Implemented |  |
| `UntagSAMLProvider` | ❌ Not implemented |  |
| `UntagServerCertificate` | ❌ Not implemented |  |
| `UntagUser` | ✅ Implemented |  |
| `UpdateAccessKey` | ✅ Implemented |  |
| `UpdateAccountPasswordPolicy` | ❌ Not implemented |  |
| `UpdateAssumeRolePolicy` | ✅ Implemented |  |
| `UpdateGroup` | ❌ Not implemented |  |
| `UpdateLoginProfile` | ❌ Not implemented |  |
| `UpdateOpenIDConnectProviderThumbprint` | ❌ Not implemented |  |
| `UpdateRole` | ✅ Implemented |  |
| `UpdateRoleDescription` | ❌ Not implemented |  |
| `UpdateSAMLProvider` | ❌ Not implemented |  |
| `UpdateSSHPublicKey` | ⛔ Not applicable | CodeCommit SSH keys; the service does not exist here. |
| `UpdateServerCertificate` | ❌ Not implemented |  |
| `UpdateServiceSpecificCredential` | ⛔ Not applicable | CodeCommit credentials; the service does not exist here. |
| `UpdateSigningCertificate` | ⛔ Not applicable | X.509 signing certificates are an EC2-Classic SOAP credential, retired by AWS. |
| `UpdateUser` | ❌ Not implemented |  |
| `UploadSSHPublicKey` | ⛔ Not applicable | CodeCommit SSH keys; the service does not exist here. |
| `UploadServerCertificate` | ❌ Not implemented |  |
| `UploadSigningCertificate` | ⛔ Not applicable | X.509 signing certificates are an EC2-Classic SOAP credential, retired by AWS. |
