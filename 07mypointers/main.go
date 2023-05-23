package main

import "fmt"

//By pointer semantic we are giving safety of immutablity for more efficiency

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

	i, j := 420, 69
	p := &i
	fmt.Println("This is i: ", i)
	fmt.Println("This is j: ", j)
	fmt.Println("This is p pointing to address ofi: ", p)
	fmt.Println("This is p pointing to value at address of i: ", *p)

	a := 4
	squareAdd(&a)
	squareVal(a)
	squareAdd(&a) // paasing addrress of a here
	//here insted of copying a we copying address of a and assigning it as pointer p

	//Now if we make any changes to *p, it will change in real address block
	*p = 4200
	fmt.Println("This is i: ", i)
	fmt.Println("This is p pointing to value at address of i: ", *p)

	fmt.Println(initPerson())
}

// this function will make a copy of a and make changes on it and print it, it will not affect real value of a
func squareVal(v int) {
	v *= v
	fmt.Println(&v, v)
}

// this will use address and change the value at that specific address
func squareAdd(p *int) { //*int is
	*p *= *p //value at p
	fmt.Println(p, *p)
}

// defining a structure  person with name and age
type Person struct {
	name string
	age  int
}

// function that declare a person as m and initialize it
func initPerson() *Person {
	m := Person{name: "noname", age: 25}
	return &m
}
