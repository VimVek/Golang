package main

import "fmt"

func main() {
	fmt.Println("ALl about functions")
	greeter()
	greetertwo()

	result, mymessage := adder(3, 5)
	fmt.Println(result)
	fmt.Println("Message is ", mymessage)
	proResult := proAdder(2, 3, 4, 5, 6)
	fmt.Println(proResult)
}

func greetertwo() {
	fmt.Println("Another ")
}

//	func () {
//		fmt.Println("Another ")
//	}()   // these function can be immediatly invoked/  annoynamous function

func greeter() {
	fmt.Println("Hello from GoLang")
}

func adder(valOne int, valTwo int) (int, string) { //function signatures, what kind of value to return so adding int after()
	return valOne + valTwo, "THis returns 2 value"
}

func proAdder(values ...int) int { //all values of type integer and we will return integer
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
