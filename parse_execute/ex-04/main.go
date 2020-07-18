//Parse Template

package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//Glob (*.html)
	tpl, err := template.ParseGlob("template/*")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	//Always consider tpl as a pointer to a template, We can take tpl as a container that stores as much as Parse files into it.
	tpl, err = tpl.ParseFiles("two.html", "vespa.html")
	if err != nil {
		log.Fatalln(err)
	}
	//ExecuteTemplate takes in writer, name and data
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "two.html", nil)
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
