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

Spinifex implements **75 of the 159** operations in the IAM `2010-05-08` API model, as pinned in `aws-sdk-go v1.55.8` — **47.2%**.

### Scope

All IAM operations are account-scoped. The root user of an account bypasses policy evaluation entirely, as it does on AWS.

### Operations

| Operation | Status |
|---|---|
| `AddClientIDToOpenIDConnectProvider` | ❌ Not implemented |
| `AddRoleToInstanceProfile` | ✅ Implemented |
| `AddUserToGroup` | ✅ Implemented |
| `AttachGroupPolicy` | ✅ Implemented |
| `AttachRolePolicy` | ✅ Implemented |
| `AttachUserPolicy` | ✅ Implemented |
| `ChangePassword` | ❌ Not implemented |
| `CreateAccessKey` | ✅ Implemented |
| `CreateAccountAlias` | ❌ Not implemented |
| `CreateGroup` | ✅ Implemented |
| `CreateInstanceProfile` | ✅ Implemented |
| `CreateLoginProfile` | ❌ Not implemented |
| `CreateOpenIDConnectProvider` | ✅ Implemented |
| `CreatePolicy` | ✅ Implemented |
| `CreatePolicyVersion` | ❌ Not implemented |
| `CreateRole` | ✅ Implemented |
| `CreateSAMLProvider` | ❌ Not implemented |
| `CreateServiceLinkedRole` | ❌ Not implemented |
| `CreateServiceSpecificCredential` | ❌ Not implemented |
| `CreateUser` | ✅ Implemented |
| `CreateVirtualMFADevice` | ❌ Not implemented |
| `DeactivateMFADevice` | ❌ Not implemented |
| `DeleteAccessKey` | ✅ Implemented |
| `DeleteAccountAlias` | ❌ Not implemented |
| `DeleteAccountPasswordPolicy` | ❌ Not implemented |
| `DeleteGroup` | ✅ Implemented |
| `DeleteGroupPolicy` | ✅ Implemented |
| `DeleteInstanceProfile` | ✅ Implemented |
| `DeleteLoginProfile` | ❌ Not implemented |
| `DeleteOpenIDConnectProvider` | ✅ Implemented |
| `DeletePolicy` | ✅ Implemented |
| `DeletePolicyVersion` | ❌ Not implemented |
| `DeleteRole` | ✅ Implemented |
| `DeleteRolePermissionsBoundary` | ❌ Not implemented |
| `DeleteRolePolicy` | ✅ Implemented |
| `DeleteSAMLProvider` | ❌ Not implemented |
| `DeleteSSHPublicKey` | ❌ Not implemented |
| `DeleteServerCertificate` | ❌ Not implemented |
| `DeleteServiceLinkedRole` | ❌ Not implemented |
| `DeleteServiceSpecificCredential` | ❌ Not implemented |
| `DeleteSigningCertificate` | ❌ Not implemented |
| `DeleteUser` | ✅ Implemented |
| `DeleteUserPermissionsBoundary` | ❌ Not implemented |
| `DeleteUserPolicy` | ✅ Implemented |
| `DeleteVirtualMFADevice` | ❌ Not implemented |
| `DetachGroupPolicy` | ✅ Implemented |
| `DetachRolePolicy` | ✅ Implemented |
| `DetachUserPolicy` | ✅ Implemented |
| `EnableMFADevice` | ❌ Not implemented |
| `GenerateCredentialReport` | ❌ Not implemented |
| `GenerateOrganizationsAccessReport` | ❌ Not implemented |
| `GenerateServiceLastAccessedDetails` | ❌ Not implemented |
| `GetAccessKeyLastUsed` | ❌ Not implemented |
| `GetAccountAuthorizationDetails` | ❌ Not implemented |
| `GetAccountPasswordPolicy` | ❌ Not implemented |
| `GetAccountSummary` | ✅ Implemented |
| `GetContextKeysForCustomPolicy` | ❌ Not implemented |
| `GetContextKeysForPrincipalPolicy` | ❌ Not implemented |
| `GetCredentialReport` | ❌ Not implemented |
| `GetGroup` | ✅ Implemented |
| `GetGroupPolicy` | ✅ Implemented |
| `GetInstanceProfile` | ✅ Implemented |
| `GetLoginProfile` | ❌ Not implemented |
| `GetMFADevice` | ❌ Not implemented |
| `GetOpenIDConnectProvider` | ✅ Implemented |
| `GetOrganizationsAccessReport` | ❌ Not implemented |
| `GetPolicy` | ✅ Implemented |
| `GetPolicyVersion` | ✅ Implemented |
| `GetRole` | ✅ Implemented |
| `GetRolePolicy` | ✅ Implemented |
| `GetSAMLProvider` | ❌ Not implemented |
| `GetSSHPublicKey` | ❌ Not implemented |
| `GetServerCertificate` | ❌ Not implemented |
| `GetServiceLastAccessedDetails` | ❌ Not implemented |
| `GetServiceLastAccessedDetailsWithEntities` | ❌ Not implemented |
| `GetServiceLinkedRoleDeletionStatus` | ❌ Not implemented |
| `GetUser` | ✅ Implemented |
| `GetUserPolicy` | ✅ Implemented |
| `ListAccessKeys` | ✅ Implemented |
| `ListAccountAliases` | ❌ Not implemented |
| `ListAttachedGroupPolicies` | ✅ Implemented |
| `ListAttachedRolePolicies` | ✅ Implemented |
| `ListAttachedUserPolicies` | ✅ Implemented |
| `ListEntitiesForPolicy` | ❌ Not implemented |
| `ListGroupPolicies` | ✅ Implemented |
| `ListGroups` | ✅ Implemented |
| `ListGroupsForUser` | ✅ Implemented |
| `ListInstanceProfileTags` | ✅ Implemented |
| `ListInstanceProfiles` | ✅ Implemented |
| `ListInstanceProfilesForRole` | ✅ Implemented |
| `ListMFADeviceTags` | ❌ Not implemented |
| `ListMFADevices` | ❌ Not implemented |
| `ListOpenIDConnectProviderTags` | ✅ Implemented |
| `ListOpenIDConnectProviders` | ✅ Implemented |
| `ListPolicies` | ✅ Implemented |
| `ListPoliciesGrantingServiceAccess` | ❌ Not implemented |
| `ListPolicyTags` | ✅ Implemented |
| `ListPolicyVersions` | ✅ Implemented |
| `ListRolePolicies` | ✅ Implemented |
| `ListRoleTags` | ✅ Implemented |
| `ListRoles` | ✅ Implemented |
| `ListSAMLProviderTags` | ❌ Not implemented |
| `ListSAMLProviders` | ❌ Not implemented |
| `ListSSHPublicKeys` | ❌ Not implemented |
| `ListServerCertificateTags` | ❌ Not implemented |
| `ListServerCertificates` | ❌ Not implemented |
| `ListServiceSpecificCredentials` | ❌ Not implemented |
| `ListSigningCertificates` | ❌ Not implemented |
| `ListUserPolicies` | ✅ Implemented |
| `ListUserTags` | ✅ Implemented |
| `ListUsers` | ✅ Implemented |
| `ListVirtualMFADevices` | ❌ Not implemented |
| `PutGroupPolicy` | ✅ Implemented |
| `PutRolePermissionsBoundary` | ❌ Not implemented |
| `PutRolePolicy` | ✅ Implemented |
| `PutUserPermissionsBoundary` | ❌ Not implemented |
| `PutUserPolicy` | ✅ Implemented |
| `RemoveClientIDFromOpenIDConnectProvider` | ❌ Not implemented |
| `RemoveRoleFromInstanceProfile` | ✅ Implemented |
| `RemoveUserFromGroup` | ✅ Implemented |
| `ResetServiceSpecificCredential` | ❌ Not implemented |
| `ResyncMFADevice` | ❌ Not implemented |
| `SetDefaultPolicyVersion` | ❌ Not implemented |
| `SetSecurityTokenServicePreferences` | ❌ Not implemented |
| `SimulateCustomPolicy` | ❌ Not implemented |
| `SimulatePrincipalPolicy` | ❌ Not implemented |
| `TagInstanceProfile` | ✅ Implemented |
| `TagMFADevice` | ❌ Not implemented |
| `TagOpenIDConnectProvider` | ✅ Implemented |
| `TagPolicy` | ✅ Implemented |
| `TagRole` | ✅ Implemented |
| `TagSAMLProvider` | ❌ Not implemented |
| `TagServerCertificate` | ❌ Not implemented |
| `TagUser` | ✅ Implemented |
| `UntagInstanceProfile` | ✅ Implemented |
| `UntagMFADevice` | ❌ Not implemented |
| `UntagOpenIDConnectProvider` | ✅ Implemented |
| `UntagPolicy` | ✅ Implemented |
| `UntagRole` | ✅ Implemented |
| `UntagSAMLProvider` | ❌ Not implemented |
| `UntagServerCertificate` | ❌ Not implemented |
| `UntagUser` | ✅ Implemented |
| `UpdateAccessKey` | ✅ Implemented |
| `UpdateAccountPasswordPolicy` | ❌ Not implemented |
| `UpdateAssumeRolePolicy` | ✅ Implemented |
| `UpdateGroup` | ❌ Not implemented |
| `UpdateLoginProfile` | ❌ Not implemented |
| `UpdateOpenIDConnectProviderThumbprint` | ❌ Not implemented |
| `UpdateRole` | ✅ Implemented |
| `UpdateRoleDescription` | ❌ Not implemented |
| `UpdateSAMLProvider` | ❌ Not implemented |
| `UpdateSSHPublicKey` | ❌ Not implemented |
| `UpdateServerCertificate` | ❌ Not implemented |
| `UpdateServiceSpecificCredential` | ❌ Not implemented |
| `UpdateSigningCertificate` | ❌ Not implemented |
| `UpdateUser` | ❌ Not implemented |
| `UploadSSHPublicKey` | ❌ Not implemented |
| `UploadServerCertificate` | ❌ Not implemented |
| `UploadSigningCertificate` | ❌ Not implemented |
