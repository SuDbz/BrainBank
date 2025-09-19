# Generating Swagger 2.0 (OpenAPI) from Protocol Buffers (Protobuf)

## Table of Contents
- [Introduction](#introduction)
- [What is Swagger 2.0 (OpenAPI)?](#what-is-swagger-20-openapi)
- [What is Protocol Buffers (Protobuf)?](#what-is-protocol-buffers-protobuf)
- [Why Generate Swagger from Protobuf?](#why-generate-swagger-from-protobuf)
- [Required Tools](#required-tools)
- [Step-by-Step Guide](#step-by-step-guide)
  - [1. Install Prerequisites](#1-install-prerequisites)
  - [2. Example Proto File](#2-example-proto-file)
  - [3. Generate Swagger 2.0 JSON](#3-generate-swagger-20-json)
  - [4. Customizing Swagger Output](#4-customizing-swagger-output)
  - [5. Advanced Proto Features](#5-advanced-proto-features)
  - [6. Troubleshooting & Common Issues](#6-troubleshooting--common-issues)
- [Options and Variables](#options-and-variables)
- [Best Practices](#best-practices)
- [References & Further Reading](#references--further-reading)

---

## Introduction
This guide provides a comprehensive walkthrough for generating a Swagger 2.0 (OpenAPI) specification from Protocol Buffers (protobuf) files. It is designed for beginners and advanced users, with detailed examples, configuration options, and troubleshooting tips.

## What is Swagger 2.0 (OpenAPI)?
Swagger 2.0 (OpenAPI 2.0) is a widely adopted specification for describing RESTful APIs. It enables both humans and computers to understand the capabilities of a service without access to source code, facilitating documentation, client generation, and testing.

- [Official Swagger 2.0 Spec](https://swagger.io/specification/v2/)

## What is Protocol Buffers (Protobuf)?
Protocol Buffers is a language-neutral, platform-neutral extensible mechanism for serializing structured data, used by gRPC and other systems. Protobuf files (`.proto`) define messages and services.

- [Protobuf Language Guide](https://developers.google.com/protocol-buffers/docs/proto3)

## Why Generate Swagger from Protobuf?
- To provide RESTful API documentation for gRPC services
- To enable API consumers to understand and test your service
- To support code generation for clients in multiple languages
- To bridge gRPC and HTTP/JSON ecosystems

## Required Tools
- **protoc**: Protocol Buffers compiler
- **protoc-gen-openapiv2**: Plugin to generate Swagger 2.0 from proto files
- **protoc-gen-grpc-gateway**: (Optional) For generating RESTful HTTP reverse-proxy

### Installation Example (macOS/Linux)
```sh
brew install protobuf # or use your OS package manager
GO111MODULE=on go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
GO111MODULE=on go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
```
Make sure `$GOPATH/bin` or `$HOME/go/bin` is in your `$PATH`.

## Step-by-Step Guide

### 1. Install Prerequisites
Install `protoc` and the required plugins as shown above. You may also need the [googleapis](https://github.com/googleapis/googleapis) repository for HTTP annotations.

### 2. Example Proto File (Basic)
Create a file named `hello.proto`:
```proto
syntax = "proto3";
package example;

import "google/api/annotations.proto";

service HelloService {
  rpc SayHello (HelloRequest) returns (HelloReply) {
    option (google.api.http) = {
      post: "/v1/hello"
      body: "*"
    };
  }
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```

#### Example: HTTP Mapping
The `google.api.http` annotation maps the gRPC method to an HTTP endpoint. You can use `get`, `post`, `put`, `delete`, etc.

### 3. Generate Swagger 2.0 JSON
Run the following command:
```sh
protoc -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --openapiv2_out . \
  hello.proto
```
This will generate `hello.swagger.json` in the current directory. If you see errors about missing imports, ensure you have the correct path to `google/api/annotations.proto` (often in a `third_party/googleapis` directory).

#### Example Output (Excerpt)
```json
{
  "swagger": "2.0",
  "info": {
    "title": "",
    "version": "",
    "description": ""
  },
  "paths": {
    "/v1/hello": {
      "post": {
  "operationId": "SayHello",
        "parameters": [
          {
            "in": "body",
            "name": "body",
            "schema": { "$ref": "#/definitions/exampleHelloRequest" }
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response",
            "schema": { "$ref": "#/definitions/exampleHelloReply" }
          },
          "default": {
            "description": "Unexpected error",
            "schema": { "$ref": "#/definitions/Error" }
          }
        }
      }
    }
  }
}
```

### 4. Customizing Swagger Output
You can customize the generated Swagger using options in your proto file. Import the options proto:
```proto
import "protoc-gen-openapiv2/options/annotations.proto";

option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = {
  info: {
    title: "Hello API"
    version: "1.0"
    description: "A simple Hello World API"
    contact: { name: "API Support", email: "support@example.com" }
    license: { name: "Apache 2.0", url: "https://www.apache.org/licenses/LICENSE-2.0.html" }
  }
  schemes: HTTP
  consumes: "application/json"
  produces: "application/json"
  tags: [ { name: "hello", description: "Hello endpoints" } ]
};
```

#### Per-Method Customization (with Error Response)
```proto
rpc SayHello (HelloRequest) returns (HelloReply) {
  option (google.api.http) = {
    post: "/v1/hello"
    body: "*"
  };
  option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
    summary: "Say hello to the user"
    description: "Returns a greeting message."
    tags: "hello"
    responses: {
      key: "200"
      value: { description: "A successful response" }
      key: "default"
      value: { description: "Unexpected error", schema: { ref: "Error" } }
    }
  };
}
```

#### Per-Message Customization
```proto
message HelloRequest {
  string name = 1 [(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = {
    description: "Name of the person to greet"
    example: '"Alice"'
  }];
}
```

### 5. Advanced Proto Features

#### Enums
```proto
enum Status {
  UNKNOWN = 0;
  ACTIVE = 1;
  INACTIVE = 2;
}

message StatusRequest {
  Status status = 1;
}
```
Enums are mapped to string values in Swagger. You can add descriptions and examples using field options.

#### Repeated Fields
```proto
message BatchRequest {
  repeated string ids = 1;
}
```
Repeated fields are mapped to arrays in Swagger.

#### Nested Messages
```proto
message Parent {
  Child child = 1;
  message Child {
    string value = 1;
  }
}
```

#### File Uploads
Swagger 2.0 supports file uploads via `type: file`. In proto, use `bytes` and document accordingly.

#### Authentication
You can specify security definitions in the Swagger options:
```proto
option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = {
  security_definitions: {
    security: {
      key: "ApiKeyAuth"
      value: {
        type: TYPE_API_KEY
        in: IN_HEADER
        name: "X-API-KEY"
        description: "API Key authentication"
      }
    }
  }
  security: [ { security_requirement: { key: "ApiKeyAuth" } } ]
};
```

### 6. Troubleshooting & Common Issues

- **Enums not showing as strings**: Use the latest plugin and check your proto syntax.
- **Missing HTTP annotations**: Ensure `google/api/annotations.proto` is imported and available.
- **Required fields**: Proto3 does not support required fields; use comments and documentation.
- **Header parameters**: Use `additional_bindings` for custom headers.
- **Validation**: Use [Swagger Editor](https://editor.swagger.io/) to validate the generated JSON.
- **Plugin errors**: Check your `$PATH` and plugin installation.

## Options and Variables

### openapiv2_swagger (File-level Swagger Options)
This option customizes the top-level Swagger document. Place it at the top of your proto file:

```proto
import "protoc-gen-openapiv2/options/annotations.proto";

option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = {
  info: {
    title: "My API"
    version: "1.0.0"
    description: "Detailed API description."
    terms_of_service: "https://example.com/terms"
    contact: {
      name: "API Team"
      url: "https://example.com/contact"
      email: "support@example.com"
    }
    license: {
      name: "Apache 2.0"
      url: "https://www.apache.org/licenses/LICENSE-2.0.html"
    }
  }
  host: "api.example.com"
  base_path: "/v1"
  schemes: HTTPS
  consumes: "application/json"
  produces: "application/json"
  tags: [
    { name: "user", description: "User operations" },
    { name: "admin", description: "Admin operations" }
  ]
  external_docs: {
    url: "https://docs.example.com"
    description: "Find more info here"
  }
  security_definitions: {
    security: {
      key: "ApiKeyAuth"
      value: {
        type: TYPE_API_KEY
        in: IN_HEADER
        name: "X-API-KEY"
        description: "API Key auth"
      }
    }
  }
  security: [ { security_requirement: { key: "ApiKeyAuth" } } ]
  vendor_extension: [
    { key: "x-custom-extension", value: { string_value: "custom" } }
  ]
};
```

**Key fields:**
- `info`: Metadata about the API (title, version, description, contact, license, terms_of_service)
- `host`: API host (e.g., api.example.com)
- `base_path`: Base path for all endpoints
- `schemes`: Supported protocols (HTTP, HTTPS)
- `consumes`/`produces`: MIME types
- `tags`: Group endpoints
- `external_docs`: Link to external documentation
- `security_definitions`/`security`: Auth config
- `vendor_extension`: Custom extensions (x-*)

See the [full proto definition](https://buf.build/grpc-ecosystem/protoc-gen-openapiv2/docs/main:grpc.gateway.protoc_gen_openapiv2.options#openapiv2swagger) for all fields.

---

### openapiv2_operation (Per-method Swagger Options)
This option customizes the Swagger operation for a specific RPC method.

```proto
rpc CreateUser (CreateUserRequest) returns (User) {
  option (google.api.http) = {
    post: "/v1/users"
    body: "*"
  };
  option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
    summary: "Create a new user"
    description: "Creates a user and returns the user object."
    operation_id: "CreateUser"
    tags: "user"
    deprecated: false
    responses: {
      key: "201"
      value: { description: "User created successfully" }
    }
    external_docs: {
      url: "https://docs.example.com/users#create"
      description: "More info"
    }
    security: [ { security_requirement: { key: "ApiKeyAuth" } } ]
    vendor_extension: [
      { key: "x-operation-extension", value: { string_value: "value" } }
    ]
  };
}
```

**Key fields:**
- `summary`: Short summary of the operation
- `description`: Detailed description
- `operation_id`: Unique string for the operation
- `tags`: Grouping tags
- `deprecated`: Mark as deprecated
- `responses`: Map of HTTP status code to response description
- `external_docs`: Link to more info
- `security`: Security requirements
- `vendor_extension`: Custom extensions

See the [full proto definition](https://buf.build/grpc-ecosystem/protoc-gen-openapiv2/docs/main:grpc.gateway.protoc_gen_openapiv2.options#openapiv2operation) for all fields.

---

### openapiv2_schema (Per-message Swagger Schema Options)
This option customizes the Swagger schema for a message type.

```proto
message User {
  option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_schema) = {
    json_schema: {
      title: "User"
      description: "A user object in the system."
      required: ["id", "email"]
      read_only: false
      example: '{"id": "123", "email": "user@example.com"}'
      default: '{"active": true}'
      multiple_of: 1
      maximum: 1000
      exclusive_maximum: false
      minimum: 1
      exclusive_minimum: false
      max_length: 255
      min_length: 1
      pattern: "^[a-zA-Z0-9]+$"
      max_items: 10
      min_items: 1
      unique_items: true
      enum: ["ADMIN", "USER"]
    }
    external_docs: {
      url: "https://docs.example.com/user-schema"
      description: "User schema docs"
    }
    example: '{"id": "123", "email": "user@example.com"}'
    default: '{"active": true}'
    vendor_extension: [
      { key: "x-schema-extension", value: { string_value: "value" } }
    ]
  };

  string id = 1;
  string email = 2;
  bool active = 3;
}
```

**Key fields:**
- `json_schema`: JSON schema options (title, description, required, read_only, example, default, validation constraints)
- `external_docs`: Link to more info
- `example`: Example value
- `default`: Default value
- `vendor_extension`: Custom extensions

See the [full proto definition](https://buf.build/grpc-ecosystem/protoc-gen-openapiv2/docs/main:grpc.gateway.protoc_gen_openapiv2.options#openapiv2schema) for all fields.

---

- **openapiv2_field**: Per-field options (description, example, default, required, etc.)
- **google.api.http**: Maps gRPC methods to HTTP endpoints (get, post, put, delete, custom verbs)

See [protoc-gen-openapiv2 options reference](https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/customizing_openapi_output/) and [Buf Docs](https://buf.build/grpc-ecosystem/protoc-gen-openapiv2/docs/main:grpc.gateway.protoc_gen_openapiv2.options) for all available options and field details.

## Best Practices

- Use clear descriptions and examples in your proto comments and options
- Group related endpoints with tags
- Use enums and repeated fields carefully (see [known issues](https://github.com/grpc-ecosystem/grpc-gateway/issues))
- Validate the generated Swagger using [Swagger Editor](https://editor.swagger.io/)
- Keep your proto and Swagger documentation in sync
- Use versioning in your API paths (e.g., `/v1/`)
- Document authentication and error responses clearly

## Practical Examples, Use Cases & Required Items

These examples fill gaps commonly needed when producing production-ready OpenAPI output from protobuf.

### Security Examples

1) API Key (Header)

```proto
option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = {
  security_definitions: {
    security: {
      key: "ApiKeyAuth"
      value: {
        type: TYPE_API_KEY
        in: IN_HEADER
        name: "X-API-KEY"
        description: "API Key authentication"
      }
    }
  }
  security: [ { security_requirement: { key: "ApiKeyAuth" } } ]
};
```

2) OAuth2 (Password Flow example)

```proto
option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = {
  security_definitions: {
    security: {
      key: "oauth2"
      value: {
        type: TYPE_OAUTH2
        flow: FLOW_PASSWORD
        authorization_url: "https://auth.example.com/authorize"
        token_url: "https://auth.example.com/token"
        scopes: { key: "read", value: "Read access" }
      }
    }
  }
  security: [ { security_requirement: { key: "oauth2", value: ["read"] } } ]
};
```

### Error Model (Recommended)

Define a common error message and reference it in responses to make client handling predictable.

```proto
message Error {
  string code = 1; // machine-readable error code
  string message = 2; // human-friendly message
  map<string, string> details = 3; // optional details
}

service UserService {
  rpc GetUser (GetUserRequest) returns (User) {
    option (google.api.http) = { get: "/v1/users/{id}" };
    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
      responses: {
        key: "404"
        value: { description: "User not found", schema: { ref: "Error" } }
      }
    };
  }
}
```

Note: The `schema` reference above needs to use the generated definition name (e.g., `#/definitions/Error`) in the Swagger JSON; the plugin helps wire this automatically when using the standard message type.

### additional_bindings (Multiple HTTP Mappings)

To support multiple HTTP verbs / paths for the same RPC, use `additional_bindings`:

```proto
rpc UpdateUser (UpdateUserRequest) returns (User) {
  option (google.api.http) = {
    put: "/v1/users/{id}"
    body: "user"
    additional_bindings: {
      patch: "/v1/users/{id}"
      body: "user"
    }
  };
}
```

This generates both `PUT /v1/users/{id}` and `PATCH /v1/users/{id}` in the final Swagger.

#### File Uploads

gRPC/proto doesn't have a native `file` type. Use `bytes` for binary payloads and annotate appropriately for consumers. For best Swagger UI compatibility, use a POST endpoint and document the field:

```proto
message UploadRequest {
  bytes content = 1 [(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = {
    description: "File content (base64-encoded)"
    format: "byte"
    example: "...base64..."
  }];
  string filename = 2 [(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = {
    description: "Original filename"
  }];
}

rpc Upload (UploadRequest) returns (UploadResponse) {
  option (google.api.http) = {
    post: "/v1/upload"
    body: "*"
  };
}
```

In Swagger, this will appear as a string field. If you want a true file upload in Swagger UI, you may need to post-process the generated Swagger JSON to set `type: file` for the parameter.

### Streaming RPCs

Streaming gRPC methods (client, server, bidi) do not map cleanly to HTTP/JSON. The generator omits streaming endpoints from the OpenAPI output. Document streaming methods in your API docs and provide gRPC client usage examples. Example:

```proto
rpc Chat (stream ChatMessage) returns (stream ChatMessage);
```

Add a note in `openapiv2_operation.description` such as:

```proto
option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
  description: "This is a bidirectional streaming RPC. Use a gRPC client to interact with this endpoint."
};
```

### openapiv2_field (Field-level annotations)

Use `openapiv2_field` to document individual fields, including examples and constraints.

```proto
message CreateUserRequest {
  string email = 1 [(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = {
    description: "User email address"
    example: '"user@example.com"'
  }];
  string password = 2 [(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = {
    description: "Plain-text password (store hashed)."
    min_length: 8
  }];
}
```

### Vendor Extensions (x-*)

Vendor extensions allow adding custom metadata to generated Swagger. Use `vendor_extension` in the proto options to inject `x-*` fields.

```proto
option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = {
  vendor_extension: [ { key: "x-company-meta", value: { string_value: "acme" } } ]
};
```


### Buf, Docker, and CI Integration

#### 1) Using `buf` to generate OpenAPI

Create a `buf.yaml` and a `buf.gen.yaml` with a plugin configuration. Example `buf.gen.yaml`:

```yaml
version: v1
plugins:
  - name: openapiv2
    out: gen/openapiv2
    opt: logtostderr=true
```

Run:
```sh
buf generate
```

#### 2) Docker example

Use a lightweight container with `protoc` and Go tools installed. Example Dockerfile:

```dockerfile
FROM golang:1.20-alpine
RUN apk add --no-cache protoc bash
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
ENV PATH="$PATH:/go/bin"
WORKDIR /workspace
```

Build and run the container to generate Swagger without installing tools locally.

#### 3) CI validation (GitHub Actions example)

Add a step that runs generation and validates the JSON with `swagger-cli`:

```yaml
- name: Generate OpenAPI
  run: |
    protoc -I. -Ithird_party/googleapis --openapiv2_out=gen proto/*.proto

- name: Validate OpenAPI
  run: |
    npm install -g @apidevtools/swagger-cli
    swagger-cli validate gen/*.swagger.json
```


### Validation & Testing

- Use `swagger-cli validate` or the [Swagger Editor](https://editor.swagger.io/) to ensure your generated Swagger is valid.
- Add round-trip tests: generate from proto, then run a linter/validator in CI.
- If you use custom vendor extensions or advanced options, always check the generated Swagger for compatibility with your API gateway or documentation tools.


### Quick 'Try It' Commands

```sh
# generate
protoc -I. -I$GOPATH/src -Ithird_party/googleapis --openapiv2_out . api/*.proto

# validate
swagger-cli validate api.swagger.json

# open in editor (macOS)
open https://editor.swagger.io/ # then paste the JSON
```

---



## Troubleshooting Appendix

- **Plugin not found**: Ensure `$PATH` includes the directory where `protoc-gen-openapiv2` is installed (often `$HOME/go/bin`).
- **Missing google/api/annotations.proto**: Download [googleapis](https://github.com/googleapis/googleapis) and add `-Ithird_party/googleapis` to your `protoc` command.
- **Swagger UI not showing file upload**: The plugin maps `bytes` to `string` (base64). For true file upload, post-process the Swagger JSON to set `type: file`.
- **Streaming endpoints missing**: This is expected; document streaming APIs in prose.
- **Required fields**: Proto3 does not support required fields. Use `openapiv2_field` and `json_schema.required` for documentation only.

## References & Further Reading

- [Swagger 2.0 Specification](https://swagger.io/specification/v2/)
- [gRPC-Gateway OpenAPI Docs](https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/customizing_openapi_output/)
- [protoc-gen-openapiv2 Options](https://buf.build/grpc-ecosystem/protoc-gen-openapiv2/docs/main:grpc.gateway.protoc_gen_openapiv2.options)
- [Documenting gRPC with OpenAPI](https://rotational.io/blog/documenting-grpc-with-openapi/)
- [gRPC-Gateway GitHub](https://github.com/grpc-ecosystem/grpc-gateway)
- [Google API HTTP Annotations](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto)
- [Swagger Editor](https://editor.swagger.io/)
- [Buf Schema Registry](https://buf.build/)