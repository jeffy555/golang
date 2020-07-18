//Parse Template

package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//Parsefiles take variadic parameters
	//func ParseFiles(filenames...string) and will return (*Template, error)
	//Parse many files in the parameters and will have a pointer to a template
	tpl, err := template.ParseFiles("index.html", "index2.html")
	if err != nil {
		log.Fatalln(err)
	}
	// Execute the Stdout which implements writer interface, passing in no data
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
