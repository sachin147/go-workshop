package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"html/template"
	"net/url"
	"log"
	"time"
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
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, r.Form.Get("x"))
}

func headers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(r.Header)
}

func index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ctx := struct {
		Path *url.URL
		Method string
		Query url.Values
		Header http.Header
		Raw template.HTML
		Now time.Time
	}{
		Path: r.URL,
		Method: r.Method,
		Query: r.Form,
		Header: r.Header,
		Raw: template.HTML("<b>bold text</b>"),
		Now: time.Now(),
	}
	// TODO: Use base.html and index.html
	template.Must(
		template.ParseFiles("templates/index.html", "templates/base.html"),
	).Execute(w, ctx)
}

func about(w http.ResponseWriter, r *http.Request) {
	// TODO: render templates: base.html and about.html
    template.Must(
		template.ParseFiles("templates/about.html", "templates/base.html"),
	).Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	// TODO: render templates: base.html and contact.html
    template.Must(
		template.ParseFiles("templates/contact.html", "templates/base.html"),
	).Execute(w, nil)
}

func main() {
	fmt.Println("hello world")

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/method", method)
	http.HandleFunc("/path/", path)
	http.HandleFunc("/query", query)
    http.HandleFunc("/", index)
    
    http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	// TODO: /about
	// TODO: /contact

	log.Panic(http.ListenAndServe("127.0.0.1:8000", nil))
}
