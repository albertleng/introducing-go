1. What is the difference between a method and a function?  
*My answer*:  `method` has a receiver which has the format of `func` `receiver` `function_name` `return_type` and its body while `function` has only `func`, `function_name` and `return_type` and its body

**Book's answer**:  The difference between a method and a function is that a method has a receiver while a function does not. 

2. Why would you use an embedded anonymous field instead of a normal named field?
*My answer*: To show the logical relationship of the field with the parent struct which is a relationship of `is-a` instead of `has-a`  
   
**Book's answer**: You would use an embedded anonymous field instead of a normal named field in order to use methods directly on the containing type.
   
3. Add a new `perimeter` method to the `Shape` interface to calculate the perimeter of a shape. Implement the method for `Circle` and `Rectangle`.  
   
*My answer*:
```go
type Shape interface {
	area float64
	perimeter float64
}

func (r *Rectangle) func perimeter() float64 {
	return r.x * r.y
}

func (c *Circle) func perimeter() float64 {
	return 2 * math.pi * c.r
}
```

**Book's answer**:  
```go
type Shape interface {
	perimeter() float64
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (c *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}
```
