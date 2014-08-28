package main

import (
	"fmt"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (st Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, st)
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"hello", ":", "Gophers!"})
	http.ListenAndServe("localhost:4000", nil)
}

