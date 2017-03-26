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

type Info struct {
	Path *url.URL
	Method string
	Query url.Values
	Header http.Header
	Raw template.HTML
	Now time.Time
}

func getInfo(r *http.Request) Info{
	r.ParseForm()
	return Info{
		Path: r.URL,
		Method: r.Method,
		Query: r.Form,
		Header: r.Header,
		Raw: template.HTML("<b>bold text</b>"),
		Now: time.Now(),
	}
}
func index(w http.ResponseWriter, r *http.Request) {
	ctx := getInfo(r)
	template.Must(
		template.ParseFiles("templates/index.html", "templates/base.html"),
	).Execute(w, ctx)
}

func about(w http.ResponseWriter, r *http.Request) {
	template.Must(
		template.ParseFiles("templates/about.html", "templates/base.html"),
	).Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	template.Must(
		template.ParseFiles("templates/contact.html", "templates/base.html"),
	).Execute(w, nil)
}

func cookies(w http.ResponseWriter, r *http.Request) {
	template.Must(
		template.ParseFiles("templates/cookies.html", "templates/base.html"),
	).Execute(w, r.Cookies())
}

func cookieAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintln(w, "only POST allowed")
		return
	}
	r.ParseForm()
	name := r.Form.Get("name")
	value := r.Form.Get("value")
	if name == "" || value == "" {
		fmt.Fprintln(w, "name and value are required")
		return
	}

	cookie := &http.Cookie{Name: name, Value: value}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/cookies", http.StatusTemporaryRedirect)
}

func cookieDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintln(w, "only POST allowed")
		return
	}
	r.ParseForm()
	name := r.Form.Get("name")
	if name == "" {
		fmt.Fprintln(w, "name is required")
		return
	}

	cookie := &http.Cookie{Name: name, Expires: time.Now()}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/cookies", http.StatusTemporaryRedirect)
}


type Response struct {
	Success bool `json:"success"`
	Error string `json:"error"`
	Result interface{} `json:"result"`
}

func fail(w http.ResponseWriter, msg string) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(Response{Success: false, Error: msg})
}

func success(w http.ResponseWriter, result interface{}) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(Response{Success: true, Result: result})
}

func apiInfo(w http.ResponseWriter, r *http.Request) {
	success(w, getInfo(r))
}

func apiCookies(w http.ResponseWriter, r *http.Request) {
	success(w, r.Cookies())
}

func apiCookieAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fail(w, "only POST allowed")
		return
	}
	r.ParseForm()
	name := r.Form.Get("name")
	value := r.Form.Get("value")
	if name == "" || value == "" {
		fail(w, "name and value are required")
		return
	}

	cookie := &http.Cookie{Name: name, Value: value}
	http.SetCookie(w, cookie)
	success(w, "added")
}

func apiCookieDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fail(w, "only POST allowed")
		return
	}
	r.ParseForm()
	name := r.Form.Get("name")
	if name == "" {
		fail(w, "name is required")
		return
	}

	cookie := &http.Cookie{Name: name, Expires: time.Now()}
	http.SetCookie(w, cookie)
	success(w, "added")
}

func main() {
	fmt.Println("hello world")

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/method", method)
	http.HandleFunc("/path/", path)
	http.HandleFunc("/query", query)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/cookies", cookies)
	http.HandleFunc("/cookie-add", cookieAdd)
	http.HandleFunc("/cookie-delete", cookieDelete)

	// API:
	http.HandleFunc("/api/info", apiInfo)
	http.HandleFunc("/api/cookies", apiCookies)
	http.HandleFunc("/api/cookie-add", apiCookieAdd)
	http.HandleFunc("/api/cookie-delete", apiCookieDelete)

	log.Panic(http.ListenAndServe("127.0.0.1:8000", nil))
}
