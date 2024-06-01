package main

import "fmt"

func main() {
	fmt.Println("Let's study about pointers")

	var one int = 230
	var ptr *int = &one

	fmt.Println("address of one is:", ptr)
	fmt.Println("value of pointer is:", *ptr)
	fmt.Println("address of one is:", &ptr)
}
