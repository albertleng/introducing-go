Go was designed to be a language that encourages good software engineering practices. An important part of high-quality software is code reuse - embodied in the principle "Don't Repeat Yourself."  

Functions are the first layer we utilize to allow code reuse. Go also provides another mechanism for code reuse: packages. 
```go
import "fmt"
```
`fmt` is the name of a package that includes a variety of functions related to formatting and output to the screen. Bundling code in this way serves three purposes:  
- It reduces the chance of having overlapping names, and in turn keeps our function names short and succinct.
- It organizes code so that it's easier to find code you want to reuse.
- It speeds up the compiler by only requiring recompilation of smaller chunks of a program. Although we use the package `fmt`, we don't have to recompile it every time we change our program.  

## The Core Packages
Instead of writing everything from scratch, most real-world programming depends on our ability to interface with existing libraries.  

### Strings
Go includes a large number of functions to work with strings in the `strings` package.  

`Contains`, `Count`, `HasPrefix`, `HasSuffix`, `Index`, `Join`, `Repeat`, `Replace`, `Split`, `ToLower`, `ToUpper`, 

### Input/Output
The `io` package consists of a few functions, but mostly interfaces used in other packages. The two main interfaces are `Reader` and `Writer`. `Reader`s support reading via the `Read` method. `Writer`s support writing via the `Write` method. Many functions in Go take `Reader`s and `Writer`s as arguments. For example, the `io` package has a `Copy` function that copies data from a `Reader` to a `Writer`:
```go
func Copy(dst Writer, src Reader) (written int64, err error)
```
To read or write to a `[]byte` or a `string`, you can use the `Buffer` struct found in the `bytes` package:  
```go
var buf bytes.Buffer
buf.Write ([]byte("test"))
```
A `Buffer` doesn't have to be initialized, and it supports both the `Reader` and `Writer` interfaces. You can convert it into a `[]byte` by calling `buf.Bytes()`. If you only need to read from a string, you can also use the `strings.NewReader` function, which is more efficient thatn using a buffer.

### Files and Folders
To open a file in Go, use the `Open` function from the `os` package.