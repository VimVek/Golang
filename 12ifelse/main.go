package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	numb, _ := strconv.ParseInt(strings.TrimSpace(input), 0, 255)
	fmt.Println(numb)
	if numb > 10 {
		fmt.Printf("Number is greater than 10 and is %v \n", numb)
	} else {
		fmt.Println("what")
	}
}
