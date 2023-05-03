package main

import "fmt"

func main() {
	fmt.Println("Pointerss")
	var ptr *int //creating a pointer
	//*int is used to declare variable ptr to be a pointer to an int
	fmt.Println("Value of Pointer is: ", ptr) // default value of pointer is nil

	myNumber := 28
	var ptr1 = &myNumber //creating pointer which is also refrencing to some memory
	//ptr1 is assigned the value that is the address of where variable b is stored
	fmt.Println("Value of actual pointer is ", ptr1)  //actual memory address
	fmt.Println("Value of actual pointer is ", *ptr1) //to see what inside the pointer
	*ptr1 = *ptr1 + 1                                 //this will change actual value as it is pointing to address
	fmt.Println("Value now is", myNumber)
}
