package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in GoLang")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is: ", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2 you can move 2 spots")
	case 3:
		fmt.Println("Dice value is 3 you can move 3 spots")
		fallthrough //if 3 comes we want to run 4 also
	case 4:
		fmt.Println("Dice value is 4 you can move 4 spots")
	case 5:
		fmt.Println("Dice value is 5 you can move 5 spots")
		fallthrough
		//if 5 comes we want to run 6 also
	case 6:
		fmt.Println("Dice value is 6 you can move 6 spots and roll the dice again")
	default:
		fmt.Println("Unkwon error!! what was that")
	}
}
