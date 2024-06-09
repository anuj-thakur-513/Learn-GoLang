package main

import "fmt"

func main() {
	fmt.Println("Understanding loops in golang")
	// days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	// fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {
	// 	fmt.Printf("value is: %v\n", day)
	// }

	rogueValue := 1

	for rogueValue < 10 {
		if rogueValue == 5 {
			rogueValue++
			continue
		}

		fmt.Println(rogueValue)
		rogueValue += 1
	}
}
