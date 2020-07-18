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
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("jeffy.html")
	//Creating an new html file

	if err != nil {
		log.Println("Error creating a file", err)
	}

	defer nf.Close()

	// Execute the Stdout which implements writer interface, passing in no data
	//Executes takes a writer and a data
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
