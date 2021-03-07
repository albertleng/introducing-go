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