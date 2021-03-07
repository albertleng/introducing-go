### Structs
A struct is a type that contains named fields. For example, we could represent a circle like this:  
```go
type Circle struct {
	x float64
	y float64
	r float64
}
```
The `type` keyword introduces a new type. It's followed by the name of the type (`Circle`), the keyword `struct` to indicate that we are defining a `struct` type, and a list of fields inside of curly braces.  

### Initialization
We can create an instance of our new `Circle` type in a variety of ways:  
```go
var c Circle
```
Like with other data types, this will create a local `Circle` variable that is by default set to zero. For a `struct`, zero means each of the fields is set to their corresponding zero value (0 for `int`s, 0.0 for `float`s, `nil` for pointers, etc.) We can also use the `new` function:  
```go
C := new(Circle)
```
This allocates memory for all the fields, sets each of them to their zero value, and returns a pointer to the struct (`*Circle`). Pointers are often used with structs so that functions can modify their contents.  

Using `new` in this way is somewhat uncommon. Typically, we want to give each of the fields an initial value.  

The first option:  
```go
c := Circle{x:0, y:0, r: 5}
```
The second option:
```go
c := Circle{0, 0, 5}
```
If you want a pointer to the struct, use `&`:  
```go
c := &Circle{0, 0, 0}
```

### Fields
We can access fields using the `.` operator:  
```go
fmt.Println(c.x, c.y, c.r)
c.x = 10
c.y = 5
```
Arguments are always copied in Go. If we attempted to modify one of the fields inside of another function, it would not modify the original variable. Because of this, we would typically write the function using a pointer to the `Circle`:  
```go
func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}
```
And change `main` to use `&` before `c`:
```go
c := Circle(0, 0, 5)
fmt.Println(circleArea(&c))
```

### Methods
```go
func (c *Circle) area() float64 {
	return math.Pi * c.r*c.r
}
```
In between the keyword `func` and the name of the function, we've added a *receiver*. The receiver is like a parameter - it has a name and a type - but by creating the function in this way, it allows us to call the function using `.` operator:
```go
fmt.Pritln(c.area())
```
### Interfaces
```go
type Shape interface {
	area() float64
}
```
Like a struct, an interface is created using the `type` keyword, followed by a name and the keyword `interface`. But instead of defining fields, we define a *method set*. A method set is a list of methods that a type must have in order to *implement* the interface.  

Both `Rectangle` and `Circle` have area methods that return `float64`s, so both types implement the `Shape` interface. By itself, this wouldn't be particularly useful, but we can use interface types as arguments to functions.  

This is the problem interfaces are designed to solve. Because both of our shapes have an `area` method, they both implement the `Shape` interface and we can change our function to this:  

```go
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
        area += s.area()
    }
    return area
}
```
We would call this function like this:  
```go
fmt.Println(totalArea(&c, &r))
```
All `totalArea` knows about each shape is that it has an `area` method:  
```go
type Shape interface {
	area() float64
}
```
In Go, you generally focus more on the behavior of your program than on a taxonomy of types. Create a few small structs that do what you want, add in methods that you need, and as you build your programs, useful interfaces will tend to emerge. There's no need to have them all figured out ahead of time.  

Interfaces are particularly useful as software projects grow and become more complex. They allow us to hide the incidental details of implementations (e.g., the fields of our struct), which makes it easier to reason about software components in isolation.  

Go also has a mechanism for combining interfaces, types, variables, and functions together into a single component known as a *package*.