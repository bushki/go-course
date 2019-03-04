package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {
	fmt.Println("Hello World")
	mySum := adder()
	for counter := 1; counter < 10; counter++ {
		fmt.Println(mySum(counter))
	}
}
