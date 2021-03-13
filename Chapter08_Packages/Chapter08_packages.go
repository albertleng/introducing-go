package main

import (
	"fmt"
	"hash/crc32"
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

	//file, err := os.Open("test.txt")
	//if err != nil {
	//	// handle the error here
	//	return
	//}
	//defer file.Close()
	//
	//// get the file size
	//stat, err := file.Stat()
	//if err != nil {
	//	return
	//}
	//
	//// read the file
	//bs := make([]byte, stat.Size())
	//_, err = file.Read(bs)
	//if err != nil {
	//	return
	//}
	//
	//str := string(bs)
	//fmt.Println(str)
	
	
	//bs, err := ioutil.ReadFile("test.txt")
	//if err != nil {
	//	return
	//}
	//str := string(bs)
	//fmt.Println(str)


	//file, err := os.Create("test.txt")
	//if err != nil {
	//	return
	//}
	//defer file.Close()
	//file.WriteString("test")


	//dir, err := os.Open(".")
	//if err != nil {
	//	return
	//}
	//defer dir.Close()
	//
	//fileInfos, err := dir.Readdir(-1)
	//if err != nil {
	//	return
	//}
	//for _, fi := range fileInfos {
	//	fmt.Println(fi.Name())
	//}


	//filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
	//	fmt.Println(path, info.Size(), err)
	//	return nil
	//})


	//err := errors.New("error message")
	//fmt.Println(err)


//	Containers and Sort

//	LIST
//	var x list.List
//	x.PushBack(1)
//	x.PushBack(2)
//	x.PushBack(3)
//
//	for e:= x.Front(); e != nil; e=e.Next() {
//		fmt.Println(e.Value.(int))
//	}

	//SORT
	//kids := []Person {
	//	{"Jill", 9},
	//	{"Jack", 10},
	//}
	//sort.Sort(ByName(kids))
	//fmt.Println(kids)
	//
	//sort.Sort(ByAge(kids))
	//fmt.Println(kids)


	//Hashes and Cryptography
//	Non-cryptography

	// create a hasher
	h := crc32.NewIEEE()
	// write our data to it
	h.Write([]byte("test"))
	// calculate the crc32 checksum
	v := h.Sum32()
	fmt.Println(v)

}
type Person struct {
	Name string
	Age int
}

type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type ByAge []Person
func (this ByAge) Len() int {
	return len(this)
}
func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}
func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}