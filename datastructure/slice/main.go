package main

import (
	"log"
	"os"
	"text/template"
)

var value *template.Template

func init() {
	value = template.Must(template.ParseFiles("kings.html"))
}

func main() {
	kings := []string{"David", "Solomon", "Jehu", "Josaiah"}
	err := value.ExecuteTemplate(os.Stdout, "kings.html", kings)
	if err != nil {
		log.Fatalln(err)
	}
}
