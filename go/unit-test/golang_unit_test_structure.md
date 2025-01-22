# Recommended Structure and Rules for a Go Unit Test File

## Structure

1. **Package Declaration**
   - Test files should belong to the same package as the code being tested, or use the `<package_name>_test` convention for black-box testing.

   ```go
   package yourpackage_test
   ```

2. **Imports**
   - Import the `testing` package.
   - Import any other required packages explicitly.

   ```go
   import (
       "testing"
       "github.com/stretchr/testify/assert" // Example for assertion
   )
   ```

3. **Test Function Naming**
   - Use the `TestXxx` naming convention where `Xxx` describes the function being tested.
   - Ensure function names are descriptive.

   ```go
   func TestAddNumbers(t *testing.T) {
       // Test logic here
   }
   ```

4. **Test Cases**
   - Use table-driven tests for better readability and maintainability.
   - Include test cases for edge cases, happy paths, and error scenarios.

   ```go
   func TestAddNumbers(t *testing.T) {
       testCases := []struct {
           name     string
           a, b     int
           expected int
       }{
           {name: "Positive numbers", a: 2, b: 3, expected: 5},
           {name: "Negative numbers", a: -2, b: -3, expected: -5},
           {name: "Mixed numbers", a: -2, b: 3, expected: 1},
       }

       for _, tc := range testCases {
           t.Run(tc.name, func(t *testing.T) {
               result := AddNumbers(tc.a, tc.b)
               if result != tc.expected {
                   t.Errorf("expected %d, got %d", tc.expected, result)
               }
           })
       }
   }
   ```

5. **Setup and Teardown**
   - Use helper functions or setup/teardown logic for reusable test initialization.

   ```go
   func setup() {
       // Initialization code
   }

   func teardown() {
       // Cleanup code
   }

   func TestMain(m *testing.M) {
       setup()
       code := m.Run()
       teardown()
       os.Exit(code)
   }
   ```

6. **Benchmarks (Optional)**
   - Include benchmark tests to measure performance if required.

   ```go
   func BenchmarkAddNumbers(b *testing.B) {
       for i := 0; i < b.N; i++ {
           AddNumbers(2, 3)
       }
   }
   ```

7. **Example Tests (Optional)**
   - Provide examples to be included in GoDoc.

   ```go
   func ExampleAddNumbers() {
       fmt.Println(AddNumbers(2, 3))
       // Output: 5
   }
   ```

## Rules

1. **Isolation**
   - Each test should be independent of others.

2. **Assertions**
   - Use assertions for better readability and debugging.
   - Example: `github.com/stretchr/testify/assert`.

3. **Error Handling**
   - Validate and handle errors appropriately in tests.

   ```go
   result, err := SomeFunction()
   if err != nil {
       t.Fatalf("unexpected error: %v", err)
   }
   ```

4. **Coverage**
   - Aim for high code coverage but prioritize meaningful tests over 100% coverage.

5. **Logging**
   - Use `t.Log` or `t.Logf` for informative logging in tests.

6. **Parallel Tests**
   - Use `t.Parallel()` for tests that can run concurrently.

   ```go
   func TestSomething(t *testing.T) {
       t.Parallel()
       // Test logic
   }
   ```

7. **Avoid Global State**
   - Minimize reliance on global variables to ensure test reliability.

8. **Use Mocks/Stubs**
   - Mock dependencies or external services for isolated testing.

   ```go
   type MockService struct {}

   func (m *MockService) DoSomething() string {
       return "mocked result"
   }
   ```

9. **Fail Early**
   - Use `t.FailNow` or `t.Fatalf` for critical failures where subsequent checks are irrelevant.

10. **CI Integration**
    - Ensure tests are runnable in CI/CD pipelines without manual intervention.
