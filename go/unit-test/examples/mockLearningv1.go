package main

import "fmt"
type Math interface {
    Add(a, b int) int
    Multiply(a, b int) int
}

type MathImpl struct{}

// Add performs addition of two integers.
func (m *MathImpl) Add(a, b int) int {
    return a + b
}

// Multiply performs multiplication of two integers.
func (m *MathImpl) Multiply(a, b int) int {
    return a * b
}

func Calculator(m Math, a, b int) int {
    sum := m.Add(a, b)
    product := m.Multiply(a, b)
    return sum + product
}


func main() {
    math := MathImpl{}

    a, b := 3, 5

    sum := math.Add(a, b)
    product := math.Multiply(a, b)

    fmt.Printf("Add(%d, %d) = %d\n", a, b, sum)
    fmt.Printf("Multiply(%d, %d) = %d\n", a, b, product)
}