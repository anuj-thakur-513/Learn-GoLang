package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time package in Go")
	currentTime := time.Now()
	fmt.Println("Current date time:", currentTime.Format("01-02-2006 15:04:05 Monday"))


	createdDate := time.Date(2003, time.January, 5, 7, 30, 0, 0, time.UTC)
	fmt.Println("custom date:", createdDate.Format("01-02-2006 15:04:05 Monday"))
}