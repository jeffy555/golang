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
	tpl, err := template.ParseFiles("one.gmao")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	//Always consider tpl as a pointer to a template, We can take tpl as a container that stores as much as Parse files into it.
	tpl, err = tpl.ParseFiles("two.gmao", "vespa.gmao")
	if err != nil {
		log.Fatalln(err)
	}
	//ExecuteTemplate takes in writer, name and data
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
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
