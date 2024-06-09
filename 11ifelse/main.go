package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	var age int = 23
	var result string
	if age < 18 {
		result = "Underage"
	} else {
		result = "Adult"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}
}
