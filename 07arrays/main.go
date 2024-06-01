package main

import "fmt"

func main() {
	fmt.Println("Learn arrays in golang")

	var stringArr [4]string
	stringArr[0] = "1"
	stringArr[1] = "2"
	stringArr[3] = "4"

	fmt.Println("array is:", stringArr, "\nlen of array is:", len(stringArr))

	var intArr = [3]int{0, 1, 2}
	fmt.Println(intArr)
}
