//Parse Template

package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	// 	func Must(t *Template, err error) *Template
	// Must is a helper that wraps a call to a function returning (*Template, error) and panics if the error is non-nil. It is intended for use in variable initializations such as
	// var t = template.Must(template.New("name").Parse("text"))
	//func Must(t *Template, err error) *Template --> Takes a pointer to a template and an error
	//Must will do the error checking.
	tpl = template.Must(template.ParseGlob("*.html"))
}
func main() {

	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	//ExecuteTemplate takes in writer, name and data
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "second.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "one.html", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Execute the Stdout which implements writer interface, passing in no data
	//Executes takes a writer and a data
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
