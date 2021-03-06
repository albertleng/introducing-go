A function (also known as a procedure or a subroutine) is an independent section of code that maps zero or more input parameters to zero or more output parameters.  

### Some notes:
- *Parameter names can be different*
  - The calling and callee functions are allowed to use different names for the parameters.
  
- *Variables must be passed to functions*
- *Functions form call stacks*
  - Functions are built in call a call stack.
  - Each time we call a function, we push it onto the call stack , and each time we return from a function, we pop the last function off of the stack.
    
- *Return types can have names*
- *Multiple values can be returned*
  - Multiple values are often used to return an error value along with the result `(*, err := f())`, or a boolean to indicate success `(x, ok := f())`.
    
### Variadic Functions
- This is a special form available for the last parameter in a Go function.
- By using an ellipsis(`...`) before the type name of the last parameter, you can indicate that it takes zero or more of those parameters.
- This is precisely how the `fmt.Prinltn` function is implemented:
```go
func Println(a ...interface{}) (n int, err error)
```

### Closure
- It's possible to create functions inside of functions.
- One way to use closure is by writing a function that returns another function, which when called, can generate a sequence of numbers. 

### Recursion
- Finally, a function is able to call itself.
- Closure and recursion are powerful programming technique that form the basis of a paradigm known as `functional programming`. Most people will find functional programming more difficult to understand than an approach based on for `loops`, `if` statements, variables, and simple functions.

### defer, panic and recover
- Go has a special statement called `defer` that schedules a function call to be run after the function completes.
- `defer` is often used when resources need to be freed in some way. For example, when we open a file, we need to make sure to close it later. With `defer`:
```go
f, _ := os.Open(filename)
defer f.Close()
```
- This has three advantages:
  - It keeps our `Close` call near our `Open` call so it's easier to understand.
  - If our function had multiple return statements (perhaps one in an `if` and one in an `else`), `Close` will happen before both of them.
  - Deferred functions are run even if a runtime panic occurs.
  
### panic and recover
- We can handle a runtime panic with the built-in `recover` function. `recover` stops the panic and returns the value that was passed to the call to `panic`. 
- We have to pair it with `defer`.
- A `panic` generally indicates a programmer error (e.g., attempting to access an index of an array that's out of bounds, forgetting to initialize a map, etc.) or an exceptional condition that's no easy way to recover from (hence the name *panic*).

### Pointers
Pointers reference a location in memory where a value is stored rather than the value itself.

### The * and & operators
In Go, a pointer is represented using an asterisk (`x`) followed by the type of the stored value.  

An asterisk is also used to *dereference* pointer variables. Dereferencing a pointer gives us access to the value the pointer points to. When we write `*xPtr = 0`, we are saying "store the `int 0` in the memory location `xPtr` refers to." If we try `xPtr = 0` instead, we will get a compile-time error because `xPtr` is not an `int`; it's a `*int`. which can only be given another `*int`.  

Finally, we use the `&` operator to find the address of a variable, `&x` returns a `*int` (pointer to an int) because x is an int. 

### new
`new` takes a type as an argument, allocates enough memory to fit a value of that type, and returns a pointer to it.  

In some programming languages, there is a significant difference between using `new` and `&`, with great care being needed to eventually delete anything created with `new`. You don't have to worry about this with Go - it's a garbage collected programming language, which means memory is cleaned up automatically when nothing refers to it anymore.  

Pointers are rarely used with Go's built-in types, they are extremely useful when paired with structs.