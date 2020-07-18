package main

import (
	"log"
	"os"
	"text/template"
)

var value *template.Template

func init() {
	value = template.Must(template.ParseFiles("char.html"))
}

func main() {
	Characters := map[string]string{
		"Genesis": "Abhraham",
		"Exodus":  "Moses",
		"Numbers": "Joshua",
		"Samuel":  "Samuel",
		"King":    "David",
	}
	err := value.ExecuteTemplate(os.Stdout, "char.html", Characters)
	if err != nil {
		log.Fatalln(err)
	}
}
