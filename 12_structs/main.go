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
	city      string
}

//greeting method (value receiver)
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//increase age
func (p *Person) increaseAge() {
	p.age++
}

//update city
func (p *Person) updateCity(newCity string) {
	p.city = newCity
}

func main() {
	fmt.Println("Hello World")

	//init person
	person1 := Person{firstName: "John", lastName: "Smith", age: 26, city: "LA"}
	fmt.Println(person1)

	//alternative init
	person2 := Person{"Jane", "Doe", 26, "NYC"}
	fmt.Println(person2)

	//change and print field
	person2.age = 67
	fmt.Println(person2.age)

	//call value method
	fmt.Println("*** call value method ***")
	fmt.Println(person2.greet())

	//using pointer method
	fmt.Println("*** call pointer method ***")
	fmt.Println(person2)
	person2.increaseAge()
	fmt.Println(person2)

	//update city using pointer method
	fmt.Println("*** call pointer method ***")
	fmt.Println(person2)
	person2.updateCity("Las Vegas")
	fmt.Println(person2)

}
