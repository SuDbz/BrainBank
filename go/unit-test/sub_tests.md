# Sub-tests in Go

## Table of Contents
- [What are Sub-tests?](#what-are-sub-tests)
- [Writing Sub-tests](#writing-sub-tests)
- [Running Sub-tests](#running-sub-tests)
  - [Run All Sub-tests](#run-all-sub-tests)
  - [Run Specific Sub-tests by Name](#run-specific-sub-tests-by-name)
  - [Run Sub-tests Matching a Pattern](#run-sub-tests-matching-a-pattern)

## What are Sub-tests?
Sub-tests in Go allow you to create and run related tests grouped under a single parent test. They are particularly useful for testing multiple scenarios of a function or feature while maintaining logical organization. Sub-tests leverage Go's `testing.T` object and are typically created using `t.Run`.

### Benefits of Sub-tests:
- Logical grouping of related test cases.
- Isolated execution and reporting of each scenario.
- Ability to selectively run specific sub-tests.

## Writing Sub-tests
Here is a basic example of how to create sub-tests in Go:

```go
package main

import (
    "testing"
)

func TestAddition(t *testing.T) {
    testCases := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"both positive", 2, 3, 5},
        {"one negative", -1, 4, 3},
        {"both zero", 0, 0, 0},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := tc.a + tc.b
            if result != tc.expected {
                t.Errorf("%s: expected %d, got %d", tc.name, tc.expected, result)
            }
        })
    }
}
```

### Explanation:
1. Define test cases in a slice of structs.
2. Use a loop to iterate over test cases.
3. For each test case, call `t.Run` with a descriptive name and a test function.
4. Perform assertions or checks inside the test function.

## Running Sub-tests

### Run All Sub-tests
To run all sub-tests, simply run the parent test:

```bash
go test -v
```

### Run Specific Sub-tests by Name
You can run specific sub-tests by using the `-run` flag with a regular expression matching the test name.

#### Example:
To run only the `both positive` sub-test:

```bash
go test -v -run=TestAddition/both_positive
```

### Run Sub-tests Matching a Pattern
You can use partial matches or patterns in the `-run` flag. For example:

```bash
go test -v -run=TestAddition/both
```
This will match and run any sub-tests with names containing "both".

