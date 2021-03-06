package main

import (
	"fmt"
	"math"
)

func main() {
	//xs := []float64{98, 63, 77, 82, 83}
	//fmt.Println(average(xs))

	/*
	 Variadic Functions
	 */
	//fmt.Println(add(1, 2, 3))
	//xs := []int{1, 2, 3}
	//fmt.Println(add(xs...))
	
	/*
	Closure
	 */
	
	//add := func(x, y int) int {
	//	return x + y
	//}
	//fmt.Println(add(1, 1))

	//x := 0
	//increment := func() int {
	//	x++
	//	return x
	//}
	//fmt.Println(increment())
	//fmt.Println(increment())

	//nextEven := makeEvenGenerator()
	//fmt.Println(nextEven())
	//fmt.Println(nextEven())
	//fmt.Println(nextEven())

	/*
	Recursion
	 */
	//fmt.Println(factorial(2))

	/*
	defer, panic, and recover
	 */
	//defer second()
	//first()

	/*
	panic and recover
	 */
	//defer func() {
	//	str := recover()
	//	fmt.Println(str)
	//}()
	//panic("PANIC")

	/*
	new
	 */
	//xPtr := new(int)
	//one(xPtr)
	//fmt.Println(*xPtr)

/*
	Exercise 2
 */
	//quotient, isEven := halves(1)
	//fmt.Println(quotient, isEven)
	//quotient, isEven = halves(2)
	//fmt.Println(quotient, isEven)

	/*
	Exercise 3
	 */
	//nums := []int{1, 4, 5, 9, 2, 33, 100, 92, 12}
	//fmt.Println(greatestNum(nums...))
	//fmt.Println(greatestNum(9, 1, 2, 4, 5, 7, 11))


	/*
	Exercise 4
	 */
//nextOdd := makeOddGenerator()
//fmt.Println(nextOdd())
//fmt.Println(nextOdd())
//fmt.Println(nextOdd())

	/*
	Exercise 5
	 */
//fmt.Println(fibo(0))
//fmt.Println(fibo(1))
//fmt.Println(fibo(2))
//fmt.Println(fibo(3))
//fmt.Println(fibo(4))
//fmt.Println(fibo(5))

	/*
	Exercise 11
	 */
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)
}

/*
Exercise 11
 */

func swap(x, y *int) {
	*x, *y = *y, *x
}

/*
Exercise 5
 */
func fibo(num uint) uint {
	if num == 0 || num == 1 {
		return num
	}
	return fibo(num-1) + fibo(num-2)
}
/*
Exercise 4
 */

//func makeOddGenerator() func() uint {
//	i := uint(1)
//	return func() (ret uint) {
//		ret = i
//		i += 2
//		return
//	}
//}

/*
Exercise 3
 */
func greatestNum(args ...int) int {
	greatest := math.MinInt32

	for _, num := range args {
		if num > greatest {
		    greatest = num
		}
	}

	return greatest

}

/*
Exercise 2
 */

func halves(num int) (int, bool) {
	isEven := num % 2 == 0

	return num/2, isEven
}

func one(xPtr *int) {
	*xPtr = 1
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
 }
func average(xs []float64) float64 {
  total := 0.0
  for _, v := range xs {
  	total += v
  }
  return total / float64(len(xs))
}