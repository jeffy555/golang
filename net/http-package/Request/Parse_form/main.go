package main

import (
	"html/template"
	"log"
	"net/http"
)

type jeffy int

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func (m jeffy) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.html", req.Form)
}

func main() {
	var l jeffy
	http.ListenAndServe(":8081", l)
}
