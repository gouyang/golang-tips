package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

type Struct struct {
	Greeting, Punct, Who string
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
}

func main() {
	var h Hello
	http.Handle("/string", String("Test string."))
	http.Handle("/hello", h)
	var s = Struct{
		"Hello",
		":",
		"Gophers",
	}
	http.Handle("/struct", s)
	http.ListenAndServe(":8081", nil)
}
