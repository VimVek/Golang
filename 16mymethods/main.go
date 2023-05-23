package main

import "fmt"

func main() {
	fmt.Println("Methods") //functions when go into classes are called methods
	//we dont have classes in go and have structs

	Vivek := User{"Vivek", "vivek.joshi", true, 25}
	fmt.Println(Vivek)
	fmt.Printf("VIveks details are: %+v\n", Vivek)
	fmt.Printf("Name is %v, Age is %v \n", Vivek.Name, Vivek.Age)
	Vivek.GetStatus()
	Vivek.NewMail()
	fmt.Printf("Name is %v, Age is %v \n", Vivek.Name, Vivek.Email)

}

// User has capital U as it is acting like a class or template that we will be using
// and needs to be exported out, same for fields as it can be used by anybody
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	//oneAge int, if we write like this than this property will not be exportable
}

// this is how we create method for user
func (u User) GetStatus() {
	fmt.Println("Is User active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@gmail.com"
	// when we pass thing along multiple functions, usually copy of object is pass along
	fmt.Println("Email is changed to: ", u.Email)
}
