# API Versioning

## Table of Contents
- API Versioning Approaches: A Comparative Summary
- How to Handle and Generate a Single Swagger Document for URL Versioning
- Handling API Deprecation in Protobuf for URL Versioning
- Handling Common OpenAPI Options with Protobuf
- Best Practices for Handling Versioned APIs with Unchanged Endpoints
- Implementation Guide: The Redirect Approach
- Links

---

## API Versioning Approaches: A Comparative Summary

| Approach                  | How It Works                                      | Example                                      | Pros                                                                 | Cons                                                        |
|---------------------------|---------------------------------------------------|----------------------------------------------|---------------------------------------------------------------------|-------------------------------------------------------------|
| URI Path Versioning       | API version in URL path                           | https://api.example.com/v1/users             | Simple, Cache-Friendly, Browser-Friendly                           | Violates REST Principles, Code Duplication                   |
| Query Parameter Versioning| Version as query parameter                        | https://api.example.com/users?version=1      | Clean URIs, Backward Compatibility                                  | Less Visible, Not Purely RESTful                             |
| Custom Header Versioning  | Version in custom HTTP header                     | X-API-Version: 2                             | Clean URIs, Flexible                                               | Hard to Test, Less Discoverable                              |
| Media Type Versioning     | Version in Accept header (media type)             | Accept: application/vnd.myapi.v2+json        | Most RESTful, Granular Control                                     | Complex Implementation, Less Intuitive                       |

---

## How to Handle and Generate a Single Swagger Document for URL Versioning

When using URL versioning with gRPC-Gateway, generate a single OpenAPI (Swagger) document that combines all API versions using the `protoc-gen-openapiv2` plugin.

### Core Concept: The "Top-Level" Protobuf File

Create a main Protobuf file (e.g., `api/swagger.proto`) that imports all version-specific service definitions and defines global metadata.

```proto
syntax = "proto3";

package api;

import "google/api/annotations.proto";
import "protoc-gen-openapiv2/options/annotations.proto";
import "api/v1/user_service.proto";
import "api/v2/user_service.proto";

option (openapiv2_swagger) = {
  info: {
    title: "User Management API";
    version: "2.0";
    description: "A combined API for user management with multiple versions.";
  };
  schemes: [HTTP, HTTPS];
  consumes: "application/json";
  produces: "application/json";
};
```

### Step 2: Generate a Single Swagger Document

```bash
protoc -I. \
    --openapiv2_out=. \
    --openapiv2_opt logtostderr=true,json_names_for_fields=true \
    --openapiv2_opt Mapi/v1/user_service.proto=path/to/project/api/v1/user_service.proto,Mapi/v2/user_service.proto=path/to/project/api/v2/user_service.proto \
    api/swagger.proto
```

### Step 3: Verifying the Combined Document

Open the generated Swagger JSON and verify that all endpoints and definitions from both v1 and v2 are present.

#### Example Snippet

```json
{
  "swagger": "2.0",
  "info": {
    "title": "User Management API",
    "version": "2.0",
    "description": "A combined API for user management with multiple versions."
  },
  "paths": {
    "/v1/users/{user_id}": { ... },
    "/v2/users/{user_id}": { ... }
  },
  "definitions": {
    "v1.User": { ... },
    "v2.User": { ... }
  }
}
```

---

## Handling API Deprecation in Protobuf for URL Versioning

Use the `deprecated` option in Protobuf and the `openapiv2_operation` annotation for OpenAPI.

### Example

```proto
service UserService {
  // The GetUser method is deprecated. Please use the v2 API instead.
  rpc GetUser(GetUserRequest) returns (User) {
    option deprecated = true;
    option (google.api.http) = { get: "/v1/users/{user_id}" };
    option (openapiv2_operation) = {
      deprecated: true;
      description: "This endpoint is deprecated. Use the v2 API for user details.";
    };
  }
}
```

After regenerating code and documentation, the deprecated endpoint will be marked in both Go code and Swagger UI.

---

## Handling Common OpenAPI Options with Protobuf

Define common options in a single Protobuf file and import it in all service files.

### Step 1: Create a Common Options Protobuf File

```proto
syntax = "proto3";

package common;

import "google/protobuf/descriptor.proto";
import "protoc-gen-openapiv2/options/annotations.proto";

extend google.protobuf.FileOptions {
  option (protoc_gen_openapiv2.options.openapiv2_swagger) openapiv2_swagger_info = {
    info: {
      title: "User Management API";
      version: "2.0";
      description: "A combined API for user management with multiple versions.";
      contact: {
        name: "User API Team";
        url: "https://github.com/api-versioning-demo";
        email: "api@example.com";
      };
    };
    schemes: HTTP;
    schemes: HTTPS;
    consumes: "application/json";
    produces: "application/json";
    responses: {
      key: "404";
      value: { description: "Resource not found"; }
    }
  };
}
```

### Step 2: Import and Use in Service Protobuf Files

```proto
import "api/common/options.proto";
option (common.openapiv2_swagger_info) = {};
```

### Step 3: Generate the Single Swagger Document

Include the mapping for the common options file in your `protoc` command.

---

## Best Practices for Handling Versioned APIs with Unchanged Endpoints

**Do not redirect v2 endpoints to v1 logic.** Instead, duplicate unchanged endpoints in the new version's definition for clarity and maintainability.

---

## Implementation Guide: The Redirect Approach

If you must implement redirection, define both v1 and v2 services, and have v2 handlers call v1 logic.

### Example Go Server

```go
type server struct {
    v1.UnimplementedUserServiceServer
    v2.UnimplementedUserServiceServer
}

func (s *server) GetUserV2(ctx context.Context, req *v2.GetUserRequest) (*v2.User, error) {
    v1Req := &v1.GetUserRequest{UserId: req.GetUserId()}
    v1User, err := s.GetUserV1(ctx, v1Req)
    if err != nil {
        return nil, err
    }
    return &v2.User{Id: v1User.Id, Name: v1User.Name}, nil
}
```

---

## Links

- [Buf Schema Registry](https://buf.build)
- [gRPC-Gateway OpenAPI Options](https://buf.build/grpc-ecosystem/grpc-gateway/file/7c1895c8bc284b96a022ec096a3ebb22:protoc-gen-openapiv2/options/openapiv2.proto)
