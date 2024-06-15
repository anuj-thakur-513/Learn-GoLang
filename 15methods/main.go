package main

import "fmt"

func main() {
	var anuj User = User{"anuj thakur", "anujthakur0103@gmail.com", true, 21}
	temp := User{"temp user", "temp@temp.com", false, 10}
	fmt.Println(anuj, temp)

	fmt.Printf("anuj details are: %+v\ntemp details are: %+v\n", anuj, temp)
	anuj.GetStatus()
	anuj.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@test.com"
	fmt.Println("Email of this user is:", u.Email)
}
