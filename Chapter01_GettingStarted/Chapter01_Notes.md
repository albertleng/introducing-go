There are two types of Go programs: `executables` and `libraries`.  

The `import` keyword is how we include code from other packages to use with your program.  

Files in the `fmt` package start with package `fmt`.  

Functions are the building blocks of a Go program. All functions start with the keyword `func` followed by the name of the function, a list of zero or more `parameters` surrounded by parenthesis, an optional return type, and a `body` which is surrounded by curly braces.  

```go
fmt.Println("Hello, World)
```
This statement is made of three components. First, we access another function inside of the `fmt` package called `Println`; Then, we create a new string that contains `Hello, World` and *invoke* (a.k.a. *call* or *execute*) that function with the string as the first and only argument.  

You can find out more about `Println` function via  
```bash
godoc fmt Println
```
(*Note: Install `godoc` via `sudo apt install golang-golang-x-tools`*)  
Output:
```bash
godoc fmt Println
use 'godoc cmd/fmt' for documentation on the fmt command 

func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline
    is appended. It returns the number of bytes written and any write error
    encountered.
```