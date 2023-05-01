package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("SLices")
	var fruitlist = []string{}
	//var fruitlist = []string{"Tomato", "Pear"}
	fmt.Printf("type of fruitlist is %T\n", fruitlist)

	fruitlist = append(fruitlist, "Mango", "Banana", "Tomato", "Pear")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:])
	fmt.Println(fruitlist)

	highscores := make([]int, 4)
	highscores[0] = 123
	highscores[1] = 231
	highscores[2] = 321
	highscores[3] = 541
	//highscores[4] = 123 it will not work

	highscores = append(highscores, 4332, 2323, 331, 12)
	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	// remove value from slices based on index
	var courses = []string{"python", "java", "react", "javascript", "golang"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	// we add ... as append takes 1 argument and we are giving 2
	fmt.Println(courses)
}
