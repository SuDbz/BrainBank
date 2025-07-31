# API Versioning with Bazel

This page explains how to structure Bazel BUILD files for multi-version APIs, with examples and best practices.

## Table of Contents
- [Directory Structure](#directory-structure)
- [BUILD.bazel for Each Version](#buildbazel-for-each-version)
- [Parent BUILD.bazel Features](#parent-buildbazel-features)
- [Build Commands](#build-commands)
- [Usage in Go Code](#usage-in-go-code)
- [Best Practices](#best-practices)
- [Benefits Table](#benefits-table)
- [Next: Migration and Compatibility](./migration.md)

---

## Directory Structure
```
proto/health/
├── BUILD.bazel           # Parent build file with aliases and combined targets
├── v1/
│   ├── health.proto      # V1 proto definition
│   └── BUILD.bazel       # V1-specific build rules
└── v2/
    ├── health.proto      # V2 proto definition  
    └── BUILD.bazel       # V2-specific build rules
```

## BUILD.bazel for Each Version
- V1 uses simple names (`health_proto`, `health`)
- V2 uses versioned names (`health_v2_proto`, `health_v2`)
- Each version has its own `go_proto_library` and `go_library`

## Parent BUILD.bazel Features
- **Version Aliases:** `health_stable`, `health_latest`
- **Combined Targets:** `health_all_versions` for migration
- **Version-Specific OpenAPI Generation**
- **Merged OpenAPI Specification**

## Build Commands
| Command | Description |
|---------|-------------|
| `bazel build //proto/health/v1:health` | Build only V1 |
| `bazel build //proto/health/v2:health_v2` | Build only V2 |
| `bazel build //proto/health:health_all_versions` | Build both versions |
| `bazel build //proto/health:health_stable` | Build stable alias |
| `bazel build //proto/health:health_latest` | Build latest alias |
| `bazel build //proto/health:health_openapi_v1` | Generate V1 OpenAPI |
| `bazel build //proto/health:health_openapi_v2` | Generate V2 OpenAPI |
| `bazel build //proto/health:health_openapi_merged` | Generate merged OpenAPI |

## Usage in Go Code
```go
import (
    healthv1 "github.com/example/grpc-gateway-demo/generated/go/ecommerce/health/v1"
    healthv2 "github.com/example/grpc-gateway-demo/generated/go/ecommerce/health/v2"
)

func main() {
    v1Client := healthv1.NewHealthServiceClient(conn)
    v2Client := healthv2.NewHealthServiceClient(conn)
}
```

## Best Practices
- Use versioned names and import paths
- Keep versions independent
- Use aliases for migration
- Test both versions

## Benefits Table
| Benefit                | Description                                  |
|------------------------|----------------------------------------------|
| Clear Separation       | Each version has its own build rules         |
| Backward Compatibility | V1 remains unchanged and functional          |
| Easy Migration         | Aliases allow gradual migration              |
| Parallel Development   | Teams can work on different versions         |
| Flexible Deployment    | Deploy services with different version combos|
| Documentation          | Generate separate or merged API docs         |
| Testing                | Test versions independently or together      |

---
**Next:** [Migration and Compatibility](./migration.md)
