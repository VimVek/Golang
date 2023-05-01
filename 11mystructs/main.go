package main

import "fmt"

func main() {
	// NO inheritance, No Super,parents or child
	fmt.Println("Structs in Golang")

	Vivek := User{"Vivek", "vivek.joshi", true, 25}
	fmt.Println(Vivek)
	fmt.Printf("VIveks details are: %+v\n", Vivek)
	fmt.Printf("Name is %v, Age is %v \n", Vivek.Name, Vivek.Age)
}

// User has capital U as it is acting like a class or template that we will be using
// and needs to be exported out, same for fields as it can be used by anybody
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
