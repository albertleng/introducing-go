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
