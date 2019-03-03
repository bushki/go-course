package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println("Hello World")

	x := 5
	y := 10

	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	//else if
	color := "     red      "

	if strings.Trim(color, " ") == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is NOT red or blue")
	}

	//switch

	switch strings.Trim(color, " ") {
	case "red":
		fmt.Println("switch: Color is red")
	case "blue":
		fmt.Println("switch: Color is red")
	default:
		fmt.Println("switch: Color is unknown")
	}
}
