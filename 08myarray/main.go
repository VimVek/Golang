package main

import "fmt"

func main() {
	fmt.Println("Array in Golang")
	var fruitlist [4]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Banana"
	fruitlist[2] = "Tomato"
	fruitlist[3] = "PineAPple"
	fmt.Println("Fruit list is :", fruitlist)

	var veglist = [3]string{"Potato", "onion", "Beans"}
	fmt.Println("Veglist is ", veglist)
	fmt.Println("len of veglist is", len(veglist))

}
