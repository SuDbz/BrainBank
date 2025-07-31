# Migration and Compatibility

This page covers strategies for migrating between API versions and ensuring compatibility.

## Table of Contents
- [Migration Phases](#migration-phases)
- [Fallback and Compatibility](#fallback-and-compatibility)
- [Testing Both Versions](#testing-both-versions)
- [Summary Table](#summary-table)
- [Next: References](./references.md)

---

## Migration Phases
1. **Experimental:** V2 is experimental, V1 is stable
2. **Stable:** V2 becomes stable, V1 is legacy
3. **Deprecated:** V1 is deprecated but still available

## Fallback and Compatibility
- Use fallback logic in clients to try V2, then V1 if needed
- Convert requests/responses between versions as needed

**Example:**
```go
type HealthService struct {
    v1Client healthv1.HealthServiceClient
    v2Client healthv2.HealthServiceClient
}

func (s *HealthService) Check(ctx context.Context, req *healthv2.HealthCheckRequest) (*healthv2.HealthCheckResponse, error) {
    resp, err := s.v2Client.Check(ctx, req)
    if err != nil {
        v1Req := convertToV1Request(req)
        v1Resp, v1Err := s.v1Client.Check(ctx, v1Req)
        if v1Err != nil {
            return nil, v1Err
        }
        return convertToV2Response(v1Resp), nil
    }
    return resp, nil
}
```

## Testing Both Versions
- Use Bazel `test_suite` to test v1, v2, and compatibility

## Summary Table
| Phase         | Alias Example         | Description                  |
|--------------|----------------------|------------------------------|
| Experimental  | health_experimental  | V2 is new, V1 is stable      |
| Stable        | health_stable        | V2 is stable, V1 is legacy   |
| Deprecated    | health_deprecated    | V1 is deprecated, V2 is main |

---
**Next:** [References](./references.md)
