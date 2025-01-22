# Unit Testing in golang 

Unit testing is a crucial part of software development, especially in a language like Go where concurrency and performance are paramount.


### Go's Testing Framework

- testing Package : Go provides a built-in testing package that offers the necessary tools for writing and running unit tests.
- testing.T Interface : This interface is the core of the testing framework. It provides methods for:
    - t.Errorf(format, args...): Records an error and continues the test.
    - t.Fatalf(format, args...): Records a fatal error and stops the test immediately.
    - t.Log(args...): Prints a log message to the standard output.
    - t.Run(name, f func(*T)): Subtests allow you to group related tests under a single name

    Example:
     ```go
    package mypackage
    import "testing"
    func Add(x, y int) int {
            return x + y
    }
    func TestAdd(t *testing.T) {
            result := Add(2, 3)
            if result != 5 {
                    t.Errorf("Add(2, 3) = %d; want 5", result)
            }
    }
    ```
 ###  Key Concepts

 - Test Cases :   Each TestXXX function represents a single test case.
 - Assertions :  Use methods like t.Errorf and t.Fatalf to make assertions about the expected behavior of your code.
 - Test Coverage:  Measure how much of your code is executed by your tests. Tools like go test -cover can help you identify areas with low coverage.
 - Table-Driven Tests:   For repetitive tests with varying inputs and expected outputs, use table-driven tests:
    ```go 
    func TestAddTableDriven(t *testing.T) {
        tests := []struct {
                x, y, expected int
        }{
                {2, 3, 5},
                {0, 0, 0},
                {-1, 1, 0},
        }

        for _, tt := range tests {
                result := Add(tt.x, tt.y)
                if result != tt.expected {
                        t.Errorf("Add(%d, %d) = %d; want %d", tt.x, tt.y, result, tt.expected)
                }
        }
    }
    ```

 ### Running Tests

- go test: Run all tests in the current package.
- go test -v: Run tests in verbose mode, showing output for each test.
- go test -cover: Run tests and display test coverage information.   

## Documentation 

1. [Recommended Structure and Rules for a Go Unit Test File](golang_unit_test_structure.md)
2. [Mocking Objects in Go: A Guide with Examples](mocking_in_golang.md)
