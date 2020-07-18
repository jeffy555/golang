package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {

	//Handlefunc has "func HandleFunc(pattern string, handler func(ResponseWriter, *Request))"
	http.Handle("/", http.HandlerFunc(foo))
	http.Handle("/bar/", http.HandlerFunc(bar))
	http.Handle("/me/", http.HandlerFunc(jeffy))
	http.Handle("/hello/", http.HandlerFunc(hello))
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func bar(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "bar ran")
}

func jeffy(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("name.html")
	if err != nil {
		log.Fatalln(err)
	}
	//Executetemplate --> func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	err = tpl.ExecuteTemplate(w, "name.html", "jeffy")
	if err != nil {
		log.Fatalln(err)
	}
}
func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Olaaa Again")
}
