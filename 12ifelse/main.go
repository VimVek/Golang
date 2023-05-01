package main

import "fmt"

func main() {
	fmt.Println("If ELse")
	loginCount := 10
	var res string
	if loginCount < 10 {
		res = "Regular"
	} else if loginCount > 10 {
		res = "Premium"
	} else {
		res = "Exactly 10"
	}
	fmt.Println(res)

	if num := 3; num < 10 {
		fmt.Printf("NUmber is less than 10 and is %v \n", num)
	}
}
