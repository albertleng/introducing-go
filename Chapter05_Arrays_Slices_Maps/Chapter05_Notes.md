### Arrays
An array is a numbered sequence of elements of a single type with a fixed length. In Go, they look like this:  
```go
var x [5] int
```

The Go compiler won't allow you to create variables that you never use.  

Because of its fixed length and other limitations, you rarely see arrays used in Go code. Instead, you will usually use a `slice`, which is a type built on top of an array.

### Slices
A slice is a segment of an array. Unlike arrays, its length is allowed to change.  
```go
var x []float64
```

If you want to create a slice, you should use the built-in `make` function:  
```go
x := make([]float64, 5)
```
This creates a slice that is associated with an underlying `float64` array of length 5. Slices are always associated with some array, and although they can never be longer than the array, they can be smaller. The `make` function also allows a third parameter:  
```go
x := make([]float64, 5, 10)
```
This creates a slice of length 5 with a capacity of 10.  

Another way to create slices is to use the `[low : high]` expression:  
```go
arr := [5]float64{1, 2, 3, 4, 5}
x := arr[0:5]
```
`low` is the index of where to start the slice and `high` is the index of where to end it (but not including the index itself).  

Go includes two built-in functions to assist with slices: `append` and `copy`.

#### append
`append` adds elements onto the end of a slice. 

#### copy
`copy` takes two arguments: `dst` and `src`. All of the entries in `src` are copied into `dst` overwriting whatever is there.

### Maps
A `map` is an unordered collection of key-value pairs (maps are also sometimes called associative arrays, hash tables, or dictionaries). Maps are used to look up a value by its associated key. Example:
```go
var x map[string]int
```
*"x is a map of strings to ints."*

```go
name, ok := elements["Un"]
fmt.Println(name, ok)
```
Accessing an element of a map can return two values instead of just one. The first value is the result of the lookup, the second tells us whether or not the lookup was successful. In Go, we often see code like this:  
```go
if name, ok := elements["Un"]; ok {
	fmt.Println(name, ok)
}
```
Like we saw with arrays, there is also a shorter way to create maps:
```go
elements := map[string]string {
	"H": "Hydrogen",
	"He": "Helium",
	"Li": "Lithium",
	"Be": "Beryllium",
	"B": "Boron",
	"C": "Carbon",
	"N": "Nitrogen",
	"O": "Oxygen",
	"F": "Flourine",
	"Ne": "Neon",
}
```
