package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func checkErrorNil(err error) {
	if err != nil {
		panic(err)
	}
}

func performGetRequest() {
	const myUrl string = "http://localhost:8000/get"
	res, err := http.Get(myUrl)

	checkErrorNil(err)
	defer res.Body.Close()

	// fmt.Println("response:", res)
	// fmt.Println("Status code:", res.StatusCode)

	content, _ := io.ReadAll(res.Body)
	// fmt.Println(string(content))
	var resString strings.Builder
	byteCount, _ := resString.Write(content)
	fmt.Println("byte count is:", byteCount)
	fmt.Println(resString.String())
}

func main() {
	fmt.Println("welcome to web requests")
	performGetRequest()
}
