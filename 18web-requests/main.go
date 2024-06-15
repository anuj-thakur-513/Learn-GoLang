package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.google.com"

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Web request")
	res, err := http.Get(url)
	checkNilError(err)

	fmt.Println(res)
	defer res.Body.Close() // caller's responsibility to close the connection

	dataBytes, err := io.ReadAll(res.Body)
	checkNilError(err)

	content := string(dataBytes)
	fmt.Println(content)
}
