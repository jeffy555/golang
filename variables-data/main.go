package main

import (
	"log"
	"os"
	"text/template"
)

var value *template.Template

func init() {
	value = template.Must(template.ParseFiles("index.html"))
}

func main() {
	err := value.ExecuteTemplate(os.Stdout, "index.html", "Jesus Christ")
	if err != nil {
		log.Fatalln(err)
	}
}
