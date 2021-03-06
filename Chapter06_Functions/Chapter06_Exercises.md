1. `sum` is a function that takes a slice of numbers and adds them together. What would be its function signature look like in Go?

*My answer*:  
```go
func sum(nums []int) int
```

**Book's answer**:   

2. Write a function that takes an integer and halves it and returns true if it was even or false if it was odd. For example, `half(1)` should return `(0, false)` and `half(2)` should return `(1, true)`
*My answer*: Refer to [Chapter6_Functions.go](Chapter06_functions.go)  
   

3. Write a function with one variadic parameter that finds the greatest number in a list of numbers.  
   *My answer*: Refer to [Chapter6_Functions.go](Chapter06_functions.go)


4.  Using `makeEvenGenerator` as an example, write a `makeOddGenerator` function that generates odd numbers.
    *My answer*: Refer to [Chapter6_Functions.go](Chapter06_functions.go)  
    
5. The Fibonacci sequence is defined as: `fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2)`. Write a recursive function that can find `fib(n)`.
   *My answer*: Refer to [Chapter6_Functions.go](Chapter06_functions.go)  

6. What are `defer`, `panic` and `recover`? How do you recover from a runtime panic?  
*My answer*: `defer` schedules a function call to be run after the function completes. `panic` causes a runtime error. We can handle a runtime panic with the built-in `panic` function. `recover` stops the panic and returns the value that was passed to the call to `panic`.  
   
To recover from a runtime panic, we use `recover` by pairing it with `defer`:  
```go
defer func() {
	str := recover()
	fmt.Println(str)
}()
panic("PANIC")
```

7. How do you get the memory address of a variable?  
*My answer*: Use `&` in front of the variable. For example, memory address for a variable `x` is `&x`.  
   
8. How do you assign a value to a pointer?
*My answer*: `xPtr = &x`  
   
9. How do you create a new pointer?
*My answer*: `xPtr := new(int)`
   
10. What is the value of `x` after running this program?
```go
func square(x *float64) {
	*x = *x * *x
}
func main() {
	x := 1.5
	square(&x)
}
```
*My answer*: 2.25  

11. Write a program that can swap two integers (`x := 1, y := 2; swap(&x, &y)` should give you `x=2 and y=1`).  
    *My answer*: Refer to [Chapter6_Functions.go](Chapter06_functions.go)  
