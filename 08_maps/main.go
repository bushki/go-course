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
}
