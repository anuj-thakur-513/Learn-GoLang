package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func encodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "Learncodeonline.in", "test123", []string{"web-dev", "frontend"}},
		{"NodeJS Bootcamp", 399, "Learncodeonline.in", "test123", []string{"web-dev", "backend"}},
		{"Go Bootcamp", 399, "Learncodeonline.in", "test123", nil},
	}

	// package the above data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))
}

func main() {
	fmt.Println("Let's learn to create JSON in golang")
	encodeJson()
}
