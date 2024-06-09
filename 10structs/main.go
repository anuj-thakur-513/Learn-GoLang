package main

import "fmt"

func main() {
	// no inheritance in golang; no super or parent

	var anuj User = User{"anuj thakur", "anujthakur0103@gmail.com", true, 21}
	temp := User{"temp user", "temp@temp.com", false, 10}
	fmt.Println(anuj, temp)

	fmt.Printf("anuj details are: %+v\ntemp details are: %+v\n", anuj, temp)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
