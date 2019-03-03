package main

import "fmt"

func main() {
	//long method
	i := 1

	for i <= 10 {
		fmt.Println(i)
		//i = i + 1
		i++
	}

	//short method
	for counter := 1; counter <= 20; counter++ {
		fmt.Printf("Number %d\n", counter)
	}

	//fizzbuzz

	for counter := 1; counter <= 100; counter++ {
		if counter%15 == 0 {
			fmt.Printf("%d FizzBuzz\n", counter)
		} else if counter%3 == 0 {
			fmt.Printf("%d Fizz\n", counter)
		} else if counter%5 == 0 {
			fmt.Printf("%d Buzz\n", counter)
		} else {
			fmt.Printf("%d\n", counter)
		}
	}
}
