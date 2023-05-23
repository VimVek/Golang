package main

import (
	"fmt"
	"time"
)

// GOOS = "windows" "go build  this command build our code for windows
func main() {
	fmt.Println("Time of GOLANG ")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	// this time is standard for formatting, same with day and time

	createdDate := time.Date(2020, time.March, 28, 15, 30, 33, 0, time.Now().UTC().Location())
	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
