package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Learn slices in golang")

	var fruitList = []string{"apple", "tomato", "mango"}

	fmt.Printf("type of fruitlist is: %T\n", fruitList)

	fruitList = append(fruitList, "peach", "banana")
	fmt.Println("fruitlist is:", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	// memory management keywords
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	highScores = append(highScores, 555, 666)

	fmt.Printf("type of highScores is: %T\n", highScores)
	fmt.Println("highScores:", highScores)

	sort.Ints(highScores)
	fmt.Println("highScores:", highScores)

	// remove value from slices based on index
	var courses = []string{"reactjs", "javascript", "golang", "python"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
