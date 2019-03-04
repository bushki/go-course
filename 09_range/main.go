package main

import "fmt"

func main() {

	ids := []int{2, 5, 0, 3, 10}

	//loop through array of ids
	for index, id := range ids {
		fmt.Printf("index: %d - ID: %d\n", index, id)
	}

	//loop through array of ids not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//add all elements
	sum := 0
	for id := range ids {
		sum += id
	}

	fmt.Printf("Sum: %d\n", sum)

	//range with map
	presidents := map[int]string{1: "George Washington", 2: "John Adams", 3: "Thomas Jefferson"}
	fmt.Println("*** Presidents ***")
	for key, value := range presidents {
		fmt.Printf("Key %d: Value: %s\n", key, value)
	}

}
