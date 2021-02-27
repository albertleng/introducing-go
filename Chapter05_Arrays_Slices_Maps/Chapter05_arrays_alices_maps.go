package main

import "fmt"

func main() {

	/*
		Arrays
	*/
	//var x [5]int
	//x[4] = 100
	//fmt.Println(x)
	//
	//var x [5]float64
	//x[0] = 98
	//x[1] = 93
	//x[2] = 77
	//x[3] = 82
	//x[4] = 83

	//x := [5]float64{ 98, 93, 77, 82, 83}
	//x := [5]float64{ 98,
	//	93,
	//	77,
	//	82,
	//	83,
	//}

	//var total float64 = 0.0
	//for i := 0; i < len(x); i++ {
	//	total += x[i]
	//}
	//fmt.Println(total/float64(len(x)))

	//var total float64 = 0.0
	//for _, value := range x {
	//	total += value
	//}
	//fmt.Println(total/float64(len(x)))

	/*
		Slices
	*/

	//append
	//slice1 := []int{1, 2, 3}
	//slice2 := append(slice1, 4, 5)
	//fmt.Println(slice1, slice2)

	//copy
	//slice1 := []int{1, 2, 3}
	//slice2 := make([]int, 2)
	//copy(slice2, slice1)
	//fmt.Println(slice1, slice2)

	/*
		Slices
	*/
	//x := make(map[string]int)
	//x["key"] = 10
	//fmt.Println(x["key"], " with length ", len(x))

	//elements := make(map[string]string)
	//elements["H"] = "Hydrogen"
	//elements["He"] = "Helium"
	//elements["Li"] = "Lithium"
	//elements["Be"] = "Beryllium"
	//elements["B"] = "Boron"
	//elements["C"] = "Carbon"
	//elements["N"] = "Nitrogen"
	//elements["O"] = "Oxygen"
	//elements["F"] = "Fluorine"
	//elements["Ne"] = "Neon"
	//
	//fmt.Println(elements["Li"], elements["Al"])

	/*
		Exercise 4
	*/
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	// My answer
	//smallest := math.MaxInt64
	//for i := 0; i < len(x); i++ {
	//	if x[i] < smallest {
	//		smallest = x[i]
	//	}
	//}
	//fmt.Println("smallest number is ", smallest)

	// Book's answer
	var min int
	for i, v := range x {
		if i == 0 || min > v {
			min = v
		}
	}
	fmt.Println(min)
}
