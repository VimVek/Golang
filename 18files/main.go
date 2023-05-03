package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Golang")
	content := "Needs to go in File"
	file, err := os.Create("./mygofile.text")
	CheckNilError(err)
	//if err != nil {  isted of writing this we call the above function
	//	panic(err)
	//The panic built-in function stops normal execution of the current goroutine.
	//}
	length, err := io.WriteString(file, content)
	CheckNilError(err)
	fmt.Println("Length is: ", length)
	defer file.Close()

	ReadFile("./mygofile.text")
}

func ReadFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	CheckNilError(err)
	fmt.Println("Text data inside file is: ", string(databyte)) // databyte this will show data like in byte
}

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}
