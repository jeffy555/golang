package main

import (
	"log"
	"os"
	"text/template"
)

var value *template.Template

type characters struct {
	Name  string
	Motto string
}

func init() {
	value = template.Must(template.ParseFiles("char.html"))
}

func main() {
	char := characters{
		Name:  "Moses",
		Motto: "Free the Israel from Egypt",
	}
	// char1 := characters{
	// 	Name: "Noah",
	// 	Motto: "Ark to save the human from Satan",
	// }
	// char2 := characters{
	// 	Name: "Joshua",
	// 	Motto: "War to capture Milk and honey",
	// }
	// char3 := characters{
	// 	Name: "Jesus",
	// 	Motto: "Love",
	// }
	// ch := {char,char1,char2,char3}
	err := value.ExecuteTemplate(os.Stdout, "char.html", char)
	if err != nil {
		log.Fatalln(err)
	}
}
