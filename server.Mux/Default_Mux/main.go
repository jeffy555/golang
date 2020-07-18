package main

import (
	"io"
	"net/http"
)

func main() {

	//Handlefunc has "func HandleFunc(pattern string, handler func(ResponseWriter, *Request))"
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar/", bar)
	http.HandleFunc("/me/", name)
	http.HandleFunc("/hello/", hello)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func bar(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "bar ran")
}

func name(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello Jeffy")
}
func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Olaaa Again")
}
