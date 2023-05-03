package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome user input" //type declaration not imp in walrus operator
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin) // this line says = to read(from standard input/out)
	fmt.Println("Enter Rating of this code: ")

	// comma ok syntax// error ok sybntax (we dont hai try catch here)
	// either gonna get input or error i.e input, _  (in place of _ we can say err or anything)
	input, _ := reader.ReadString('\n') //read as long as\n is there
	fmt.Println("Thanks for Rating, ", input)
	fmt.Printf("Type of the rating: %T", input)
}
