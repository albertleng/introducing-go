package main

import "fmt"

func main() {
	//i := 1
	//for i <= 10 {
	//	fmt.Println(i)
	//	i = i + 1
	//}

	//for i := 1; i <= 10; i++ {
	//	if i % 2 == 0 {
	//		fmt.Println(i, "even")
	//	} else {
	//		fmt.Println(i, "odd")
	//	}
	//}

	//for i := 1; i <= 10; i++ {
	//	switch i {
	//	case 0: fmt.Println("zero")
	//	case 1: fmt.Println("one")
	//	case 2: fmt.Println("two")
	//	case 3: fmt.Println("three")
	//	case 4: fmt.Println("four")
	//	case 5: fmt.Println("five")
	//	case 6: fmt.Println("six")
	//	case 7: fmt.Println("seven")
	//	case 8: fmt.Println("eight")
	//	case 9: fmt.Println("nine")
	//	case 10: fmt.Println("ten")
	//	default: fmt.Println("Unknown number")
	//	}
	//}

	// Exercise 2.
	//for i := 1; i <= 100; i++ {
	//	if i % 3 == 0 {
	//		fmt.Println(i)
	//	}
	//}

	// Exercise 3.
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
