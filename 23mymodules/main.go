package main

//go mod vendor (bring modules from package and not from web)
//go run -mod-vendor main.go (this will bring everything from vendor)
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("MOD in GoLong")
	Greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", ServeHome).Methods("GET")

	http.ListenAndServe(":4000", r) //running server with port name and router
	log.Fatal(http.ListenAndServe(":4000", r))
}

func Greeter() {
	fmt.Println("Hello Mod users")
}

// request is someone sending us request and w is sending back the response(responding)
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to Golang </h1>"))
}
