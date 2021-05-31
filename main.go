package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Starting server at port 80\n")
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":80", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World")
}
