package main

import "fmt"

func main() {
	defer fmt.Println("Two")
	defer fmt.Println("One")
	defer fmt.Println("World")
	fmt.Println("Hello")
}
