package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web verb video")
	//PerformGetRequest()
	//PerformPostJsontRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	//making web request
	response, err := http.Get(myurl) //this will give response or err

	CheckError(err)
	defer response.Body.Close() //Imp must be done everytime

	fmt.Println("Status Code", response.StatusCode)
	fmt.Println("Content Length is: ", response.ContentLength)

	var responseString strings.Builder

	content, err := ioutil.ReadAll(response.Body) //to read content
	CheckError(err)

	byteCount, err := responseString.Write(content)
	CheckError(err)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String()) //covert data in form of string
	// data is inside responseString

	fmt.Println(string(content)) // this does same work

}
func PerformPostJsontRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload
	//Newreader allow us to create new reading format
	requestBody := strings.NewReader(`
	{
		"coursename": "Lets GO with GoLang",
		"price" : 0,
		"platform": "vivekjoshi.com"
	}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	CheckError(err)

	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	CheckError(err)

	fmt.Println("Content: ", string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata
	data := url.Values{}
	data.Add("Firstname", "Vivek")
	data.Add("Lastname", "Joshi")
	data.Add("Email", "Vivek@Gmail.com")

	response, err := http.PostForm(myurl, data)
	CheckError(err)

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
