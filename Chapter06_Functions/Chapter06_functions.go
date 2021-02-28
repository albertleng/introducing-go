package main

import "fmt"

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
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")

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