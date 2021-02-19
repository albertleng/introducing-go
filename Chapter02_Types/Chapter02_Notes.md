Go is a statically typed programming language. This means that variables always have a specific type and that type cannot change.  

## Numbers
### Integers
Integers - like their mathematical counterpart - are numbers without a decimal components.


Go's integer types are `uint8`, `uint16`, `uint32`, `uint64`, `int8`, `int16`, `int32` and `int64`. There are no alias types: `byte` (which is the same as `uint8`) and `rune` (which is the same as `int32`). There are also three machine dependent integer types: `uint`, `int` and `uintpr`.  

Generally, if you are working with integers, you should just use the `int` type.  


### Floating-Point Numbers  
Floating-point numbers are numbers that contain a decimal component (i.e., real numbers).  

- Floating-point numbers are inexact.
- Like integers, floating-point numbers have a certain size (32 bit or 64 bit).
- In addition to numbers, there are several other values that can be represented: "not a number" (NaN, for things like 0/0) and positive and negative infinity.


Go has two floating-point types: `float32` and `float64` (also often referred to as single precision and double precision, respectively). It also has two additional types for representing complex numbers (numbers with imaginary parts): `complex64` and `complex128`. Generally, we should stick with `float64` when working with floating-point numbers.  


The small set of basic operations, `+`, `-`, `*`, `/` and `%`, along with some of the helper functions available in the `math` package, is sufficient to create surprisingly sophisticated programs.  


### Strings
A string is a sequence of characters with a definite length used to represent text.  

String literals can be created using double quotes "Hello, World" or `` `Hello, World`.`` The difference between these is that double-quoted strings cannot contain newlines and they allow special escape sequences. For example, \n gets replaced with a newline and \t gets replaced with a tab character.  

Some common operations on strings:
  
  `len("Hello World")`
  `"Hello, World"[1]`
  `"Hello, " + "World"`


### Booleans
A boolean value is a special 1-bit integer type used to represent true and false (or on and off). Three logical operations are used with boolean values:  

&& and  
|| or  
! not  