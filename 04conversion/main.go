package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hi")
	fmt.Println("Please choose any number betwwen 0-5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating: ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //input is of type string so we will use a package
	//strconv.ParseFloat(input, 64) will give error as input contains value/n ex= 4/n as we pressing enter
	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
