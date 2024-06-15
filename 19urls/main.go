package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "http://www.lco.dev:3000/learn?coursename=backend&paymentid=gheawrjksf12123"

func main() {
	fmt.Println("Handling URLs in golang")

	fmt.Println(myUrl)
	// parsing
	res, _ := url.Parse(myUrl)

	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	fmt.Println(res.Port())
	fmt.Println(res.RawQuery)

	qparams := res.Query()

	fmt.Println(qparams) // stored in map

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "coursename=backend&paymentid=testing",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
