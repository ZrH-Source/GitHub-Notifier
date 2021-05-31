package main

import (
	"fmt"
	"net/http"

	"github.com/zrh-source/GitHub-Notifier/adapter"
)

func main() {
	fmt.Printf("Starting server at port 80\n")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/github", adapter.GitHandler)
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World")
}
