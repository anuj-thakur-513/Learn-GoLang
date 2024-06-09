package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Swith Case in Golang")

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("move 1 spot")
	case 2:
		fmt.Println("move 2 spots")
	case 3:
		fmt.Println("move 3 spots")
	case 4:
		fmt.Println("move 4 spots")
	case 5:
		fmt.Println("move 5 spots")
	case 6:
		fmt.Println("move 6 spots and roll the dice again")
	}
}
