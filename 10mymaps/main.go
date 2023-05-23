package main

import "fmt"

func main() {
	fmt.Println("Maps in GOlang")
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["Py"] = "Python"
	languages["J"] = "Java"
	languages["GO"] = "G0Lang"

	fmt.Println("List of all Languages ", languages)
	fmt.Println("JS is shortform of ", languages["JS"])

	//deleting from Maps
	delete(languages, "J")
	fmt.Println(languages)

	//iterating in maps

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}
}
