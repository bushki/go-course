package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	//output: 5 0xc00004c080
	fmt.Printf("%T %T \n", a, b)

	//read value from pointer
	fmt.Println("Value from pointer *")
	fmt.Println(*b)
	fmt.Println(*&a)

	//change value using pointer
	fmt.Println("Change value from pointer")
	*b = 17
	fmt.Println(a)

	/*
		http://goinbigdata.com/golang-pass-by-pointer-vs-pass-by-value/
			Strictly speaking, there is only one way to pass parameters in Go - by value.
			Every time a variable is passed as parameter, a new copy of the variable is
			created and passed to called function or method. The copy is allocated at a different memory address.
	*/

}
