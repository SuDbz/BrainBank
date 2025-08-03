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
7. [Complete Working Examples](#complete-working-examples)
    - [Simple Go Project with Bazel](#simple-go-project-with-bazel)
    - [Go + Protocol Buffers Project](#go--protocol-buffers-project)
    - [Microservice with gRPC and Swagger](#microservice-with-grpc-and-swagger)
8. [Automating with Gazelle](#automating-with-gazelle)
9. [Managing Dependencies Like a Pro](#managing-dependencies-like-a-pro)
10. [Understanding Workspace Structure](#understanding-workspace-structure)
11. [Advanced Tips and Best Practices](#advanced-tips-and-best-practices)
12. [Common Patterns and Troubleshooting](#common-patterns-and-troubleshooting)
13. [How to Run Each Bazel Rule and Command](#how-to-run-each-bazel-rule-and-command)
14. [Resources and Further Learning](#resources-and-further-learning)

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
        $(location //tools:json_validator) $(location input.json)
        
        # Step 2: Generate main code
        $(location //tools:go_generator) \\
            --input=$(location input.json) \\
            --output=$(@D)/output.go
        
        # Step 3: Generate tests
        $(location //tools:test_generator) \\
            --input=$(location input.json) \\
            --output=$(@D)/output_test.go
    """,
    tools = [
        "//tools:json_validator",
        "//tools:go_generator", 
        "//tools:test_generator",
    ],
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

**Detailed examples:**
```python
go_library(
    name = "user_service",
    srcs = ["service.go"],
    deps = [
        # Internal dependencies (other targets in this workspace)
        ":user_api_go_proto",             # Generated proto code
        "//pkg/database:db_lib",          # Database utilities
        "//pkg/auth:auth_lib",           # Authentication
        
        # External dependencies (third-party libraries)
        "@com_github_gorilla_mux//:mux",     # HTTP router
        "@org_golang_google_grpc//:grpc",    # gRPC framework
        "@com_github_sirupsen_logrus//:logrus", # Logging
        
        # Standard library (usually implicit, but can be explicit)
        "@io_bazel_rules_go//proto/wkt:duration_go_proto",
    ],
)
```

**Dependency management best practices:**
```python
# ✅ Good: minimal, specific dependencies
go_library(
    name = "math_utils",
    deps = [
        "@org_golang_x_exp//constraints", # Only what you actually use
    ],
)

# ❌ Bad: unnecessary dependencies
go_library(
    name = "math_utils", 
    deps = [
        "@com_github_gorilla_mux//:mux",   # Why does math need HTTP routing?
        "//services/user:user_lib",        # Creates circular dependency risk
    ],
)

# ✅ Good: well-organized dependency layers
go_library(
    name = "api_layer",
    deps = [
        ":service_layer",                  # API depends on service
        "@com_github_gorilla_mux//:mux",
    ],
)

go_library(
    name = "service_layer", 
    deps = [
        ":data_layer",                     # Service depends on data
        ":business_logic",
    ],
)

go_library(
    name = "data_layer",
    deps = [
        "@com_github_lib_pq//:pq",        # Data layer depends on database driver
    ],
)
```

## Complete Working Examples

Let's build three increasingly complex projects to demonstrate Bazel concepts in practice. Each example builds on the previous one, showing how real-world projects evolve.

### Simple Go Project with Bazel

This example shows the minimal setup for a Go project with Bazel.

**Project structure:**
```
simple-go-project/
├── WORKSPACE
├── BUILD.bazel
├── main.go
└── math/
    ├── BUILD.bazel
    ├── calculator.go
    └── calculator_test.go
```

**WORKSPACE file:**
```python
workspace(name = "simple_go_project")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Download rules_go
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "6b65cb7917b4d1709f9410ffe00ecf3e160edf674b78c54a894471320862184f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.39.1/rules_go-v0.39.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.39.1/rules_go-v0.39.1.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()
go_register_toolchains(version = "1.19.5")
```

**Root BUILD.bazel:**
```python
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

# Main application library
go_library(
    name = "main_lib",
    srcs = ["main.go"],
    importpath = "github.com/example/simple-go-project",
    deps = ["//math:calculator_lib"],
)

# Executable binary
go_binary(
    name = "calculator_app",
    embed = [":main_lib"],
)
```

**main.go:**
```go
package main

import (
    "fmt"
    "os"
    "strconv"
    
    "github.com/example/simple-go-project/math"
)

func main() {
    if len(os.Args) != 4 {
        fmt.Println("Usage: calculator <num1> <operation> <num2>")
        fmt.Println("Operations: +, -, *, /")
        os.Exit(1)
    }
    
    num1, _ := strconv.ParseFloat(os.Args[1], 64)
    operation := os.Args[2]
    num2, _ := strconv.ParseFloat(os.Args[3], 64)
    
    calc := math.NewCalculator()
    var result float64
    var err error
    
    switch operation {
    case "+":
        result = calc.Add(num1, num2)
    case "-":
        result = calc.Subtract(num1, num2)
    case "*":
        result = calc.Multiply(num1, num2)
    case "/":
        result, err = calc.Divide(num1, num2)
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            os.Exit(1)
        }
    default:
        fmt.Printf("Unknown operation: %s\n", operation)
        os.Exit(1)
    }
    
    fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operation, num2, result)
}
```

**math/BUILD.bazel:**
```python
load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "calculator_lib",
    srcs = ["calculator.go"],
    importpath = "github.com/example/simple-go-project/math",
    visibility = ["//visibility:public"],  # Allow main package to use this
)

go_test(
    name = "calculator_test",
    srcs = ["calculator_test.go"],
    embed = [":calculator_lib"],
)
```

**math/calculator.go:**
```go
package math

import "errors"

type Calculator struct{}

func NewCalculator() *Calculator {
    return &Calculator{}
}

func (c *Calculator) Add(a, b float64) float64 {
    return a + b
}

func (c *Calculator) Subtract(a, b float64) float64 {
    return a - b
}

func (c *Calculator) Multiply(a, b float64) float64 {
    return a * b
}

func (c *Calculator) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

**math/calculator_test.go:**
```go
package math

import (
    "testing"
)

func TestCalculator_Add(t *testing.T) {
    calc := NewCalculator()
    result := calc.Add(2, 3)
    if result != 5 {
        t.Errorf("Expected 5, got %f", result)
    }
}

func TestCalculator_Divide(t *testing.T) {
    calc := NewCalculator()
    
    // Test normal division
    result, err := calc.Divide(10, 2)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if result != 5 {
        t.Errorf("Expected 5, got %f", result)
    }
    
    // Test division by zero
    _, err = calc.Divide(10, 0)
    if err == nil {
        t.Error("Expected error for division by zero")
    }
}
```

**Build and run:**
```bash
# Build the application
bazel build //:calculator_app

# Run the application
bazel run //:calculator_app -- 10 + 5

# Run tests
bazel test //math:calculator_test

# Run all tests
bazel test //...
```

### Go + Protocol Buffers Project

This example adds Protocol Buffers to create a more realistic service-oriented application.

**Project structure:**
```
grpc-calculator/
├── WORKSPACE
├── BUILD.bazel
├── main.go
├── api/
│   ├── BUILD.bazel
│   └── calculator.proto
├── server/
│   ├── BUILD.bazel
│   ├── server.go
│   └── server_test.go
└── client/
    ├── BUILD.bazel
    └── client.go
```

**Enhanced WORKSPACE:**
```python
workspace(name = "grpc_calculator")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Rules Go
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "6b65cb7917b4d1709f9410ffe00ecf3e160edf674b78c54a894471320862184f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.39.1/rules_go-v0.39.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.39.1/rules_go-v0.39.1.zip",
    ],
)

# Rules Proto
http_archive(
    name = "rules_proto",
    sha256 = "dc3fb206a2cb3441b485eb1e423165b231235a1ea9b031b4433cf7bc1fa460dd",
    strip_prefix = "rules_proto-5.3.0-21.7",
    urls = [
        "https://github.com/bazelbuild/rules_proto/archive/refs/tags/5.3.0-21.7.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

go_rules_dependencies()
go_register_toolchains(version = "1.19.5")

rules_proto_dependencies()
rules_proto_toolchains()

# gRPC dependencies
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
```

**api/calculator.proto:**
```protobuf
syntax = "proto3";

package calculator.v1;

option go_package = "github.com/example/grpc-calculator/api/v1;calculatorv1";

// Calculator service provides basic arithmetic operations
service CalculatorService {
  // Add two numbers
  rpc Add(AddRequest) returns (AddResponse);
  
  // Subtract two numbers  
  rpc Subtract(SubtractRequest) returns (SubtractResponse);
  
  // Multiply two numbers
  rpc Multiply(MultiplyRequest) returns (MultiplyResponse);
  
  // Divide two numbers
  rpc Divide(DivideRequest) returns (DivideResponse);
}

message AddRequest {
  double a = 1;
  double b = 2;
}

message AddResponse {
  double result = 1;
}

message SubtractRequest {
  double a = 1;
  double b = 2;
}

message SubtractResponse {
  double result = 1;
}

message MultiplyRequest {
  double a = 1;
  double b = 2;
}

message MultiplyResponse {
  double result = 1;
}

message DivideRequest {
  double a = 1;
  double b = 2;
}

message DivideResponse {
  double result = 1;
  string error = 2;  // Error message if division fails
}
```

**api/BUILD.bazel:**
```python
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

# Protocol buffer definitions
proto_library(
    name = "calculator_proto",
    srcs = ["calculator.proto"],
    visibility = ["//visibility:public"],
)

# Generated Go code from protobuf
go_proto_library(
    name = "calculator_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],  # Include gRPC services
    importpath = "github.com/example/grpc-calculator/api/v1",
    proto = ":calculator_proto",
    visibility = ["//visibility:public"],
)
```

**server/server.go:**
```go
package server

import (
    "context"
    "errors"
    
    calculatorv1 "github.com/example/grpc-calculator/api/v1"
)

type CalculatorServer struct {
    calculatorv1.UnimplementedCalculatorServiceServer
}

func NewCalculatorServer() *CalculatorServer {
    return &CalculatorServer{}
}

func (s *CalculatorServer) Add(ctx context.Context, req *calculatorv1.AddRequest) (*calculatorv1.AddResponse, error) {
    result := req.A + req.B
    return &calculatorv1.AddResponse{Result: result}, nil
}

func (s *CalculatorServer) Subtract(ctx context.Context, req *calculatorv1.SubtractRequest) (*calculatorv1.SubtractResponse, error) {
    result := req.A - req.B
    return &calculatorv1.SubtractResponse{Result: result}, nil
}

func (s *CalculatorServer) Multiply(ctx context.Context, req *calculatorv1.MultiplyRequest) (*calculatorv1.MultiplyResponse, error) {
    result := req.A * req.B
    return &calculatorv1.MultiplyResponse{Result: result}, nil
}

func (s *CalculatorServer) Divide(ctx context.Context, req *calculatorv1.DivideRequest) (*calculatorv1.DivideResponse, error) {
    if req.B == 0 {
        return &calculatorv1.DivideResponse{
            Result: 0,
            Error:  "division by zero",
        }, nil
    }
    
    result := req.A / req.B
    return &calculatorv1.DivideResponse{Result: result}, nil
}
```

**Build commands:**
```bash
# Build everything
bazel build //...

# Test everything  
bazel test //...

# Run the server
bazel run //server:calculator_server

# Run the client
bazel run //client:calculator_client
```

### Microservice with gRPC and Swagger

This final example shows a production-ready microservice with API documentation, health checks, and proper service architecture.

**Key additions:**
- Health check endpoints
- Swagger/OpenAPI documentation generation
- Structured logging
- Configuration management
- Docker deployment support

**Extended project structure:**
```
production-service/
├── WORKSPACE
├── BUILD.bazel
├── api/
│   ├── BUILD.bazel
│   ├── user.proto
│   └── health.proto
├── cmd/
│   ├── server/
│   │   ├── BUILD.bazel
│   │   └── main.go
│   └── client/
│       ├── BUILD.bazel
│       └── main.go
├── internal/
│   ├── server/
│   │   ├── BUILD.bazel
│   │   ├── user_service.go
│   │   └── health_service.go
│   ├── config/
│   │   ├── BUILD.bazel
│   │   └── config.go
│   └── repository/
│       ├── BUILD.bazel
│       ├── user_repo.go
│       └── memory_store.go
├── docs/
│   └── BUILD.bazel  # Generated API docs
└── deployments/
    ├── BUILD.bazel
    └── Dockerfile
```

This production example would include:
- Proper error handling with status codes
- Request/response logging middleware
- Configuration from environment variables
- Database abstractions
- API documentation generation
- Health check services
- Docker containerization
- Kubernetes deployment manifests

The complexity scales naturally with Bazel's modular approach, where each component has its own BUILD file and clear dependencies.

## Automating with Gazelle

Gazelle is a powerful tool that automatically generates and maintains BUILD files for Go projects. Think of it as an intelligent assistant that watches your Go code and creates the appropriate Bazel configuration.

### What Gazelle Does

**The Problem:** Manually writing BUILD files is tedious and error-prone:
- You have to remember to add every new Go file to `srcs`
- You need to manually track and add dependencies
- Import paths must be kept in sync with your code
- BUILD files can become out of date as code evolves

**Gazelle's Solution:** Automatically generates BUILD files by:
- Scanning your Go source files
- Analyzing import statements
- Determining dependencies between packages
- Creating appropriate `go_library`, `go_binary`, and `go_test` targets
- Keeping everything in sync as your code changes

### Setting Up Gazelle

**1. Add Gazelle to your WORKSPACE:**
```python
# WORKSPACE
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Rules Go (required for Gazelle)
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "6b65cb7917b4d1709f9410ffe00ecf3e160edf674b78c54a894471320862184f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.39.1/rules_go-v0.39.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.39.1/rules_go-v0.39.1.zip",
    ],
)

# Gazelle
http_archive(
    name = "bazel_gazelle",
    sha256 = "ecba0f04f96b4960a5b250c8e8eeec42281035970aa8852dda73098613154a93",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

go_rules_dependencies()
go_register_toolchains(version = "1.19.5")
gazelle_dependencies()
```

**2. Create the root BUILD.bazel file:**
```python
# BUILD.bazel (at workspace root)
load("@bazel_gazelle//:def.bzl", "gazelle")

# Gazelle configuration
gazelle(
    name = "gazelle",
    prefix = "github.com/mycompany/myproject",  # Your module's import prefix
)

# Update BUILD files and go.mod
gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)
```

**3. Create go.mod file (if you don't have one):**
```go
// go.mod
module github.com/mycompany/myproject

go 1.19

require (
    github.com/gorilla/mux v1.8.0
    github.com/sirupsen/logrus v1.9.0
    google.golang.org/grpc v1.50.0
    google.golang.org/protobuf v1.28.1
)
```

### Using Gazelle

**Basic workflow:**

1. **Write your Go code** as you normally would
2. **Run Gazelle** to generate/update BUILD files
3. **Build and test** with Bazel

**Commands:**
```bash
# Generate/update BUILD files for your Go code
bazel run //:gazelle

# Update external dependencies from go.mod
bazel run //:gazelle-update-repos

# Do both in one command
bazel run //:gazelle && bazel run //:gazelle-update-repos
```

### Real Example: Gazelle in Action

**Before Gazelle - Manual BUILD files:**

Let's say you have this Go code structure:
```
myproject/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── api/
│   │   └── handler.go
│   └── database/
│       └── connection.go
└── pkg/
    └── utils/
        └── logger.go
```

**cmd/server/main.go:**
```go
package main

import (
    "log"
    "net/http"
    
    "github.com/gorilla/mux"
    "github.com/mycompany/myproject/internal/api"
    "github.com/mycompany/myproject/pkg/utils"
)

func main() {
    logger := utils.NewLogger()
    handler := api.NewHandler(logger)
    
    r := mux.NewRouter()
    r.PathPrefix("/api/").Handler(handler)
    
    log.Fatal(http.ListenAndServe(":8080", r))
}
```

**internal/api/handler.go:**
```go
package api

import (
    "net/http"
    
    "github.com/mycompany/myproject/internal/database"
    "github.com/mycompany/myproject/pkg/utils"
)

type Handler struct {
    db     *database.Connection
    logger *utils.Logger
}

func NewHandler(logger *utils.Logger) *Handler {
    return &Handler{
        db:     database.NewConnection(),
        logger: logger,
    }
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    h.logger.Info("Handling request")
    w.WriteHeader(http.StatusOK)
}
```

**Without Gazelle, you'd need to manually create:**

**cmd/server/BUILD.bazel:**
```python
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "server_lib",
    srcs = ["main.go"],
    importpath = "github.com/mycompany/myproject/cmd/server",
    deps = [
        "//internal/api:api_lib",
        "//pkg/utils:utils_lib",
        "@com_github_gorilla_mux//:mux",
    ],
)

go_binary(
    name = "server",
    embed = [":server_lib"],
)
```

**internal/api/BUILD.bazel:**
```python
load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "api_lib",
    srcs = ["handler.go"],
    importpath = "github.com/mycompany/myproject/internal/api",
    deps = [
        "//internal/database:database_lib",
        "//pkg/utils:utils_lib",
    ],
)
```

**With Gazelle - Automatic generation:**

1. **Run Gazelle:**
```bash
bazel run //:gazelle
```

2. **Gazelle automatically creates all BUILD files** by analyzing your import statements and dependencies!

3. **Generated files are identical** to the manual ones above, but created instantly and kept in sync automatically.

### Advanced Gazelle Configuration

**Gazelle directives** - Add comments to your Go files to control Gazelle behavior:

```go
// Package api provides HTTP handlers
//
// gazelle:ignore
// gazelle:build_file_name BUILD.bazel
// gazelle:prefix github.com/mycompany/myproject
package api
```

**Common directives:**
- `gazelle:ignore` - Skip this file/directory
- `gazelle:build_file_name BUILD.bazel` - Use BUILD.bazel instead of BUILD
- `gazelle:prefix github.com/mycompany/myproject` - Set import path prefix
- `gazelle:resolve go github.com/external/pkg //path/to:target` - Custom dependency resolution

**BUILD file directives:**
```python
# gazelle:prefix github.com/mycompany/myproject
# gazelle:resolve go github.com/special/package //custom:target

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "api_lib",
    srcs = ["handler.go"],
    importpath = "github.com/mycompany/myproject/internal/api",
    deps = [
        "//custom:target",  # Custom resolution
        "//internal/database:database_lib",
    ],
)
```

### Gazelle Best Practices

1. **Run Gazelle regularly:**
   ```bash
   # After adding new files
   bazel run //:gazelle
   
   # After changing dependencies
   bazel run //:gazelle-update-repos
   ```

2. **Use go.mod for dependency management:**
   - Keep go.mod updated with your dependencies
   - Let Gazelle sync Bazel with go.mod
   - Don't manually edit generated dependency files

3. **Commit generated BUILD files:**
   - Include BUILD files in version control
   - Makes builds reproducible for team members
   - Allows builds without requiring Gazelle

4. **Use gazelle:ignore sparingly:**
   - Only ignore files that truly shouldn't be built
   - Consider if the file structure could be improved instead

5. **Set up automation:**
   ```bash
   # Add to your CI/CD pipeline
   bazel run //:gazelle
   git diff --exit-code  # Fail if BUILD files are out of sync
   ```

### Troubleshooting Gazelle

**Common issues and solutions:**

1. **"Import path not found"**
   ```bash
   # Solution: Update external dependencies
   bazel run //:gazelle-update-repos
   ```

2. **Wrong import paths in generated code**
   ```bash
   # Solution: Check your gazelle prefix configuration
   # gazelle:prefix github.com/your-actual-module-name
   ```

3. **Missing dependencies**
   ```bash
   # Solution: Ensure go.mod is up to date
   go mod tidy
   bazel run //:gazelle-update-repos
   ```

4. **Gazelle ignoring certain files**
   ```bash
   # Check for gazelle:ignore comments
   # Check file naming (must end in .go)
   # Check if directory contains valid Go package
   ```

## Managing Dependencies Like a Pro

Dependency management is one of Bazel's strongest features, but it requires understanding several concepts to use effectively. This section shows you how to handle dependencies in real-world scenarios.

### Understanding Bazel Dependency Types

**1. Internal Dependencies (Workspace targets)**
These are targets within your own workspace:
```python
go_library(
    name = "user_service",
    deps = [
        "//internal/database:db_lib",     # Internal dependency
        "//pkg/auth:auth_lib",            # Internal dependency
        ":user_proto_go",                 # Same package dependency
    ],
)
```

**2. External Dependencies (Third-party packages)**
These are packages from external repositories:
```python
go_library(
    name = "api_server",
    deps = [
        "@com_github_gorilla_mux//:mux",           # HTTP router
        "@org_golang_google_grpc//:grpc",          # gRPC framework
        "@com_github_sirupsen_logrus//:logrus",    # Logging
    ],
)
```

**3. Generated Dependencies (From proto, etc.)**
These are created by other Bazel rules:
```python
go_library(
    name = "service_impl",
    deps = [
        ":user_api_go_proto",            # Generated from proto_library
        ":swagger_generated_models",      # Generated from OpenAPI spec
    ],
)
```

### Adding External Dependencies

**Method 1: Using go.mod + Gazelle (Recommended)**

This is the easiest and most maintainable approach:

1. **Add dependency to go.mod:**
   ```go
   // go.mod
   module github.com/mycompany/myproject
   
   go 1.19
   
   require (
       github.com/gorilla/mux v1.8.0
       github.com/lib/pq v1.10.7        // ← New dependency
       github.com/stretchr/testify v1.8.1
   )
   ```

2. **Update Bazel dependencies:**
   ```bash
   bazel run //:gazelle-update-repos
   ```

3. **Use in your BUILD files:**
   ```python
   go_library(
       name = "database_lib",
       srcs = ["db.go"],
       deps = [
           "@com_github_lib_pq//:pq",    # ← Now available
       ],
   )
   ```

**Method 2: Manual go_repository (For fine control)**

When you need precise control over versions or patches:

```python
# WORKSPACE
load("@bazel_gazelle//:deps.bzl", "go_repository")

go_repository(
    name = "com_github_lib_pq",
    importpath = "github.com/lib/pq",
    version = "v1.10.7",
    # Optional: apply patches
    patches = ["//patches:pq-custom.patch"],
    patch_args = ["-p1"],
)

go_repository(
    name = "com_github_gorilla_mux",
    importpath = "github.com/gorilla/mux", 
    version = "v1.8.0",
    # Optional: use specific commit
    commit = "00bdffe0f3c77e27d2cf6f5c70232a2d3e4d9c15",
)

go_repository(
    name = "org_golang_google_grpc",
    build_file_proto_mode = "disable",    # Don't auto-generate proto rules
    importpath = "google.golang.org/grpc",
    version = "v1.50.0",
)
```

### Advanced Dependency Management

**1. Version Pinning and Updates**

Create a dedicated file for dependency management:

```python
# deps.bzl - Generated by Gazelle
load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_gorilla_mux",
        importpath = "github.com/gorilla/mux",
        version = "v1.8.0",
    )
    
    go_repository(
        name = "com_github_lib_pq", 
        importpath = "github.com/lib/pq",
        version = "v1.10.7",
    )
    
    # ... more dependencies
```

Then in WORKSPACE:
```python
# WORKSPACE
load("//:deps.bzl", "go_dependencies")
go_dependencies()
```

**2. Handling Conflicting Dependencies**

When different parts of your project need different versions:

```python
# WORKSPACE
go_repository(
    name = "com_github_lib_pq_v1_10",
    importpath = "github.com/lib/pq",
    version = "v1.10.7",
)

go_repository(
    name = "com_github_lib_pq_v1_9", 
    importpath = "github.com/lib/pq",
    version = "v1.9.0",
)

# In BUILD files, use specific versions
go_library(
    name = "old_service",
    deps = ["@com_github_lib_pq_v1_9//:pq"],  # Old version
)

go_library(
    name = "new_service", 
    deps = ["@com_github_lib_pq_v1_10//:pq"], # New version
)
```

**3. Private Dependencies**

For internal or private repositories:

```python
# WORKSPACE
go_repository(
    name = "com_mycompany_internal_lib",
    importpath = "github.com/mycompany/internal-lib",
    vcs = "git",
    remote = "https://github.com/mycompany/internal-lib.git",
    commit = "abc123def456",  # Specific commit for reproducibility
    # For private repos, you might need authentication
)

# For Git repositories with authentication
go_repository(
    name = "com_mycompany_private",
    importpath = "github.com/mycompany/private-repo",
    remote = "git@github.com:mycompany/private-repo.git",
    vcs = "git", 
    commit = "latest",
)
```

**4. Dependency Overrides**

Sometimes you need to patch or replace dependencies:

```python
# WORKSPACE  
go_repository(
    name = "com_github_broken_lib",
    importpath = "github.com/broken/lib",
    version = "v1.0.0",
    patches = [
        "//patches:fix-critical-bug.patch",
        "//patches:add-missing-feature.patch", 
    ],
    patch_args = ["-p1"],
)

# Use local replacement during development
go_repository(
    name = "com_github_external_dep",
    importpath = "github.com/external/dep",
    path = "/path/to/local/checkout",  # Local development version
)
```

### Dependency Best Practices

**1. Keep dependencies minimal:**
```python
# ✅ Good: only include what you actually use
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

## How to Run Each Bazel Rule and Command

This section provides detailed instructions on how to build, run, test, and work with each type of Bazel rule and command. Each subsection includes practical examples and common use cases.

### Running proto_library Targets

**What proto_library does when executed:**
- Validates Protocol Buffer syntax
- Generates language-agnostic protobuf descriptors
- Makes proto definitions available to other rules

**How to work with proto_library:**

```bash
# 1. Validate proto files (check for syntax errors)
bazel build //api:user_proto
# This compiles the proto files and checks for errors

# 2. Query proto dependencies
bazel query "deps(//api:user_proto)"
# Shows all dependencies of the proto target

# 3. See what files are generated
bazel build //api:user_proto --output_groups=default
ls -la bazel-bin/api/

# 4. Inspect proto descriptor
bazel build //api:user_proto
# Generated descriptor will be in bazel-bin/api/user_proto-descriptor-set.proto.bin
```

**Example with real commands:**
```bash
# Given this BUILD file:
# proto_library(
#     name = "user_api_proto",
#     srcs = ["user.proto", "auth.proto"],
# )

# Build and validate the proto files
bazel build //api:user_api_proto
# Output: Target //api:user_api_proto up-to-date:
#   bazel-bin/api/user_api_proto-descriptor-set.proto.bin

# Check what proto files are included
bazel query --output=build //api:user_api_proto
```

**Common proto_library commands:**
```bash
# Build specific proto target
bazel build //path/to:proto_target

# Build all proto targets in a package
bazel build //api:all

# Validate all proto files in workspace
bazel build //... --build_tag_filters=proto

# Show proto file dependencies
bazel query "kind(proto_library, //...)"
```

### Running go_proto_library Targets

**What go_proto_library does when executed:**
- Generates Go structs from proto messages
- Creates Go interfaces for gRPC services
- Produces importable Go packages

**How to work with go_proto_library:**

```bash
# 1. Generate Go code from proto
bazel build //api:user_api_go_proto
# Generates .go files with structs and gRPC interfaces

# 2. See generated Go files
bazel build //api:user_api_go_proto
ls -la bazel-bin/api/user_api_go_proto_/github.com/mycompany/myproject/api/user/v1/

# 3. Use in Go code (after building)
# The generated code can be imported in your Go files
```

**Example with real commands:**
```bash
# Given this BUILD file:
# go_proto_library(
#     name = "user_api_go_proto",
#     proto = ":user_api_proto",
#     importpath = "github.com/mycompany/myproject/api/user/v1",
# )

# Generate Go code from proto
bazel build //api:user_api_go_proto
# Output: Target //api:user_api_go_proto up-to-date:
#   bazel-bin/api/libuser_api_go_proto.a

# Inspect generated Go code structure
find bazel-bin/api/user_api_go_proto_/ -name "*.go" -exec head -10 {} \;

# Build projects that depend on this generated code
bazel build //services/user:user_service_lib
```

**Common go_proto_library commands:**
```bash
# Generate Go code from specific proto
bazel build //api:specific_go_proto

# Generate all Go proto libraries
bazel build //... --build_tag_filters=go_proto

# Check what Go packages are generated
bazel query "kind(go_proto_library, //...)"

# Verify generated code compiles
bazel test //... --test_tag_filters=go_proto_test
```

### Running go_library Targets

**What go_library does when executed:**
- Compiles Go source files into a package
- Creates reusable library archive (.a file)
- Makes the package available for other targets to import

**How to work with go_library:**

```bash
# 1. Build a Go library
bazel build //pkg/database:db_lib
# Compiles the library and creates .a archive file

# 2. Check library compilation
bazel build //pkg/database:db_lib --verbose_failures
# Shows detailed error messages if compilation fails

# 3. See what's included in the library
bazel query --output=build //pkg/database:db_lib

# 4. Test library compilation across workspace
bazel build //... --build_tag_filters=go_library
```

**Example with real commands:**
```bash
# Given this BUILD file:
# go_library(
#     name = "user_service_lib",
#     srcs = ["service.go", "handler.go"],
#     importpath = "github.com/mycompany/myproject/services/user",
#     deps = ["//pkg/database:db_lib"],
# )

# Build the library
bazel build //services/user:user_service_lib
# Output: Target //services/user:user_service_lib up-to-date:
#   bazel-bin/services/user/libuser_service_lib.a

# Verify the library's dependencies are satisfied
bazel build //services/user:user_service_lib --check_up_to_date

# See detailed build information
bazel build //services/user:user_service_lib -s
# Shows all commands executed during build
```

**Common go_library commands:**
```bash
# Build specific library
bazel build //path/to:library_name

# Build all libraries in a package
bazel build //services/user:all

# Build all Go libraries in workspace
bazel build //... --build_tag_filters=go_library

# Check library dependencies
bazel query "deps(//services/user:user_service_lib)"

# Find libraries that depend on a specific library
bazel query "rdeps(//..., //pkg/database:db_lib)"
```

### Running go_binary Targets

**What go_binary does when executed:**
- Links Go libraries into an executable
- Creates a runnable program
- Handles platform-specific compilation

**How to work with go_binary:**

```bash
# 1. Build an executable
bazel build //cmd/server:server
# Creates executable binary

# 2. Run the executable
bazel run //cmd/server:server
# Builds (if needed) and runs the program

# 3. Run with arguments
bazel run //cmd/server:server -- --port=8080 --config=dev.yaml

# 4. Build for different platforms
bazel build //cmd/server:server --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64
```

**Example with real commands:**
```bash
# Given this BUILD file:
# go_binary(
#     name = "user_server",
#     embed = [":server_lib"],
# )

# Build the binary
bazel build //cmd/server:user_server
# Output: Target //cmd/server:user_server up-to-date:
#   bazel-bin/cmd/server/user_server_/user_server

# Run the server
bazel run //cmd/server:user_server
# Builds and executes the server

# Run with command line arguments
bazel run //cmd/server:user_server -- --help
bazel run //cmd/server:user_server -- --port=9090 --debug=true

# Build for production deployment
bazel build //cmd/server:user_server --compilation_mode=opt
```

**Common go_binary commands:**
```bash
# Build executable
bazel build //cmd/app:app_name

# Build and run immediately
bazel run //cmd/app:app_name

# Run with arguments
bazel run //cmd/app:app_name -- arg1 arg2 --flag=value

# Build optimized version
bazel build //cmd/app:app_name --compilation_mode=opt

# Build for specific platform
bazel build //cmd/app:app_name --platforms=@io_bazel_rules_go//go/toolchain:darwin_amd64

# See executable size and info
bazel build //cmd/app:app_name && ls -lh bazel-bin/cmd/app/app_name_/app_name
```

### Running go_test Targets

**What go_test does when executed:**
- Compiles test files with the code under test
- Runs the tests and reports results
- Provides test isolation and caching

**How to work with go_test:**

```bash
# 1. Run tests for a specific target
bazel test //pkg/database:db_test
# Compiles and runs tests, shows pass/fail results

# 2. Run tests with verbose output
bazel test //pkg/database:db_test --test_output=all
# Shows detailed test output and logs

# 3. Run specific test functions
bazel test //pkg/database:db_test --test_filter="TestUserRepository.*"

# 4. Run tests in different modes
bazel test //pkg/database:db_test --test_timeout=300
```

**Example with real commands:**
```bash
# Given this BUILD file:
# go_test(
#     name = "user_service_test",
#     srcs = ["service_test.go"],
#     embed = [":user_service_lib"],
# )

# Run the tests
bazel test //services/user:user_service_test
# Output: //services/user:user_service_test PASSED in 0.8s

# Run tests with detailed output
bazel test //services/user:user_service_test --test_output=all
# Shows all test logs and printf statements

# Run specific test cases
bazel test //services/user:user_service_test --test_filter="TestCreateUser"

# Run tests with coverage
bazel coverage //services/user:user_service_test
# Generates coverage report
```

**Common go_test commands:**
```bash
# Run specific test target
bazel test //path/to:test_target

# Run all tests in a package
bazel test //services/user:all

# Run all tests in workspace
bazel test //...

# Run tests with specific tags
bazel test //... --test_tag_filters=unit
bazel test //... --test_tag_filters=integration

# Run tests excluding certain tags
bazel test //... --test_tag_filters=-flaky

# Run tests with verbose output
bazel test //path/to:test --test_output=all

# Run tests with coverage
bazel coverage //path/to:test

# Run tests with specific timeout
bazel test //path/to:test --test_timeout=60

# Debug failing tests
bazel test //path/to:test --test_output=errors --verbose_failures
```

### Running protoc_gen_swagger Targets

**What protoc_gen_swagger does when executed:**
- Parses proto service definitions
- Generates OpenAPI/Swagger documentation
- Creates JSON/YAML specification files

**How to work with protoc_gen_swagger:**

```bash
# 1. Generate Swagger documentation
bazel build //api:user_api_swagger
# Creates swagger.json or openapi.yaml files

# 2. View generated documentation
bazel build //api:user_api_swagger
cat bazel-bin/api/swagger.json | jq .

# 3. Serve documentation (if you have a server target)
bazel run //docs:swagger_server
```

**Example with real commands:**
```bash
# Given this BUILD file:
# protoc_gen_openapiv2(
#     name = "user_api_openapi",
#     proto = ":user_api_proto",
# )

# Generate the OpenAPI specification
bazel build //api:user_api_openapi
# Output: Target //api:user_api_openapi up-to-date:
#   bazel-bin/api/user_api_openapi.swagger.json

# View the generated documentation
cat bazel-bin/api/user_api_openapi.swagger.json | jq .

# Copy to a web server directory
cp bazel-bin/api/user_api_openapi.swagger.json /var/www/html/api-docs/
```

**Common swagger generation commands:**
```bash
# Generate Swagger docs
bazel build //api:swagger_target

# Generate all API documentation
bazel build //... --build_tag_filters=swagger

# Validate generated Swagger
swagger-codegen validate bazel-bin/api/swagger.json

# Generate client code from Swagger
swagger-codegen generate -i bazel-bin/api/swagger.json -l go -o ./client
```

### Running genrule Targets

**What genrule does when executed:**
- Runs custom shell commands
- Processes input files to create output files
- Executes during the build process

**How to work with genrule:**

```bash
# 1. Execute a genrule
bazel build //tools:generate_config
# Runs the shell command and creates output files

# 2. See genrule command details
bazel build //tools:generate_config -s
# Shows the exact command that was executed

# 3. Force re-execution of genrule
bazel build //tools:generate_config --force_recompile
```

**Example with real commands:**
```bash
# Given this BUILD file:
# genrule(
#     name = "generate_version",
#     outs = ["version.go"],
#     cmd = "echo 'package main; const Version = \"1.0.0\"' > $@",
# )

# Execute the genrule
bazel build //tools:generate_version
# Output: Target //tools:generate_version up-to-date:
#   bazel-bin/tools/version.go

# View the generated file
cat bazel-bin/tools/version.go
# Output: package main; const Version = "1.0.0"

# See the exact command that was run
bazel build //tools:generate_version -s --verbose_failures
```

**Common genrule commands:**
```bash
# Execute specific genrule
bazel build //path/to:genrule_target

# Execute all genrules
bazel build //... --build_tag_filters=genrule

# Force re-execution (ignore cache)
bazel build //path/to:genrule_target --spawn_strategy=standalone

# Debug genrule execution
bazel build //path/to:genrule_target -s --sandbox_debug

# See genrule outputs
bazel build //path/to:genrule_target && ls -la bazel-bin/path/to/
```

### Running Gazelle Commands

**What Gazelle does when executed:**
- Scans Go source files
- Generates/updates BUILD files
- Manages external dependencies

**How to work with Gazelle:**

```bash
# 1. Generate BUILD files
bazel run //:gazelle
# Scans Go code and creates/updates BUILD files

# 2. Update dependencies from go.mod
bazel run //:gazelle-update-repos
# Syncs external dependencies with go.mod

# 3. Fix import paths
bazel run //:gazelle -- fix
```

**Example with real commands:**
```bash
# Setup: Create these targets in root BUILD.bazel:
# gazelle(name = "gazelle")
# gazelle(name = "gazelle-update-repos", command = "update-repos")

# Generate BUILD files for all Go code
bazel run //:gazelle
# Output: Gazelle generated/updated BUILD files

# Update external dependencies
bazel run //:gazelle-update-repos
# Reads go.mod and updates deps.bzl

# Fix specific package
bazel run //:gazelle -- fix //pkg/database

# Generate with specific prefix
bazel run //:gazelle -- -go_prefix=github.com/mycompany/myproject
```

**Common Gazelle commands:**
```bash
# Basic BUILD file generation
bazel run //:gazelle

# Update external dependencies
bazel run //:gazelle-update-repos

# Fix import paths in specific directory
bazel run //:gazelle -- fix //path/to/package

# Generate with custom configuration
bazel run //:gazelle -- -build_file_name=BUILD.bazel

# Update and prune unused dependencies
bazel run //:gazelle-update-repos -- -prune

# Dry run (see what would change)
bazel run //:gazelle -- -mode=diff
```

### General Bazel Commands for All Targets

**Build Commands:**
```bash
# Build specific target
bazel build //path/to:target

# Build all targets in package
bazel build //path/to:all

# Build everything in workspace
bazel build //...

# Build with optimizations
bazel build //path/to:target --compilation_mode=opt

# Build for specific platform
bazel build //path/to:target --platforms=//tools:linux_x86_64
```

**Query Commands:**
```bash
# Show target dependencies
bazel query "deps(//path/to:target)"

# Show reverse dependencies
bazel query "rdeps(//..., //path/to:target)"

# Find targets by type
bazel query "kind(go_library, //...)"

# Show target details
bazel query --output=build //path/to:target
```

**Clean Commands:**
```bash
# Clean build outputs
bazel clean

# Clean everything including external dependencies
bazel clean --expunge

# Clean specific target
bazel clean //path/to:target
```

**Info Commands:**
```bash
# Show workspace info
bazel info

# Show build configuration
bazel info build-config

# Show output directories
bazel info bazel-bin bazel-genfiles
```

This comprehensive guide provides the exact commands and workflows for running each type of Bazel rule and target in your workspace.

## Understanding Workspace Structure

The workspace is the foundation of every Bazel project. Understanding how to structure and configure it properly is crucial for successful Bazel adoption.

### WORKSPACE File Deep Dive

The `WORKSPACE` file is the entry point for your Bazel project. It serves multiple purposes:

1. **Defines the workspace boundary** - Everything under this directory is part of your project
2. **Declares external dependencies** - Third-party libraries and tools
3. **Configures build rules** - Sets up language-specific build systems
4. **Establishes workspace name** - Used for external references

**Complete WORKSPACE example:**
```python
# WORKSPACE - Required at the root of every Bazel project

# Workspace name - used when this project is a dependency of others
workspace(name = "my_awesome_project")

# Load utilities for downloading external dependencies
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

###################
# Go Language Support
###################

# Download rules_go - Provides Go language support for Bazel
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "6b65cb7917b4d1709f9410ffe00ecf3e160edf674b78c54a894471320862184f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.39.1/rules_go-v0.39.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.39.1/rules_go-v0.39.1.zip",
    ],
)

# Download Gazelle - Automated BUILD file generation for Go
http_archive(
    name = "bazel_gazelle",
    sha256 = "ecba0f04f96b4960a5b250c8e8eeec42281035970aa8852dda73098613154a93",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
    ],
)

###################
# Protocol Buffer Support
###################

# Download rules_proto - Protocol buffer support
http_archive(
    name = "rules_proto",
    sha256 = "dc3fb206a2cb3441b485eb1e423165b231235a1ea9b031b4433cf7bc1fa460dd",
    strip_prefix = "rules_proto-5.3.0-21.7",
    urls = [
        "https://github.com/bazelbuild/rules_proto/archive/refs/tags/5.3.0-21.7.tar.gz",
    ],
)

###################
# Docker Support (Optional)
###################

# Download rules_docker - For building Docker images
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "b1e80761a8a8243d03ebca8845e9cc1ba6c82ce7c5179ce2b295cd36f7e394bf",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.25.0/rules_docker-v0.25.0.tar.gz"],
)

###################
# Initialize Language Rules
###################

# Initialize Go rules
load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

go_rules_dependencies()
go_register_toolchains(version = "1.19.5")  # Specify Go version
gazelle_dependencies()

# Initialize Protocol Buffer rules
load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")
rules_proto_dependencies()
rules_proto_toolchains()

# Initialize Docker rules (if using)
load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)
container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")
container_deps()

###################
# External Go Dependencies
###################

# Load external Go dependencies - typically generated by Gazelle
load("//:deps.bzl", "go_dependencies")
go_dependencies()

###################
# Custom External Dependencies
###################

# Example: Custom git repository
git_repository(
    name = "com_mycompany_internal_tools",
    remote = "https://github.com/mycompany/internal-tools.git",
    commit = "abc123def456789",  # Pin to specific commit
)

# Example: Local development dependency
local_repository(
    name = "local_dev_tools",
    path = "/path/to/local/development/tools",
)
```

### BUILD File Organization

BUILD files define how to build code in each directory. Here's how to organize them effectively:

**Root BUILD.bazel:**
```python
# BUILD.bazel (workspace root)

load("@bazel_gazelle//:def.bzl", "gazelle")

# Package metadata
package(default_visibility = ["//visibility:public"])

# Gazelle configuration
gazelle(
    name = "gazelle",
    prefix = "github.com/mycompany/myproject",
)

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

# Workspace-wide utilities
filegroup(
    name = "all_files",
    srcs = glob(["**/*"]),
    visibility = ["//visibility:public"],
)

# Linting and formatting
load("@com_github_bazelbuild_buildtools//buildifier:def.bzl", "buildifier")

buildifier(
    name = "buildifier_check",
    mode = "check",
)

buildifier(
    name = "buildifier_fix",
    mode = "fix",
)
```

**Service-level BUILD.bazel:**
```python
# services/user/BUILD.bazel

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

# Package-level visibility and metadata
package(default_visibility = ["//services:__subpackages__"])

# Main service library
go_library(
    name = "user_lib",
    srcs = [
        "service.go",
        "handler.go", 
        "repository.go",
    ],
    importpath = "github.com/mycompany/myproject/services/user",
    deps = [
        "//api/user/v1:user_api_go_proto",
        "//internal/database:db_lib",
        "//pkg/auth:auth_lib",
        "//pkg/logging:logging_lib",
        "@com_github_gorilla_mux//:mux",
        "@org_golang_google_grpc//:grpc",
    ],
)

# Unit tests
go_test(
    name = "user_test",
    srcs = [
        "service_test.go",
        "handler_test.go",
        "repository_test.go",
    ],
    embed = [":user_lib"],
    deps = [
        "//testutils:testutils_lib",
        "@com_github_stretchr_testify//assert",
        "@com_github_stretchr_testify//mock",
    ],
    data = glob(["testdata/**/*"]),
)

# Integration tests (separate target)
go_test(
    name = "user_integration_test",
    srcs = ["integration_test.go"],
    embed = [":user_lib"],
    deps = [
        "//testutils:integration_testutils",
        "@com_github_stretchr_testify//suite",
    ],
    tags = [
        "integration",
        "requires-database",
    ],
    size = "medium",
)
```

### Configuration Files

**`.bazelrc` - Build Configuration:**
```bash
# .bazelrc - Bazel configuration file

# Build settings
build --show_timestamps
build --announce_rc
build --symlink_prefix=bazel-

# Test settings  
test --test_output=errors
test --test_summary=detailed

# Go-specific settings
build --incompatible_enable_proto_toolchain_resolution

# Performance settings
build --jobs=auto
build --local_ram_resources=HOST_RAM*0.75

# Remote caching (if available)
# build --remote_cache=grpc://cache.example.com:443

# Different configurations for different environments
build:dev --compilation_mode=fastbuild
build:prod --compilation_mode=opt
build:debug --compilation_mode=dbg --strip=never

# CI-specific settings
build:ci --progress_report_interval=60
build:ci --show_progress_rate_limit=5
build:ci --curses=no
build:ci --color=no

# Try import user-specific settings
try-import %workspace%/.bazelrc.user
```

**`.bazelignore` - Exclude Directories:**
```
# .bazelignore - Directories to exclude from Bazel

# Version control
.git/
.svn/

# IDE files
.idea/
.vscode/
*.swp
*.swo

# Build artifacts from other systems
node_modules/
target/
build/
dist/

# Temporary directories
tmp/
temp/
.tmp/

# OS-specific files
.DS_Store
Thumbs.db

# Dependencies that shouldn't be built by Bazel
vendor/
third_party/some-external-tool/
```

### Directory Layout Best Practices

**Recommended project structure:**
```
my-project/
├── WORKSPACE                    # Workspace definition
├── BUILD.bazel                  # Root build file
├── .bazelrc                     # Build configuration
├── .bazelignore                 # Excluded directories
├── go.mod                       # Go module definition
├── deps.bzl                     # External dependencies (generated)
├── README.md
├──
├── api/                         # API definitions (proto files)
│   ├── BUILD.bazel
│   ├── user/
│   │   └── v1/
│   │       ├── BUILD.bazel
│   │       └── user.proto
│   └── common/
│       ├── BUILD.bazel
│       └── types.proto
├──
├── cmd/                         # Main applications
│   ├── BUILD.bazel
│   ├── server/
│   │   ├── BUILD.bazel
│   │   └── main.go
│   ├── cli/
│   │   ├── BUILD.bazel
│   │   └── main.go
│   └── migrator/
│       ├── BUILD.bazel
│       └── main.go
├──
├── internal/                    # Private application code
│   ├── BUILD.bazel
│   ├── config/
│   │   ├── BUILD.bazel
│   │   └── config.go
│   ├── database/
│   │   ├── BUILD.bazel
│   │   ├── connection.go
│   │   └── migrations/
│   └── middleware/
│       ├── BUILD.bazel
│       ├── auth.go
│       └── logging.go
├──
├── pkg/                         # Public library code
│   ├── BUILD.bazel
│   ├── auth/
│   │   ├── BUILD.bazel
│   │   ├── jwt.go
│   │   └── jwt_test.go
│   ├── logging/
│   │   ├── BUILD.bazel
│   │   └── logger.go
│   └── utils/
│       ├── BUILD.bazel
│       ├── strings.go
│       └── strings_test.go
├──
├── services/                    # Business logic services
│   ├── BUILD.bazel
│   ├── user/
│   │   ├── BUILD.bazel
│   │   ├── service.go
│   │   ├── service_test.go
│   │   └── repository.go
│   └── order/
│       ├── BUILD.bazel
│       ├── service.go
│       └── handler.go
├──
├── testutils/                   # Shared test utilities
│   ├── BUILD.bazel
│   ├── database.go
│   └── fixtures/
│       ├── BUILD.bazel
│       └── users.json
├──
├── deployments/                 # Deployment configurations
│   ├── BUILD.bazel
│   ├── docker/
│   │   ├── BUILD.bazel
│   │   └── Dockerfile
│   └── k8s/
│       ├── BUILD.bazel
│       ├── deployment.yaml
│       └── service.yaml
└──
└── docs/                        # Documentation
    ├── BUILD.bazel
    ├── api/                     # Generated API docs
    └── development.md
```

### Workspace Patterns

**1. Monorepo Pattern:**
```
company-monorepo/
├── WORKSPACE
├── projects/
│   ├── frontend/
│   │   ├── BUILD.bazel
│   │   └── [React/Angular app]
│   ├── backend/
│   │   ├── BUILD.bazel  
│   │   └── [Go services]
│   └── mobile/
│       ├── BUILD.bazel
│       └── [iOS/Android apps]
├── shared/
│   ├── protos/
│   ├── libraries/
│   └── tools/
└── third_party/
    └── [External dependencies]
```

**2. Multi-Language Pattern:**
```
polyglot-project/
├── WORKSPACE
├── go/
│   ├── BUILD.bazel
│   └── [Go services]
├── java/
│   ├── BUILD.bazel
│   └── [Java services]
├── python/
│   ├── BUILD.bazel
│   └── [Python scripts]
├── proto/
│   ├── BUILD.bazel
│   └── [Shared API definitions]
└── web/
    ├── BUILD.bazel
    └── [JavaScript/TypeScript frontend]
```

**3. Microservices Pattern:**
```
microservices-platform/
├── WORKSPACE
├── services/
│   ├── user-service/
│   ├── order-service/
│   ├── payment-service/
│   └── notification-service/
├── libraries/
│   ├── shared-models/
│   ├── auth-lib/
│   └── database-lib/
├── infrastructure/
│   ├── deployment/
│   └── monitoring/
└── tools/
    ├── code-generators/
    └── build-scripts/
```

This structure provides clear separation of concerns, makes dependencies explicit, and scales well as your project grows.

## Advanced Tips and Best Practices

This section covers advanced techniques and hard-learned lessons for using Bazel effectively in production environments.

### Performance Optimization

**1. Build Performance:**

```bash
# .bazelrc - Optimize build performance
build --jobs=auto                              # Use all available CPU cores
build --local_ram_resources=HOST_RAM*0.75     # Use 75% of available RAM
build --local_cpu_resources=HOST_CPUS*0.75    # Use 75% of available CPUs

# Enable persistent workers for faster rebuilds
build --strategy=Javac=worker
build --strategy=TypeScriptCompile=worker

# Optimize Go builds
build --incompatible_enable_proto_toolchain_resolution
build --go_config=race                         # For race condition detection
```

**2. Remote Caching:**

```bash
# .bazelrc - Remote cache configuration
build --remote_cache=grpc://cache.company.com:443
build --remote_timeout=60s
build --remote_retries=3

# Authentication (if required)
build --google_default_credentials
# OR
build --remote_header=Authorization=Bearer=your-token
```

**3. Incremental Builds:**

```python
# Use specific, fine-grained targets instead of wildcards
# ✅ Good - Only builds what's needed
bazel build //services/user:user_lib

# ❌ Avoid - Builds everything
bazel build //...

# Use build flags to control what gets built
bazel build --build_tag_filters=-slow //...     # Skip slow targets
bazel build --build_tag_filters=unit //...      # Only unit test targets
```

### Testing Strategies

**1. Test Organization:**

```python
# services/user/BUILD.bazel
go_test(
    name = "unit_test",
    srcs = ["service_test.go"],
    embed = [":user_lib"],
    deps = ["@com_github_stretchr_testify//assert"],
    tags = ["unit", "fast"],
    size = "small",
)

go_test(
    name = "integration_test", 
    srcs = ["integration_test.go"],
    embed = [":user_lib"],
    deps = [
        "//testutils:database_testutils",
        "@com_github_stretchr_testify//suite",
    ],
    tags = ["integration", "requires-database"],
    size = "medium",
    timeout = "moderate",
)

go_test(
    name = "e2e_test",
    srcs = ["e2e_test.go"],
    deps = [
        "//testutils:e2e_testutils",
    ],
    tags = ["e2e", "requires-network", "slow"],
    size = "large",
    timeout = "long",
)
```

**2. Test Execution Strategies:**

```bash
# Run different test suites
bazel test --test_tag_filters=unit //...           # Fast unit tests
bazel test --test_tag_filters=integration //...    # Integration tests  
bazel test --test_tag_filters=e2e //...           # End-to-end tests

# Parallel testing
bazel test --test_tag_filters=unit --jobs=8 //...

# Test with coverage
bazel coverage --combined_report=lcov //...

# Run tests on file changes (watch mode)
bazel test --test_tag_filters=unit //... --watchfs
```

**3. Test Data Management:**

```python
go_test(
    name = "service_test",
    srcs = ["service_test.go"],
    data = [
        "//testdata:user_fixtures",
        "//testdata:database_schema",
        ":test_config.yaml",
    ],
    env = {
        "TEST_DATA_DIR": "$(location //testdata:user_fixtures)",
        "DB_SCHEMA": "$(location //testdata:database_schema)",
    },
)

# In testdata/BUILD.bazel
filegroup(
    name = "user_fixtures",
    srcs = glob(["users/*.json"]),
    visibility = ["//visibility:public"],
)
```

### Code Organization Patterns

**1. Layer-based Architecture:**

```python
# Clear dependency layers - higher layers can depend on lower layers

# Layer 4: Presentation (HTTP handlers, gRPC services)
go_library(
    name = "handlers_lib",
    deps = [
        ":services_lib",        # Can depend on service layer
        "//pkg/validation",     # Shared utilities
    ],
)

# Layer 3: Service/Business Logic
go_library(
    name = "services_lib", 
    deps = [
        ":repositories_lib",    # Can depend on data layer
        "//pkg/auth",          # Shared utilities
    ],
)

# Layer 2: Data Access
go_library(
    name = "repositories_lib",
    deps = [
        ":models_lib",         # Can depend on model layer
        "//pkg/database",      # Shared utilities
    ],
)

# Layer 1: Models/Entities
go_library(
    name = "models_lib",
    deps = [
        # No business logic dependencies - only external libraries
        "@com_github_lib_pq//:pq",
    ],
)
```

**2. Feature-based Organization:**

```
services/
├── user/
│   ├── BUILD.bazel
│   ├── api/           # HTTP/gRPC handlers
│   ├── service/       # Business logic
│   ├── repository/    # Data access
│   └── model/         # Data structures
├── order/
│   ├── BUILD.bazel
│   ├── api/
│   ├── service/
│   ├── repository/
│   └── model/
└── shared/
    ├── BUILD.bazel
    ├── auth/
    ├── validation/
    └── errors/
```

### Security and Access Control

**1. Visibility Management:**

```python
# Restrict access to internal packages
package(default_visibility = ["//visibility:private"])

# Service-level access control
go_library(
    name = "internal_service",
    visibility = [
        "//services/user:__subpackages__",  # Only user service can access
        "//cmd/admin:__pkg__",              # Admin tool can access
    ],
)

# Public APIs
go_library(
    name = "public_api",
    visibility = ["//visibility:public"],  # Anyone can access
)

# Team-based access
go_library(
    name = "team_backend_lib",
    visibility = [
        "//teams/backend/...",              # All backend team packages
        "//teams/platform/...",             # All platform team packages
    ],
)
```

**2. Sensitive Data Handling:**

```python
# Keep secrets out of BUILD files
go_binary(
    name = "server",
    env = {
        "CONFIG_FILE": "$(location :config.yaml)",
        # Don't put secrets here!
    },
    data = [":config.yaml"],
)

# Use external configuration
go_library(
    name = "config_lib",
    srcs = ["config.go"],
    # Configuration loaded from environment/external files at runtime
)
```

### Development Workflow

**1. Local Development Setup:**

```bash
# .bazelrc.user (personal settings, not committed)
build:dev --compilation_mode=fastbuild
build:dev --copt=-g
test:dev --test_output=streamed

# Quick development aliases
alias bb='bazel build'
alias bt='bazel test'
alias br='bazel run'

# Watch mode for development
alias watch-test='bazel test --test_tag_filters=unit //... --watchfs'
```

**2. IDE Integration:**

```python
# Generate IDE-friendly output
bazel run @io_bazel_rules_go//go/tools/bazel:gazelle -- -go_prefix github.com/mycompany/myproject

# For VS Code, create .vscode/settings.json:
{
    "go.toolsGopath": "bazel-myproject/external/go_sdk",
    "go.gopath": "bazel-myproject",
    "go.buildTags": "bazel"
}
```

**3. Code Quality:**

```python
# Add linting and formatting
load("@com_github_bazelbuild_buildtools//buildifier:def.bzl", "buildifier")
load("@io_bazel_rules_go//go:def.bzl", "nogo")

# Format BUILD files
buildifier(
    name = "buildifier",
    exclude_patterns = ["./external/**"],
)

# Go linting
nogo(
    name = "nogo_check",
    vet = True,
    deps = [
        "@org_golang_x_tools//go/analysis/passes/asmdecl:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/assign:go_tool_library",
        # ... more linters
    ],
)
```

**4. Debugging Tips:**

```bash
# Understand why targets are being rebuilt
bazel build //target:name --explain=explanation.txt

# See what Bazel is doing
bazel build //target:name -s

# Debug dependency resolution
bazel query 'deps(//target:name)'

# Find reverse dependencies
bazel query 'rdeps(//..., //target:name)'

# Analyze build performance
bazel analyze-profile /tmp/profile.json
```

These advanced practices help you build maintainable, scalable, and efficient Bazel workflows that work well for teams of any size.
