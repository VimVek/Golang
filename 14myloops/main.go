package main

import "fmt"

func main() {
	fmt.Println("All about Loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "Saturday"}
	fmt.Println(days)

	//for d := 0; d < len(days); d++ {
	//	fmt.Println(days[d])
	//}

	for i := range days { //i return index
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("index is %v and day is %v \n", index, day)
	}

	randomValue := 1
	for randomValue < 10 {
		if randomValue == 7 {
			randomValue++
			continue
		}
		if randomValue == 9 {
			goto vjo
		}
		fmt.Println(randomValue)
		randomValue++
	}
vjo: //label
	fmt.Println("Welcome to VIVEK JOSHII working")
}
