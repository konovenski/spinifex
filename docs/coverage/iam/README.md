---
title: "IAM API Coverage"
seoTitle: "AWS IAM API Operation Coverage on Spinifex — Spinifex Docs"
description: "The AWS IAM API operations Spinifex implements, covering users, roles, policies, groups, instance profiles and the access keys that authenticate them."
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
