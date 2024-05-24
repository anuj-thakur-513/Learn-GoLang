package main

import "fmt"

func main() {
	var username string = "test user"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)


	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)


	// implicit type
	var website = "github.com"
	numberOfUsers := 300000
	fmt.Println(website, numberOfUsers)
}