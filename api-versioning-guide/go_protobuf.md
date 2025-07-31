# API Versioning in Go + Protobuf

This page covers practical API versioning for Go projects using Protobuf, including code structure, best practices, and examples.

## Table of Contents
- [Versioned Protobuf Packages](#versioned-protobuf-packages)
- [Go Code Structure](#go-code-structure)
- [Handling Breaking Changes](#handling-breaking-changes)
- [Deprecation Communication](#deprecation-communication)
- [Example Directory and Structs](#example-directory-and-structs)
- [Next: API Versioning with Bazel](./bazel.md)

---

## Versioned Protobuf Packages
- Use versioned proto packages: `package users.v1;`, `package users.v2;`
- Directory: `proto/users/v1/users.proto`, `proto/users/v2/users.proto`

## Go Code Structure
- Generate Go code into versioned packages: `usersv1`, `usersv2`
- Keep service implementations separate for each version
- Place shared logic in `internal/common/`

## Handling Breaking Changes
- Breaking changes require a new version (e.g., v2)
- Old and new versions coexist
- Never change the contract of existing messages/services in v1

## Deprecation Communication
- Announce deprecation in docs and changelogs
- Add deprecation warnings in API responses (e.g., HTTP header: `Deprecation: true`)
- Mark deprecated fields in proto with `deprecated = true`

## Example Directory and Structs

```
proto/
  users/
    v1/
      users.proto
    v2/
      users.proto
internal/
  users/
    service.go
  common/
    validation.go
```

**Proto Example:**
```proto
syntax = "proto3";
package users.v1;

service UserService {
  rpc GetUser(GetUserRequest) returns (GetUserResponse);
}

message GetUserRequest {
  string id = 1;
}

message GetUserResponse {
  string id = 1;
  string name = 2;
}
```

---
**Next:** [API Versioning with Bazel](./bazel.md)
