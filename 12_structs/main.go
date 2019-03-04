package main

import "fmt"

//define struct
type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	fmt.Println("Hello World")

	//init person
	person1 := Person{firstName: "John", lastName: "Smith", age: 26}
	fmt.Println(person1)

	//alternative init
	person2 := Person{"Jane", "Doe", 26}
	fmt.Println(person2)

	//print field
	person2.age = 67
	fmt.Println(person2.age)
}
