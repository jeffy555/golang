//Handler interface needs (ServeHTTP)
//Routing has basic routing / ServerMux or Handlefunc
//We can create our own ServeMux pointer to a serveMux we have a handlefunc that containes a string and func of func(responseWriter, *Request)
//Since ServeMux is implementing a ServeHTTP we can pass it to ListenandServe
//DefaultServeMux uses Handle and Handlefunc
//Below are the best practice solution
package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

//Handlefunc requires something like this "c(res http.ResponseWriter, req *http.Request)"
func main() {
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)
	http.ListenAndServe(":8080", nil)
}

// 1. Type Handler interface with ServeHTTP has two parameters ServeHTTP(res http.ResponseWriter, req *http.Request)
// 2. Any value of type ServeMux will implement the handler interface. When we create a ServeMux we can pass it to the listenandserve, Because Lisetenandserve are once a handler
// 3. Pointer to a ServeMux also has several methods attached to it. They are handle and handlefunc, Handle takes a handler, Handlefunc takes a func(responsewriter and pointer to a request) this is not a handler but a own func.

//Case study
//http.Handle("/dog", http.Handlerfunc(d))
//We can also use this because Handler func contains func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
//Handle is a type of interface that has ServeHTTP(w ResponseWriter, r *Request)
//In case of Handlefunc it has func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
