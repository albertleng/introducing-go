A variable is a storage location, with a specific type and an associated name.  

Variables in Go are created by first using the `var` keyword, then specifying the variable name (`x`) and the type (`string`), and finally, assigning a value to the variable (`Hello, World`).  

Because creating a new variable with a starting value is so common, Go also supports a shorter statement:  
```go
x := "Hello, World"
```

Generally, you should use this shorter form whenever possible.  

### Idiomatic Go

### How to Name a Variable
Naming a variable properly is an important part of software development. Names must start with a letter and may contain letters, numbers, or the underscore symbols (_).  

### Scope
The range of places where you are allowed to use a variable is called the scope of the variable.  

Go is lexically scoped using blocks. Basically, this means that the variable exists within the nearest curly braces ({ }), or block, including any nested curly braces (blocks), but not outside of them.  

### Constant
Constants are essentially variables whose values cannot be changed later. They are created in the same way you create variables, but instead of using `var` keyword we use `const` keyword.  

Constants are a good way to reuse common values in a program without writing them out each time. 

### Defining Multiple Variables
Go also has another shorthand when you need to define multiple variables:  

```go
var (
  a = 5
  b = 10
  c = 15
)
```
Use the keyword `var` (or `const`) followed by parentheses with each variable on its own line.  



