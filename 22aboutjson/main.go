package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string `json:"FullName"` // `jason:"anything"` is for alias
	Price    int
	Platform string   `json:"Website"`
	Password string   `json:"-"`              //"-" meansdont want to be reflected
	Tags     []string `json:"Tags,omitempty"` // omitempty means dont print field if empty
}

func main() {
	fmt.Println("JSON")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []Course{
		{"Golang", 1599, "udemy.com", "abc123", []string{"web-dev", "language", "backend"}},
		{"Python", 999, "udemy.com", "bca123", nil},
		{"Java", 2599, "udemy.com", "deabc123", []string{"OOPs", "frontend", "backend"}},
	}

	//package this data asJSON data

	//finalJson, err := json.Marshal(courses) //Marshal  IMplementation of json
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	//"" empty as if we put anything than it will show in each line ,indented on basis of \t
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJson)
}
func DecodeJson() {
	//data will be in form of slice of byte
	jsonData := []byte(`
	{
		"FullName": "Golang",  
		"Price": 1599,
		"Website": "udemy.com",
		"Tags": ["web-dev","language","backend"]
	}
	`)
	//since Fullname and website is alias, json will be automatically converted to names defined in structs

	var myCourse Course
	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("JSON was Valid")
		json.Unmarshal(jsonData, &myCourse)
		fmt.Printf("%#v \n", myCourse)
	} else {
		fmt.Println("Json was not Valid")
	}
	// some caseswhere we just want toadd data to key value

	var myOnlineData map[string]interface{}
	//since we dont know what type of data will come from web, it can be string, array or whatever so we make it interface
	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Printf("%#v \n", myOnlineData)
	for k, v := range myOnlineData {
		fmt.Printf("Key is  %v and Value is %v and type is: %T \n", k, v, v)
	}
}
