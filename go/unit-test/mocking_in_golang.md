# Mocking Objects in Go: A Guide with Examples

## Table of Contents
- [Why Use Mocking?](#why-use-mocking)
- [Mocking Strategies in Go](#mocking-strategies-in-go)
  - [Manual Mocking](#1-manual-mocking)
  - [Using Mock Libraries](#2-using-mock-libraries)
  - [Functional Mocking](#3-functional-mocking)
- [Summary](#summary)

---

## Why Use Mocking?
1. **Isolation**: Tests focus only on the functionality of the code under test, not on its dependencies.
2. **Repeatability**: Avoid flaky tests caused by real dependencies (e.g., network or database).
3. **Performance**: Mocks are faster than real dependencies.
4. **Edge Cases**: Simulate scenarios that are difficult to reproduce in a real environment.

---

## Mocking Strategies in Go

### 1. **Manual Mocking**
Manual mocking involves creating mock implementations of interfaces by hand. This approach works best for small and straightforward dependencies.

#### Example: Manual Mocking
```go
package main

import "testing"

// Define an interface
type Greeter interface {
    Greet(name string) string
}

// Real implementation
type RealGreeter struct{}

func (rg *RealGreeter) Greet(name string) string {
    return "Hello, " + name
}
}

// Mock implementation
type MockGreeter struct{}

func (mg *MockGreeter) Greet(name string) string {
    return "Mocked greeting for " + name
}
}

// Function under test
func SayHello(g Greeter, name string) string {
    return g.Greet(name)
}

// Test using the mock
func TestSayHello(t *testing.T) {
    mock := &MockGreeter{}
    result := SayHello(mock, "Alice")

    expected := "Mocked greeting for Alice"
    if result != expected {
        t.Errorf("expected %s, got %s", expected, result)
    }
}
```

### Why Choose Manual Mocking?
- **Simplicity**: Suitable for small, lightweight projects.
- **Full Control**: You define exactly how the mock behaves.
- **No Dependencies**: Avoids introducing external libraries.

---

### 2. **Using Mock Libraries**
Mock libraries like [gomock](https://github.com/golang/mock) or [testify/mock](https://pkg.go.dev/github.com/stretchr/testify/mock) automate mock creation and provide additional features like assertions.

#### Example: Using GoMock

1. Install GoMock:
   ```bash
   go install github.com/golang/mock/mockgen@latest
   ```

2. Generate mocks:
   ```bash
   mockgen -source=greeter.go -destination=mocks/greeter_mock.go -package=mocks
   ```

3. Example Code:
```go
package main

import (
    "mocks" // Path to generated mocks
    "testing"
    "github.com/golang/mock/gomock"
)

func TestSayHelloWithGoMock(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockGreeter := mocks.NewMockGreeter(ctrl)
    mockGreeter.EXPECT().Greet("Bob").Return("Mocked greeting for Bob")

    result := SayHello(mockGreeter, "Bob")
    expected := "Mocked greeting for Bob"

    if result != expected {
        t.Errorf("expected %s, got %s", expected, result)
    }
}
```

#### Example: Using Testify Mock
```go
package main

import (
    "testing"
    "github.com/stretchr/testify/mock"
)

type Call struct {
    mock.Mock
}

func (c *Call) Foo(a int, b string) string {
    args := c.Called(a, b)
    return args.String(0)
}

func TestFoo(t *testing.T) {
    // Create the mock object
    c := new(Call)

    // Set up expectations
    c.On("Foo", 2, "bar").Return("Hello, world!")

    // Call the method and check the return value
    result := c.Foo(2, "bar")

    // Output the result
    t.Log(result) // Outputs: Hello, world!

    // Assert that all expectations were met
    c.AssertExpectations(t) // Here we pass `t`, the testing object
}
```

### Why Choose Mock Libraries?
- **Automation**: Automatically generates mock code.
- **Advanced Features**: Includes support for method call ordering and argument matching.
- **Scalability**: Handles large and complex projects efficiently.

---

### 3. **Functional Mocking**
This approach involves injecting functions instead of interfaces. It works well for simpler dependencies or single-method interfaces.

#### Example: Functional Mocking
```go
package main

import "testing"

// Function type
type GreeterFunc func(name string) string

// Function under test
func SayHelloFunc(greet GreeterFunc, name string) string {
    return greet(name)
}

// Test using a mock function
func TestSayHelloFunc(t *testing.T) {
    mockGreet := func(name string) string {
        return "Mocked greeting for " + name
    }

    result := SayHelloFunc(mockGreet, "Charlie")
    expected := "Mocked greeting for Charlie"

    if result != expected {
        t.Errorf("expected %s, got %s", expected, result)
    }
}
```
Methods That Cannot Be Mocked 
Reassigning a function directly within the test, this approach has limitations:

- Package-Level Functions: Functions declared at the package level cannot be reassigned within a function or test. This is because they are not variables that can be reassigned.
- Methods on Concrete Types: Methods defined on concrete types (structs) cannot be easily replaced or mocked without using interfaces.
 


### Why Choose Functional Mocking?
- **Simplicity**: Best for single-method interfaces or lightweight dependencies.
- **Flexibility**: No need to define new structs or interfaces.
- **Less Boilerplate**: Reduces code overhead compared to traditional interfaces.

---

## Summary

| Approach           | Use Case                                    | Pros                              | Cons                                  |
|--------------------|---------------------------------------------|-----------------------------------|---------------------------------------|
| Manual Mocking     | Small, simple projects                     | Simple, no dependencies           | Repetitive for complex interfaces     |
| Mock Libraries     | Large, complex projects                    | Automated, feature-rich           | Adds dependency                       |
| Functional Mocking | Lightweight, single-method dependencies    | Minimal boilerplate, flexible     | Limited to simple cases               |

Choose the approach that best fits your project's complexity, team familiarity, and testing requirements.

