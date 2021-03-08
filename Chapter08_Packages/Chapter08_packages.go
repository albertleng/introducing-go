package main

import (
	"fmt"
	"os"
)

func main() {

	//strings
	//fmt.Println(strings.Contains("test", "es"))
	//
	//fmt.Println(strings.Count("test", "t"))
	//
	//fmt.Println(strings.HasPrefix("test", "te"))
	//
	//fmt.Println(strings.HasSuffix("test", "st"))
	//
	//fmt.Println(strings.Index("test", "e"))
	//
	//fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	//
	//fmt.Println(strings.Repeat("a", 5))
	//
	//fmt.Println(strings.Replace("aaaaa", "a", "b", 2))
	//
	//fmt.Println(strings.Split("a-b-c-d-e", "-"))
	//
	//fmt.Println(strings.ToLower("TEST"))
	//
	//fmt.Println(strings.ToUpper("test"))

	//os

	file, err := os.Open("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)





}