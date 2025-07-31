# API Versioning: Concepts and Strategies

This page introduces the core concepts, strategies, and best practices for API versioning.

## Table of Contents
- [Why Version APIs?](#why-version-apis)
- [Common Versioning Approaches](#common-versioning-approaches)
- [Comparison Table](#comparison-table)
- [Best Practices](#best-practices)
- [Deprecation and Transition](#deprecation-and-transition)
- [Next: API Versioning in Go + Protobuf](./go_protobuf.md)

---

## Why Version APIs?
- Evolve APIs without breaking existing clients
- Support multiple client versions
- Enable safe, incremental improvements

## Common Versioning Approaches

| Approach         | Example                        | Pros                        | Cons                        |
|-----------------|--------------------------------|-----------------------------|-----------------------------|
| URI Versioning  | `/v1/resource`                 | Simple, visible             | Breaks RESTful identity     |
| Header Version  | `Accept: vnd.myapi.v1+json`    | Clean URLs, flexible        | Harder to debug             |
| Query Param     | `/resource?version=1`          | Easy to add, visible        | Not RESTful, cache issues   |
| Content Negot.  | `Accept: ...;version=1`        | Flexible, standard          | Complex, less discoverable  |

## Best Practices
- Start with versioning from v1
- Use semantic versioning for breaking changes
- Document all versions and changes
- Deprecate responsibly with clear timelines
- Communicate with clients

## Deprecation and Transition
- Announce deprecation early
- Provide migration guides
- Return warnings in deprecated versions
- Remove deprecated versions in stages

---
**Next:** [API Versioning in Go + Protobuf](./go_protobuf.md)
