package main

import "fmt"

// constant
const LoginToken string = "hiIamConstant" //capital first character i.e, LoginToken states that this is Public variable

func main() {
	var username string = "Vivek" // we have to use the variable is we make it
	fmt.Println("username")       //fp short form to for fmt.Println
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true // type boolean
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255 // untyped integer 0-255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.12341341313 // untyped integer 0-255
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// Default value and some aliases
	var anotherVar int
	var anotherStn string
	anotherVar = 123
	anotherStn = "Hello I am String"
	fmt.Println(anotherVar)
	fmt.Println(anotherStn)
	fmt.Printf("Variable is of type: %T \n", anotherVar)
	fmt.Printf("Variable is of type: %T \n", anotherStn)

	anotherVar = 234
	fmt.Println(anotherVar)

	//implicit type (it auto assign type according to value but we cant change it later on)
	var website = "hieloo.in"
	fmt.Println(website)

	website = "changing vivek.in"
	fmt.Println(website)
	//cant be change to website = 3 as it is string

	// no var style (total ignore keyword var and still declare variable)
	numberOfUser := 300000 // := is walrus operator allowed inside method only, not allowed outside
	fmt.Println(numberOfUser)

	//const
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
