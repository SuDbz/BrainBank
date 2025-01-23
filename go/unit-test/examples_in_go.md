
# Examples in Go

## Table of Contents
- [What are Examples?](#what-are-examples)
- [Writing Examples](#writing-examples)
- [Running Examples](#running-examples)
  - [Run All Examples](#run-all-examples)
  - [Verify Output](#verify-output)
  - [Benefits of Examples](#benefits-of-examples)

## What are Examples?
Examples in Go are special test-like functions that demonstrate how to use a specific package, function, or feature. They serve as both documentation and executable tests.

### Characteristics of Examples:
- Must be named `ExampleXyz`, where `Xyz` matches the function or feature being demonstrated.
- Typically contain `fmt.Println` to showcase output.
- The output of the example is compared against the comment `// Output:` to verify correctness.

## Writing Examples
Here is an example of an example function:

```go
package main

import "fmt"

func Add(a, b int) int {
    return a + b
}

func ExampleAdd() {
    fmt.Println(Add(2, 3))
    // Output: 5
}
```

### Explanation:
1. The function name starts with `Example` and optionally includes a suffix (`ExampleAdd`).
2. Use `fmt.Println` to output the result.
3. Add a comment starting with `// Output:` to specify the expected output.

## Running Examples

### Run All Examples
Examples are executed along with tests:

```bash
go test -v
```

### Verify Output
The `// Output:` comment ensures the example's output matches the expected result. If the output differs, the test fails.

### Benefits of Examples
- Double as documentation for functions or packages.
- Provide users with working usage patterns.
- Ensure the example code remains correct over time.

---

By using sub-tests and examples effectively, you can create robust test suites while maintaining clear and reusable documentation for your Go projects.

