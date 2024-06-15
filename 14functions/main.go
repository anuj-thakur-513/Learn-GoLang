package main

import (
	"fmt"
)

func greet(name string) {
	fmt.Println("Hello " + name)
}

func twoSum(val1 int, val2 int) int {
	return val1 + val2
}

func proSum(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func main() {
	fmt.Println("Welcome to functions in GoLang")
	greet("Anuj Thakur")
	fmt.Println(twoSum(1, 2))
	fmt.Println(proSum(1, 2, 3, 5, 6, 1234, 134))
}
