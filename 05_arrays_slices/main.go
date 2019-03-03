package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	//Arrays (fixed)
	var fruitArr [2]string

	//Assig values
	fruitArr[0] = "apple"
	fruitArr[1] = "orange"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	//Declare and assign
	myArr := [2]string{"banana", "grape"}
	fmt.Println(myArr)

	//Array Slice (size not fixed)
	arrSlice := []string{"Sliced banana", "grape", "apple", "orange"}
	fmt.Println(arrSlice)

	//len
	fmt.Println(len(arrSlice))

	//range
	fmt.Println(arrSlice[1:3]) //start at 1, do not include 3

}
