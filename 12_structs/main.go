package main

import (
	"fmt"
	"strconv"
)

//define struct
type Person struct {
	firstName string
	lastName  string
	age       int
}

//greeting method (value receiver)
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

func main() {
	fmt.Println("Hello World")

	//init person
	person1 := Person{firstName: "John", lastName: "Smith", age: 26}
	fmt.Println(person1)

	//alternative init
	person2 := Person{"Jane", "Doe", 26}
	fmt.Println(person2)

	//change and print field
	person2.age = 67
	fmt.Println(person2.age)

	//call value method
	fmt.Println("*** call value method ***")
	fmt.Println(person2.greet())
}
