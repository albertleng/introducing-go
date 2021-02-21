package main

import "fmt"

func main() {
	// var x string = "Hello, world"
	// fmt.Println(x)

	// var x string
	// x = "first"
	// fmt.Println(x)
	// x = "second"
	// fmt.Println(x)

	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}
