package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s! Also The GoLang Server is Working!", r.URL.Path[1:]) //slice the '/' from the url path
} //just print on the file (via url text)

func main() {
	http.HandleFunc("/", handler) //handles what next on said web page
	http.ListenAndServe(":8080", nil)
}
