//maps aka dictionary, key value pari

package main

import "fmt"

func main() {

	//init string/string dictionary
	emails := make(map[string]string)

	//assign
	emails["Bob"] = "bob@aol.com"
	emails["Jim"] = "Jim@aol.com"
	emails["Sharon"] = "sharon@aol.com"

	//print all
	fmt.Println(emails)
	fmt.Printf("length: %d\n", len(emails))

	//find element using key, print value
	fmt.Println(emails["Jim"])

	//delete from map
	fmt.Println("*** Delete ***")
	fmt.Println(emails)
	delete(emails, "Jim")
	fmt.Println(emails)

	//declare and assign kv
	fmt.Println("*** Presidents ***")
	presidents := map[int]string{1: "George Washington", 2: "John Adams", 3: "Thomas Jefferson"}
	presidents[4] = "James Madison"
	fmt.Println(presidents)
	fmt.Println(presidents[2])

}
