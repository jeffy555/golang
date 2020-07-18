package main

import (
	"log"
	"os"
	"text/template"
)

//Variable tpl is a pointer to a template

var tpl *template.Template

//func Must(t *Template, err error) *Template
//Must is a helper that wraps a call to a function returning (*Template, error) and panics if the error is non-nil.
//It is intended for use in variable initializations such as

var t = template.Must(template.New("name").Parse("text"))

func init() {
	tpl = template.Must(template.ParseFiles("index2.html"))
}

func main() {
	xs := []string{"zero", "One", "Two", "Three"}

	// //Anonymous struct
	data := struct {
		//Numbers has to be in caps N. Otherwise we cannot pull the value in index.html.
		Numbers []string
		God     string
	}{
		xs,
		"Jesus Christ",
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
