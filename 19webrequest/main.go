package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://pkg.go.dev/"

func main() {
	fmt.Println("Web Request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T \n", response)

	defer response.Body.Close() //it is important to closethe connection
	databyte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databyte) //convert to string as it will be in byte form
	fmt.Print("Data is: \n", content)
}
