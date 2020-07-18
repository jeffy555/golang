package main

import (
	"log"
	"os"
	"text/template"
)

var value *template.Template

func init() {
	value = template.Must(template.ParseFiles("queens.html"))
}

func main() {
	kings := []string{"Sheila", "Trehy", "Moni"}
	err := value.ExecuteTemplate(os.Stdout, "queens.html", kings)
	if err != nil {
		log.Fatalln(err)
	}
}
