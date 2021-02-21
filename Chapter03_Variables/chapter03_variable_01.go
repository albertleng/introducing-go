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

	// fmt.Println("Enter a number: ")
	// var input float64
	// fmt.Scanf("%f", &input)

	// output := input * 2
	// fmt.Println(output)

	// My Solution for exercise 5
	// fmt.Println("Enter a temperature in Fahrenheit")
	// var inputTempInFah float64
	// fmt.Scanf("%f", &inputTempInFah)
	// temperatureInCelsius := (inputTempInFah - 32) * 5 / 9
	// fmt.Println(strconv.FormatFloat(inputTempInFah, 'f', -1, 64) + " Fahrenheit is " + fmt.Sprintf("%f", temperatureInCelsius) + " Celcius")

	// My solution for Exercise 6
	fmt.Println("Enter a length in feet")
	var inputLengthInFeet float64
	fmt.Scanf("%f", &inputLengthInFeet)
	inputLengthInMeters := inputLengthInFeet * 0.3048
	fmt.Println(fmt.Sprintf("%f", inputLengthInFeet) + " feet is " + fmt.Sprintf("%f", inputLengthInMeters) + " meters")
}
