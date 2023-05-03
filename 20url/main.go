package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://www.udemy.com/course/data-structures-and-algorithms-java/learn/lecture/28748718#notes"

func main() {
	fmt.Println("URLS handelling in GoLang")
	fmt.Println(myurl)

	//parsing URL
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)   //it will return https
	fmt.Println(result.Host)     //it will return www.udemy.com
	fmt.Println(result.Path)     //it will return /course/data-structures-and-algorithms-java/learn/lecture/28748718
	fmt.Println(result.Port())   //it will return nothing here (port number)
	fmt.Println(result.RawQuery) //it will return nothing here( but it will return params)

	qparams := result.Query()
	fmt.Printf("The quert Params are of type %T \n", qparams) // url.VAlues are basically key value pairs
	// fmt.Println(qparams["anything"]) this will return value

	//we can iterate over these param
	for _, val := range qparams {
		fmt.Println("Param is: \n", val)
	}
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "www.udemy.com",
		Path:   "/course/data-structures-and-algorithms-java/learn/lecture/28748718",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
