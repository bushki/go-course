package main

import "fmt"

func main() {
	//MAIN TYPES
	//string, bool, int, int8, in16, int32, int64
	//uint, uint8, uint16....
	//byte - alias for uint8
	//rune - alias for int32
	//float32, float64
	//complex64, complex128 (large nums)

	//Using var
	var name string = "John"
	var age = 8
	fmt.Println(name)

	//without var (inferred)
	var last = "Smith"
	fmt.Println(last, age)

	//get type
	//https://golang.org/pkg/fmt/
	fmt.Printf("%T\n", name)

	//specify type
	var age2 int32 = 25
	fmt.Printf("%T\n", age2)

	//const
	const isMale = true
	fmt.Println(name, age2, isMale)

	//short-hand variable declaration
	myname := "Luis"
	fmt.Println(myname)

	//float
	size := 1.3
	fmt.Println(size)

	//another shorthand
	country, state := "US", "CA"
	fmt.Println(country, state)
}
