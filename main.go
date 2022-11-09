package main

import (
	"fmt"
	"net/http"
	"log"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func HandleRequest() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Server started on: http://localhost:8080")
	HandleRequest()
}