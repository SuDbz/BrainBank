# Interface
 
A powerful mechanism for defining the behavior of objects without specifying their concrete implementation. They provide a contract that types must adhere to, allowing for polymorphism and loose coupling

#### Key characteristics of Go interfaces
- Implicit implementation: Types implicitly implement an interface if they satisfy all of its methods. There's no explicit declaration required.
- Empty interfaces: The interface{} type is the empty interface and can represent any value.
- Type assertions: You can check if a value implements a specific interface using type assertions.
- Method sets: Each type has a method set, which is the set of methods it can call.

<br><br>
```go
type Shape interface {
        Area() float64
}

type Rectangle struct {
        width, height float64
}

func (r Rectangle) Area() float64 {
        return r.width   
 * r.height
}

type Circle struct {
        radius float64
}

func (c Circle) Area() float64 {
        return   
 3.14 * c.radius * c.radius
}

func main() {
        shapes := []Shape{Rectangle{width: 5, height: 3}, Circle{radius: 2}}
        for _, shape := range shapes {
                fmt.Println(shape.Area())   

        }
}
```

#### Combining Interfaces in Go

Go doesn't have a direct mechanism to combine interfaces in the same way as inheritance in object-oriented languages. However, you can achieve similar functionality using different techniques:

##### Example - 
Embed one interface within another by declaring it as a field without a name.
```go
type ReaderWriter interface {
    io.Reader
    io.Writer
}
```



#### Reference
 - https://www.youtube.com/watch?v=SX1gT5A9H-U
