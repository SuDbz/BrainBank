# Complete Bazel Guide for Go and Protocol Buffers

A comprehensive guide to understanding and using Bazel for Go development with Protocol Buffers, designed for developers of all skill levels.

## Table of Contents
1. [What is Bazel? (Complete Introduction)](#what-is-bazel-complete-introduction)
2. [Why Choose Bazel Over Other Build Tools?](#why-choose-bazel-over-other-build-tools)
3. [Core Concepts Explained](#core-concepts-explained)
4. [Setting Up Your First Bazel Project](#setting-up-your-first-bazel-project)
5. [Understanding Bazel Rules for Go Projects](#understanding-bazel-rules-for-go-projects)
    - [proto_library - Protocol Buffer Definitions](#proto_library---protocol-buffer-definitions)
    - [go_proto_library - Go Code from Protobuf](#go_proto_library---go-code-from-protobuf)
    - [go_library - Reusable Go Packages](#go_library---reusable-go-packages)
    - [go_binary - Executable Programs](#go_binary---executable-programs)
    - [go_test - Testing Your Code](#go_test---testing-your-code)
    - [protoc_gen_swagger - API Documentation](#protoc_gen_swagger---api-documentation)
    - [genrule - Custom Build Commands](#genrule---custom-build-commands)
6. [Understanding Key Bazel Attributes](#understanding-key-bazel-attributes)
    - [visibility - Access Control](#visibility---access-control)
    - [importpath - Go Package Paths](#importpath---go-package-paths)
    - [tags - Metadata and Filtering](#tags---metadata-and-filtering)
    - [embed - Including Libraries](#embed---including-libraries)
    - [deps - Dependencies](#deps---dependencies)
7. [Running Development Servers with Bazel](#running-development-servers-with-bazel)
    - [py_binary with entry_point](#py_binary-with-entry_point)
    - [local_server_test](#local_server_test)
    - [http_archive with server](#http_archive-with-server)
8. [Complete Working Examples](#complete-working-examples)
    - [Simple Go Project with Bazel](#simple-go-project-with-bazel)
    - [Go + Protocol Buffers Project](#go--protocol-buffers-project)
    - [Microservice with gRPC and Swagger](#microservice-with-grpc-and-swagger)
9. [Automating with Gazelle](#automating-with-gazelle)
10. [Managing Dependencies Like a Pro](#managing-dependencies-like-a-pro)
11. [Understanding Workspace Structure](#understanding-workspace-structure)
12. [Advanced Tips and Best Practices](#advanced-tips-and-best-practices)
13. [Common Patterns and Troubleshooting](#common-patterns-and-troubleshooting)
14. [How to Run Each Bazel Rule and Command](#how-to-run-each-bazel-rule-and-command)
15. [Resources and Further Learning](#resources-and-further-learning)

---

## What is Bazel? (Complete Introduction)

**Bazel** is Google's open-source build system that was originally developed to handle Google's massive codebase. Think of it as a sophisticated replacement for traditional build tools like Make, Maven, or Gradle, but designed for the modern era of software development.

### The Problem Bazel Solves

Imagine you're working on a large software project with:
- Multiple programming languages (Go, Java, Python, etc.)
- Hundreds of developers making changes daily
- Complex dependencies between different parts of the code
- Need for fast, reliable builds and tests

Traditional build tools struggle with these challenges. Bazel was designed to excel in exactly these scenarios.

### Key Benefits:

1. **Speed**: Bazel only rebuilds what's changed and caches everything else
2. **Correctness**: Reproducible builds ensure the same input always produces the same output
3. **Scale**: Handles projects with millions of lines of code
4. **Multi-language**: Build Go, Java, Python, C++, and more in the same project
5. **Remote Execution**: Distribute builds across multiple machines
6. **Incremental**: Only test and build what's affected by your changes

### How Bazel Works (Simple Explanation)

Think of Bazel like a smart chef in a restaurant:

1. **Recipe Book (BUILD files)**: Each directory has a "recipe" describing how to build the code in that directory
2. **Ingredients (Source files)**: Your Go files, proto files, etc.
3. **Dependencies (deps)**: What other "dishes" this recipe needs
4. **Smart Caching**: The chef remembers what they've already cooked and reuses it
5. **Parallel Cooking**: Multiple dishes can be prepared simultaneously

## Why Choose Bazel Over Other Build Tools?

| Feature | Bazel | Go Modules | Make | Maven/Gradle |
|---------|-------|------------|------|--------------|
| Multi-language support | ✅ Excellent | ❌ Go only | ⚠️ Basic | ⚠️ Limited |
| Incremental builds | ✅ Excellent | ✅ Good | ⚠️ Manual | ✅ Good |
| Remote caching | ✅ Built-in | ❌ No | ❌ No | ⚠️ Plugins |
| Reproducible builds | ✅ Guaranteed | ✅ Good | ❌ No | ⚠️ Configurable |
| Large monorepo support | ✅ Excellent | ⚠️ Limited | ❌ Poor | ⚠️ Limited |
| Learning curve | ⚠️ Steep | ✅ Easy | ✅ Easy | ⚠️ Moderate |

**Choose Bazel when:**
- You have a multi-language project
- You're building a monorepo
- Build speed and reproducibility are critical
- You have a team of 10+ developers

**Stick with Go modules when:**
- You have a simple Go-only project
- You're just starting out
- Quick setup is more important than advanced features

## Core Concepts Explained

### 1. Workspace
The **workspace** is the root directory of your project. It contains a special file called `WORKSPACE` that tells Bazel "this directory and everything below it is a Bazel project."

```
my-project/           ← This is your workspace root
├── WORKSPACE         ← This file makes it a Bazel workspace
├── BUILD.bazel       ← Build rules for this directory
├── go/
│   ├── BUILD.bazel   ← Build rules for the go/ directory
│   └── main.go
└── protos/
    ├── BUILD.bazel   ← Build rules for the protos/ directory
    └── api.proto
```

### 2. Targets
A **target** is something that Bazel can build, test, or run. Every target has:
- A **name** (how you refer to it)
- A **type** (what kind of thing it is: library, binary, test, etc.)
- **Inputs** (source files, dependencies)
- **Outputs** (what gets created)

Examples:
```python
# This creates a target named "api_proto"
proto_library(
    name = "api_proto",    # ← Target name
    srcs = ["api.proto"],  # ← Input files
)

# This creates a target named "server"
go_binary(
    name = "server",           # ← Target name
    embed = [":server_lib"],   # ← Dependencies
)
```

### 3. Rules
**Rules** are templates that tell Bazel how to build certain types of targets. Think of them as "build recipes":

- `go_library` rule: "Here's how to build a Go library"
- `go_binary` rule: "Here's how to build a Go executable"
- `proto_library` rule: "Here's how to process Protocol Buffer files"

### 4. Labels
**Labels** are how you refer to targets. They follow this pattern:
```
//path/to/package:target_name
```

Examples:
- `//protos:api_proto` - The "api_proto" target in the "protos" package
- `:server_lib` - The "server_lib" target in the current package
- `@external_repo//pkg:target` - A target from an external repository

### 5. BUILD Files
**BUILD files** (or BUILD.bazel) contain the rules that define targets. One BUILD file per directory that contains code.

## Setting Up Your First Bazel Project

Let's create a simple "Hello World" project to understand the basics:

### Step 1: Create the Project Structure
```bash
mkdir my-first-bazel-project
cd my-first-bazel-project
```

### Step 2: Create the WORKSPACE File
```python
# WORKSPACE
workspace(name = "my_first_bazel_project")

# Load Go rules
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Download rules_go
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "...",  # Use latest SHA from rules_go releases
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.39.1/rules_go-v0.39.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.39.1/rules_go-v0.39.1.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains(version = "1.19.5")
```

### Step 3: Create Your Go Code
```go
// main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Bazel!")
}
```

### Step 4: Create the BUILD File
```python
# BUILD.bazel
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "hello_lib",
    srcs = ["main.go"],
    importpath = "github.com/example/hello",
)

go_binary(
    name = "hello",
    embed = [":hello_lib"],
)
```

### Step 5: Build and Run
```bash
# Build the binary
bazel build //:hello

# Run the binary
bazel run //:hello
```

**What just happened?**
1. Bazel read your WORKSPACE file and downloaded the Go rules
2. It found your BUILD file and understood you want to build a Go binary
3. It compiled your Go code and created an executable
4. The `bazel run` command executed your program

## Understanding Bazel Rules for Go Projects

### proto_library - Protocol Buffer Definitions

The `proto_library` rule tells Bazel how to handle Protocol Buffer (.proto) files. Think of it as saying "here are some API definitions that other parts of my project might want to use."

**What it does:**
- Validates your .proto files for syntax errors
- Makes proto definitions available to other rules
- Handles dependencies between proto files

**When to use it:**
- You have .proto files defining APIs, data structures, or services
- You want to generate code in multiple languages from the same proto
- You need to share proto definitions across different parts of your project

**Detailed example:**
```python
proto_library(
    name = "user_api_proto",           # Target name - how other rules refer to this
    srcs = [                          # The .proto files to include
        "user.proto",
        "auth.proto",
    ],
    deps = [                          # Other proto_library targets this depends on
        "//common/protos:timestamp_proto",
        "@googleapis//google/api:annotations_proto",
    ],
    visibility = ["//visibility:public"], # Who can use this proto definition
)
```

**Real-world example - user.proto:**
```protobuf
syntax = "proto3";

package user.v1;

import "google/api/annotations.proto";
import "common/protos/timestamp.proto";

service UserService {
  rpc GetUser(GetUserRequest) returns (User) {
    option (google.api.http) = {
      get: "/v1/users/{user_id}"
    };
  }
}

message User {
  string user_id = 1;
  string email = 2;
  common.Timestamp created_at = 3;
}

message GetUserRequest {
  string user_id = 1;
}
```

### go_proto_library - Go Code from Protobuf

The `go_proto_library` rule generates Go code from your proto definitions. It's like having a translator that converts your API specification into actual Go structs and interfaces.

**What it does:**
- Generates Go structs from proto messages
- Creates Go interfaces from proto services (for gRPC)
- Handles imports and dependencies automatically
- Produces a Go library you can import in your code

**When to use it:**
- You want to use proto-defined messages in Go code
- You're building gRPC services
- You need type-safe data structures based on your API definitions

**Detailed example:**
```python
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

go_proto_library(
    name = "user_api_go_proto",              # Target name
    proto = ":user_api_proto",               # The proto_library to generate from
    importpath = "github.com/mycompany/myproject/api/user/v1",  # Go import path
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],  # Include gRPC service code
    deps = [                                 # Go dependencies for generated code
        "//common/protos:timestamp_go_proto",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
    ],
    visibility = ["//visibility:public"],
)
```

**What gets generated:**
```go
// This Go code is automatically generated
package userv1

type User struct {
    UserId    string    `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3"`
    Email     string    `protobuf:"bytes,2,opt,name=email,proto3"`
    CreatedAt *Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3"`
}

type UserServiceClient interface {
    GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
}
```

### go_library - Reusable Go Packages

The `go_library` rule is the foundation of Go development in Bazel. It compiles Go source files into a reusable package that other parts of your project can import.

**What it does:**
- Compiles Go source files into a package
- Manages dependencies between Go packages
- Provides a reusable unit that binaries and tests can use
- Handles Go module imports and vendoring

**When to use it:**
- You have Go code that multiple other packages need to import
- You want to organize your code into logical, reusable modules
- You're following good software engineering practices (separation of concerns)

**Detailed example:**
```python
go_library(
    name = "user_service_lib",              # Target name
    srcs = [                               # Go source files
        "service.go",
        "repository.go", 
        "handler.go",
    ],
    importpath = "github.com/mycompany/myproject/services/user",  # Go import path
    deps = [                               # Dependencies on other Go libraries
        ":user_api_go_proto",             # Our generated proto code
        "//pkg/database:database_lib",     # Database utilities
        "//pkg/auth:auth_lib",            # Authentication utilities
        "@com_github_gorilla_mux//:mux",  # External HTTP router
        "@org_golang_google_grpc//:grpc", # gRPC library
    ],
    visibility = [                         # Who can depend on this library
        "//services:__subpackages__",      # All packages under services/
        "//cmd:__subpackages__",           # All packages under cmd/
    ],
    tags = ["service"],                    # Metadata for tooling
)
```

**Example Go code that would use this library:**
```go
// In another package's Go file
package main

import (
    "github.com/mycompany/myproject/services/user"
    userv1 "github.com/mycompany/myproject/api/user/v1"
)

func main() {
    userService := user.NewService()
    user, err := userService.GetUser(&userv1.GetUserRequest{UserId: "123"})
    // ... handle user and error
}
```

### go_binary - Executable Programs

The `go_binary` rule creates executable programs from Go libraries. This is what you use to build the actual applications that users run.

**What it does:**
- Links Go libraries together into a single executable
- Handles the main() function and program entry point
- Creates platform-specific binaries (Linux, macOS, Windows)
- Manages static vs dynamic linking

**When to use it:**
- You need to create a runnable program (CLI tool, server, etc.)
- You have a main() function that serves as an entry point
- You want to deploy or distribute an executable

**Detailed example:**
```python
go_binary(
    name = "user_server",                  # Target name (also the binary name)
    embed = [":user_service_lib"],         # The main library containing main()
    deps = [                              # Additional dependencies if needed
        "//pkg/config:config_lib",
    ],
    goos = "linux",                       # Target operating system (optional)
    goarch = "amd64",                     # Target architecture (optional)
    pure = "on",                          # Static linking for better deployment
    visibility = ["//visibility:public"],
    tags = ["manual"],                    # Don't build unless explicitly requested
    x_defs = {                           # Compile-time variable injection
        "main.version": "{BUILD_VERSION}",
        "main.buildTime": "{BUILD_TIMESTAMP}",
    },
)
```

**Example main.go that would be embedded:**
```go
// main.go
package main

import (
    "flag"
    "log"
    "net/http"
    
    "github.com/mycompany/myproject/services/user"
)

var (
    version   = "dev"    // Can be overridden with x_defs
    buildTime = "unknown" // Can be overridden with x_defs
)

func main() {
    port := flag.String("port", "8080", "Server port")
    flag.Parse()
    
    log.Printf("Starting user server version %s (built %s)", version, buildTime)
    
    userService := user.NewService()
    http.HandleFunc("/users", userService.HandleUsers)
    log.Fatal(http.ListenAndServe(":"+*port, nil))
}
```

### go_test - Testing Your Code

The `go_test` rule runs Go tests. It's essential for maintaining code quality and ensuring your changes don't break existing functionality.

**What it does:**
- Compiles and runs Go test files (*_test.go)
- Provides test isolation and parallelization
- Handles test dependencies and setup
- Integrates with Bazel's caching for fast test runs

**When to use it:**
- You have unit tests for your Go code
- You want integration tests for your services
- You need to verify your code works correctly

**Detailed example:**
```python
go_test(
    name = "user_service_test",            # Target name
    srcs = [                              # Test files
        "service_test.go",
        "repository_test.go",
        "integration_test.go",
    ],
    embed = [":user_service_lib"],         # The library being tested
    deps = [                              # Test-specific dependencies
        "@com_github_stretchr_testify//assert",
        "@com_github_stretchr_testify//mock", 
        "//testutils:database_testutils",
    ],
    data = [                              # Test data files
        "testdata/users.json",
        "testdata/schema.sql",
    ],
    env = {                               # Environment variables for tests
        "TEST_DATABASE_URL": "sqlite://memory",
        "LOG_LEVEL": "debug",
    },
    tags = [
        "unit",                           # Tag for filtering
        "requires-network",               # Metadata about test requirements
    ],
    size = "medium",                      # Test size hint (small/medium/large)
    timeout = "short",                    # How long tests can run
)
```

**Example test file:**
```go
// service_test.go
package user

import (
    "context"
    "testing"
    
    "github.com/stretchr/testify/assert"
    userv1 "github.com/mycompany/myproject/api/user/v1"
)

func TestUserService_GetUser(t *testing.T) {
    // Setup
    service := NewService()
    
    // Test
    user, err := service.GetUser(context.Background(), &userv1.GetUserRequest{
        UserId: "test-user-123",
    })
    
    // Assertions
    assert.NoError(t, err)
    assert.NotNil(t, user)
    assert.Equal(t, "test-user-123", user.UserId)
}
```

### protoc_gen_swagger - API Documentation

The `protoc_gen_swagger` rule generates OpenAPI/Swagger documentation from your Protocol Buffer service definitions. This creates interactive API documentation that developers can use to understand and test your APIs.

**What it does:**
- Parses gRPC service definitions from proto files
- Generates OpenAPI 3.0/Swagger 2.0 specification files
- Creates interactive documentation websites
- Includes HTTP annotations for REST API mapping

**When to use it:**
- You have gRPC services that also expose HTTP endpoints
- You want to generate API documentation automatically
- You need to provide API specifications to frontend developers or partners
- You want to ensure your documentation stays in sync with your code

**Detailed example:**
```python
load("@rules_proto_grpc//doc:defs.bzl", "doc_swagger_compile")

doc_swagger_compile(
    name = "user_api_swagger",             # Target name
    protos = [":user_api_proto"],          # Proto files to document
    options = [                           # Swagger generation options
        "logtostderr=true",
        "allow_merge=true",
        "merge_file_name=user_api",
    ],
    visibility = ["//visibility:public"],
)

# Alternative using protoc-gen-openapiv2
load("@grpc_ecosystem_grpc_gateway//protoc-gen-openapiv2:defs.bzl", "protoc_gen_openapiv2")

protoc_gen_openapiv2(
    name = "user_api_openapi",
    proto = ":user_api_proto",
    options = {
        "allow_merge": "true",
        "merge_file_name": "user_api",
        "json_names_for_fields": "false",
    },
)
```

**Generated Swagger/OpenAPI output example:**
```json
{
  "swagger": "2.0",
  "info": {
    "title": "User API",
    "version": "1.0"
  },
  "paths": {
    "/v1/users/{user_id}": {
      "get": {
        "summary": "Get user by ID",
        "parameters": [
          {
            "name": "user_id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "User found",
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "User": {
      "type": "object",
      "properties": {
        "user_id": {"type": "string"},
        "email": {"type": "string"},
        "created_at": {"$ref": "#/definitions/Timestamp"}
      }
    }
  }
}
```

### genrule - Custom Build Commands

The `genrule` is Bazel's most flexible rule that allows you to run arbitrary shell commands to generate files. Think of it as a way to integrate custom build steps, scripts, or tools that don't have dedicated Bazel rules.

**What it does:**
- Executes shell commands as part of the build process
- Takes input files and produces output files
- Provides a bridge between Bazel and external tools
- Enables custom code generation and file processing

**When to use it:**
- You need to run custom scripts or tools during the build
- You want to generate code from templates or configuration files
- You need to process files in ways not covered by existing rules
- You want to integrate legacy build tools into Bazel
- You need to create configuration files or documentation during build

**Basic syntax:**
```python
genrule(
    name = "rule_name",
    srcs = ["input_files"],     # Input files
    outs = ["output_files"],    # Output files that will be generated
    cmd = "shell_command",      # Command to run
    tools = ["executables"],    # Additional tools needed
)
```

**Detailed examples:**

**1. Simple file generation:**
```python
genrule(
    name = "generate_version",
    outs = ["version.go"],
    cmd = '''
cat > $@ << 'EOF'
package main

const Version = "$(BUILD_VERSION)"
const BuildTime = "$(BUILD_TIMESTAMP)"
EOF
    ''',
)
```

**2. Processing configuration files:**
```python
genrule(
    name = "process_config",
    srcs = [
        "config.template.yaml",
        "//deployment:environment_vars.txt",
    ],
    outs = ["config.yaml"],
    cmd = """
        envsubst < $(location config.template.yaml) > $@ && \\
        echo "# Generated at build time" >> $@
    """,
    tools = ["@envsubst_tool//:envsubst"],  # External tool dependency
)
```

**3. Code generation from protobuf (custom generator):**
```python
genrule(
    name = "generate_custom_client",
    srcs = [":user_api_proto"],
    outs = [
        "user_client.go",
        "user_models.go",
    ],
    cmd = """
        $(location //tools:proto_gen_client) \\
            --input=$(location :user_api_proto) \\
            --output_dir=$(@D) \\
            --package=client
    """,
    tools = ["//tools:proto_gen_client"],
    visibility = ["//visibility:public"],
)
```

**4. Documentation generation:**
```python
genrule(
    name = "generate_docs",
    srcs = [
        "//api:all_protos",
        "README.template.md",
    ],
    outs = [
        "api_docs.md",
        "README.md",
    ],
    cmd = """
        # Generate API documentation
        $(location @protoc_gen_doc//:protoc-gen-doc) \\
            --doc_out=$(@D) \\
            --doc_opt=markdown,api_docs.md \\
            $(locations //api:all_protos)
        
        # Process README template
        sed 's/{{VERSION}}/$(BUILD_VERSION)/g' $(location README.template.md) > $(@D)/README.md
    """,
    tools = ["@protoc_gen_doc//:protoc-gen-doc"],
)
```

**5. Asset bundling and compression:**
```python
genrule(
    name = "bundle_assets",
    srcs = glob([
        "static/**/*.js",
        "static/**/*.css",
        "static/**/*.html",
    ]),
    outs = [
        "assets.tar.gz",
        "assets_manifest.json",
    ],
    cmd = """
        # Create asset bundle
        tar -czf $(@D)/assets.tar.gz -C static .
        
        # Generate manifest
        cat > $(@D)/assets_manifest.json << 'EOF'
{
    "version": "$(BUILD_VERSION)",
    "files": [
$(for file in $(SRCS); do echo "        \\"$$file\\","; done | sed '$$s/,$$//')
    ],
    "build_time": "$(BUILD_TIMESTAMP)"
}
EOF
    """,
)
```

**6. Database schema generation:**
```python
genrule(
    name = "generate_schema",
    srcs = [
        "//database/migrations:all_migrations",
        "schema_template.sql",
    ],
    outs = [
        "schema.sql",
        "schema_version.txt",
    ],
    cmd = """
        # Combine all migrations into single schema
        cat $(locations //database/migrations:all_migrations) > $(@D)/schema.sql
        
        # Add template content
        cat $(location schema_template.sql) >> $(@D)/schema.sql
        
        # Generate version file
        echo "Schema version: $(BUILD_VERSION)" > $(@D)/schema_version.txt
        echo "Generated: $(BUILD_TIMESTAMP)" >> $(@D)/schema_version.txt
    """,
    visibility = ["//database:__subpackages__"],
)
```

**Key genrule features and variables:**

**Built-in variables:**
- `$@` - Single output file (only works when outs has exactly one file)
- `$(@D)` - Directory containing the output files
- `$(SRCS)` - All source files
- `$(location target)` - Path to a specific file or target
- `$(locations target)` - Paths to all files from a target (if multiple)

**Build variables:**
- `$(BUILD_VERSION)` - Build version (if configured)
- `$(BUILD_TIMESTAMP)` - When the build was run
- `$(TARGET_CPU)` - Target CPU architecture

**Advanced genrule patterns:**

**1. Conditional generation:**
```python
genrule(
    name = "conditional_config",
    srcs = [
        "config.prod.yaml",
        "config.dev.yaml",
    ],
    outs = ["config.yaml"],
    cmd = select({
        "//config:prod": "cp $(location config.prod.yaml) $@",
        "//config:dev": "cp $(location config.dev.yaml) $@",
        "//conditions:default": "cp $(location config.dev.yaml) $@",
    }),
)
```

**2. Multi-step processing:**
```python
genrule(
    name = "multi_step_generation",
    srcs = ["input.json"],
    outs = [
        "output.go",
        "output_test.go",
    ],
    cmd = """
        # Step 1: Validate input
        $(location //tools:json_validator) $(location input.json) || exit 1
        
        # Step 2: Generate main code
        $(location //tools:go_generator) \\
            --input=$(location input.json) \\
            --output=$(@D)/output.go || exit 1
        
        # Step 3: Generate tests
        $(location //tools:test_generator) \\
            --input=$(location input.json) \\
            --output=$(@D)/output_test.go || exit 1
    """,
    tools = [
        "//tools:json_validator",
        "//tools:go_generator", 
        "//tools:test_generator",
    ],
)
```

**3. Custom script execution:**
```python
genrule(
    name = "run_custom_script",
    srcs = ["script.sh"],
    outs = ["output.txt"],
    cmd = "bash $(location script.sh) > $@",
    tools = ["//tools:bash"],
)
```

**4. File transformation:**
```python
genrule(
    name = "transform_data",
    srcs = ["input.csv"],
    outs = ["output.json"],
    cmd = """
        echo "[" > $@
        cat $(SRCS) | sed 's/,/","/g' | sed 's/^/"/' | sed 's/$$/"/' >> $@
        echo "]" >> $@
    """,
)
```

**5. Configuration file generation:**
```python
genrule(
    name = "generate_config",
    srcs = ["config.template"],
    outs = ["config.yaml"],
    cmd = """
        echo "server:" > $@
        echo "  port: 8080" >> $@
        echo "  host: localhost" >> $@
    """,
)
```

**6. Asset compression:**
```python
genrule(
    name = "compress_assets",
    srcs = glob(["static/**/*"]),
    outs = ["assets.zip"],
    cmd = "zip -r $@ $<",
)
```

**Best practices for genrule:**

1. **Keep commands simple and focused:**
   ```python
   # ✅ Good: single responsibility
   genrule(
       name = "generate_version",
       cmd = "echo 'const Version = \"1.0.0\"' > $@",
   )
   
   # ❌ Bad: too complex for genrule
   genrule(
       cmd = "complex_script.sh && compile_stuff && run_tests && deploy",
   )
   ```

2. **Use explicit file paths:**
   ```python
   # ✅ Good: explicit file references
   genrule(
       cmd = "$(location //tools:generator) $(location input.txt) > $@",
   )
   
   # ❌ Bad: relying on implicit paths
   genrule(
       cmd = "generator input.txt > output.txt",  # May not find files
   )
   ```

3. **Handle errors properly:**
   ```python
   genrule(
       cmd = """
           $(location //tools:validator) $(location input.json) || exit 1
           $(location //tools:generator) $(location input.json) > $@ || exit 1
       """,
   )
   ```

4. **Use tools attribute for executables:**
   ```python
   genrule(
       tools = [
           "//tools:custom_generator",    # Your custom tools
           "@nodejs//:bin/node",          # External tools
       ],
       cmd = "$(location //tools:custom_generator) --input=$(SRCS) --output=$@",
   )
   ```

**Integration with Go projects:**
```python
# Generate Go code that can be used by go_library
genrule(
    name = "generate_go_constants",
    srcs = ["constants.yaml"],
    outs = ["constants.go"],
    cmd = """
        echo 'package constants' > $@
        echo '' >> $@
        $(location //tools:yaml_to_go) $(location constants.yaml) >> $@
    """,
    tools = ["//tools:yaml_to_go"],
)

go_library(
    name = "constants_lib",
    srcs = [":generate_go_constants"],  # Use generated file
    importpath = "github.com/mycompany/myproject/constants",
)
```

The `genrule` is incredibly powerful for custom build steps, but should be used judiciously. For common tasks, prefer dedicated rules like `go_library`, `proto_library`, etc., and use `genrule` for the special cases where you need custom processing.

## Understanding Key Bazel Attributes

Understanding Bazel attributes is crucial for effectively organizing and controlling your build. These attributes control how targets interact with each other and how the build system behaves.

### visibility - Access Control

**What is visibility?**
`visibility` is like a security system for your code. It controls which other targets in your workspace can depend on (use) a particular target. This helps maintain clean architecture and prevents unwanted dependencies.

**Why is this important?**
- Prevents accidental dependencies that could create circular imports
- Enforces architectural boundaries (e.g., only allow certain packages to use internal APIs)
- Makes refactoring safer by clearly defining public vs private interfaces
- Helps with large team development by creating clear contracts

**Common visibility patterns:**
```python
go_library(
    name = "public_api",
    visibility = ["//visibility:public"],  # Anyone can use this
)

go_library(
    name = "internal_utils", 
    visibility = ["//visibility:private"], # Only this package can use it
)

go_library(
    name = "service_shared",
    visibility = [
        "//services:__subpackages__",      # Only packages under services/
        "//cmd/servers:__pkg__",           # Only the cmd/servers package
    ],
)

go_library(
    name = "team_shared",
    visibility = [
        "//team/backend/...",              # All packages under team/backend/
        "//team/shared/...",               # All packages under team/shared/
    ],
)
```

**Real-world example:**
```python
# Database package - only accessible to service layer
go_library(
    name = "database_lib",
    srcs = ["db.go", "migrations.go"],
    visibility = [
        "//services:__subpackages__",      # Services can access database
        "//cmd/migrate:__pkg__",           # Migration tool can access it
    ],
    # Frontend packages CANNOT access database directly
)
```

### importpath - Go Package Paths

**What is importpath?**
`importpath` is the string you use in Go `import` statements. It tells the Go compiler how to refer to this package and must be unique across your entire project.

**Why is this critical?**
- Go uses import paths to uniquely identify packages
- Must match what you write in your Go source files
- Affects go.mod files and module resolution
- Determines where generated code gets placed

**Best practices for import paths:**
```python
# ✅ Good: follows standard Go conventions
go_library(
    name = "user_service",
    importpath = "github.com/mycompany/myproject/services/user",
    # Source code can now: import "github.com/mycompany/myproject/services/user"
)

# ✅ Good: matches your repository structure
go_proto_library(
    name = "user_api_go_proto",
    importpath = "github.com/mycompany/myproject/api/user/v1",
    # Source code can now: import userv1 "github.com/mycompany/myproject/api/user/v1"
)

# ❌ Bad: doesn't follow Go conventions
go_library(
    importpath = "my-random-path/stuff",  # Should match repository
)

# ❌ Bad: conflicts with standard library
go_library(
    importpath = "fmt",                   # Will conflict with built-in fmt
)
```

**Versioning with import paths:**
```python
# API v1
go_proto_library(
    importpath = "github.com/mycompany/myproject/api/user/v1",
)

# API v2 (breaking changes)
go_proto_library(
    importpath = "github.com/mycompany/myproject/api/user/v2",
)
```

### tags - Metadata and Filtering

**What are tags?**
`tags` are arbitrary labels you attach to targets. They don't affect the build output but provide metadata that tools (including Bazel itself) can use for filtering, categorization, and special handling.

**Common use cases:**
1. **Build filtering** - exclude certain targets from bulk operations
2. **Test categorization** - group tests by type (unit, integration, e2e)
3. **Deployment stages** - mark targets for different environments
4. **Tool integration** - provide hints to IDEs and other tools

**Practical examples:**
```python
# Don't build unless explicitly requested
go_binary(
    name = "debug_tool",
    tags = ["manual"],                    # Won't be built by "bazel build //..."
)

# Test categorization
go_test(
    name = "unit_tests",
    tags = ["unit", "fast"],              # Run with: bazel test --test_tag_filters=unit
)

go_test(
    name = "integration_tests", 
    tags = ["integration", "slow", "requires-database"],
)

go_test(
    name = "e2e_tests",
    tags = ["e2e", "requires-network", "flaky"],
)

# Environment-specific builds
go_binary(
    name = "server_dev",
    tags = ["dev", "local"],
)

go_binary(
    name = "server_prod",
    tags = ["prod", "release"],
)
```

**Using tags for selective building/testing:**
```bash
# Build only fast tests
bazel test --test_tag_filters=fast //...

# Build everything except manual targets
bazel build --build_tag_filters=-manual //...

# Run only integration tests
bazel test --test_tag_filters=integration //...

# Skip flaky tests
bazel test --test_tag_filters=-flaky //...
```

### embed - Including Libraries

**What is embed?**
`embed` is how you include Go libraries directly into binaries and tests. Think of it as "bake this library's code directly into the final executable" rather than treating it as an external dependency.

**When to use embed vs deps:**
- **embed**: For libraries that contain the main package or are core to the binary/test
- **deps**: For external dependencies that your embedded libraries need

**Key differences:**
```python
go_binary(
    name = "server",
    embed = [":server_lib"],              # This library's code goes INTO the binary
    deps = [                             # These are available to embedded libraries
        "@com_github_gorilla_mux//:mux",
        "//pkg/database:db_lib", 
    ],
)
```

**Detailed examples:**
```python
# Library containing main() function
go_library(
    name = "server_main_lib",
    srcs = ["main.go"],                   # Contains func main()
    deps = [
        ":handlers_lib",
        "@com_github_gorilla_mux//:mux",
    ],
)

# Binary that embeds the main library
go_binary(
    name = "server",
    embed = [":server_main_lib"],         # Embeds main() and related code
    # Note: deps from embedded library are automatically included
)

# Test that embeds the library being tested
go_test(
    name = "handlers_test",
    srcs = ["handlers_test.go"],
    embed = [":handlers_lib"],            # Test code can access unexported functions
    deps = [
        "@com_github_stretchr_testify//assert",
    ],
    data = [                              # Test data files
        "testdata/users.json",
        "testdata/schema.sql",
    ],
    env = {                               # Environment variables for tests
        "TEST_DATABASE_URL": "sqlite://memory",
        "LOG_LEVEL": "debug",
    },
    tags = [
        "unit",                           # Tag for filtering
        "requires-network",               # Metadata about test requirements
    ],
    size = "medium",                      # Test size hint (small/medium/large)
    timeout = "short",                    # How long tests can run
)
```

**Multiple embeds (less common but useful):**
```python
go_binary(
    name = "multi_module_tool",
    embed = [
        ":tool_main_lib",                 # Contains main()
        ":shared_utils_lib",              # Shared utilities
        ":config_lib",                    # Configuration handling
    ],
)
```

### deps - Dependencies

**What are deps?**
`deps` lists other targets that this target needs to build successfully. These are the external libraries, other packages in your project, and generated code that your target depends on.

**Types of dependencies:**
1. **Internal dependencies** - other targets in your workspace
2. **External dependencies** - third-party libraries
3. **Generated dependencies** - proto-generated code, etc.

**Detailed example:**
```python
go_library(
    name = "user_service",
    srcs = ["service.go"],
    deps = [
        "//api/user/v1:user_api_go_proto",  # Generated proto code
        "//pkg/database:db_lib",            # Database utilities
        "//pkg/auth:auth_lib",              # Authentication utilities
        "@com_github_gorilla_mux//:mux",    # External HTTP router
        "@org_golang_google_grpc//:grpc",    # gRPC library
        "@com_github_sirupsen_logrus//:logrus", # Logging
    ],
)
```

**Dependency management best practices:**
```python
# ✅ Good: minimal, specific dependencies
go_library(
    name = "http_client",
    deps = [
        "@com_github_gorilla_mux//:mux",  # Actually used for routing
    ],
)

# ❌ Bad: including unnecessary dependencies  
go_library(
    name = "http_client",
    deps = [
        "@com_github_gorilla_mux//:mux",
        "@org_golang_google_grpc//:grpc",     # Not used in this library
        "@com_github_sirupsen_logrus//:logrus", # Could use standard log
    ],
)
```

**2. Use consistent versioning:**
```bash
# Keep go.mod and Bazel in sync
go mod tidy
bazel run //:gazelle-update-repos

# Regularly update dependencies
go get -u ./...
go mod tidy
bazel run //:gazelle-update-repos
```

**3. Document dependency choices:**
```python
# deps.bzl
def go_dependencies():
    # HTTP router - chosen for performance and simplicity
    go_repository(
        name = "com_github_gorilla_mux",
        importpath = "github.com/gorilla/mux",
        version = "v1.8.0",
    )
    
    # Database driver - PostgreSQL support
    # Note: pinned to v1.10.7 due to connection pooling issue in v1.11.0
    go_repository(
        name = "com_github_lib_pq",
        importpath = "github.com/lib/pq", 
        version = "v1.10.7",
    )
```

**4. Test dependency updates:**
```bash
# Create a script to test dependency updates
#!/bin/bash
set -e

echo "Updating go.mod..."
go get -u ./...
go mod tidy

echo "Updating Bazel dependencies..."
bazel run //:gazelle-update-repos

echo "Running tests..."
bazel test //...

echo "All tests passed! Dependencies updated successfully."
```

## Running Development Servers with Bazel

Bazel provides several ways to run development servers, depending on your language and workflow.

### py_binary with entry_point

Use this rule to run a Python server. It's simple and works well for local development.

**Example:**
```python
py_binary(
    name = "server_app",
    srcs = ["server/app.py"],
    entry_point = "server.app",  # module path, not a file path
)
```
**server/app.py:**
```python
from http.server import HTTPServer, SimpleHTTPRequestHandler

def main():
    server = HTTPServer(("localhost", 8080), SimpleHTTPRequestHandler)
    print("Server started at http://localhost:8080")
    server.serve_forever()

if __name__ == "__main__":
    main()
```
**Run it:**
```bash
bazel run //:server_app
```

### local_server_test

This rule (from [rules_testing](https://github.com/bazelbuild/rules_testing)) lets you spin up a server and run integration tests against it.

**Example:**
```python
load("@rules_testing//local_server:defs.bzl", "local_server_test")

local_server_test(
    name = "server_test",
    server = ":server_app",  # py_binary defined earlier
    readiness_path = "/",
    port = 8080,
    test_script = "server/test_server.py",
    deps = ["@rules_testing//lib:curl"],
)
```
**server/test_server.py:**
```python
def main():
    import requests
    response = requests.get("http://localhost:8080")
    assert response.status_code == 200
    print("Server responded OK.")

if __name__ == "__main__":
    main()
```

### http_archive with server

`http_archive` is used to import external Bazel rules, including those that provide server capabilities (e.g., rules_nodejs for frontend dev servers).

**Example:**
```python
load("@build_bazel_rules_nodejs//:index.bzl", "http_server")

http_server(
    name = "serve_portal",
    data = [":generate_docs"],
    serving_path = "docs",
    port = 8000,
)
```
Use this when you need to run a server provided by a Bazel extension (such as a frontend dev server).

---

**When to use which:**
- Use `py_binary` for simple Python servers and local development.
- Use `local_server_test` for automated integration testing of server endpoints.
- Use `http_archive` with server rules for language-specific or framework-specific dev servers.
