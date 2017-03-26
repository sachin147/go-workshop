package main

import (
	"net/http"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func method(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Method)
}

func path(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL)
}

func query(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Form)
}

func headers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Header)
}

// TODO: method handler to print http method
// TODO: path handler to print URL
// TODO: query handler to print query parameters
// TODO: headers handler to print http headers


func main() {
	fmt.Println("hello world")

	http.HandleFunc("/hello", hello)
    http.HandleFunc("/method", method)
    http.HandleFunc("/path/", path)
    http.HandleFunc("/query", query)
    http.HandleFunc("/headers", headers)
	// TODO: /method
	// TODO: /path
	// TODO: /query
	// TODO: /headers
	http.ListenAndServe("127.0.0.1:8000", nil)
}
