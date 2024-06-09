package main

import "fmt"

func main() {
	fmt.Println("Maps in go lang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["CPP"] = "C++"
	languages["go"] = "GoLang"

	fmt.Println(languages)
	fmt.Println("JS shorts for:", languages["JS"])

	delete(languages, "CPP")
	fmt.Println(languages)
}
