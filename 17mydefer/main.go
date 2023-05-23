package main

import "fmt"

// Defers run at the end and they run in order of LIFO
func main() {
	defer fmt.Println("At First")
	fmt.Println("Defers")
	defer fmt.Println("At Middle")
	myDefer()
	defer fmt.Println("At Last")
}

// Flow will be, Defers( defer iside it so 43210), myDefer, At Last, At Middle, At First
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("Value is: ", i)
	}
}
