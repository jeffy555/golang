package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("test.html"))
}

func main() {
	u1 := user{
		Name:  "Moses",
		Motto: "Freedom for Israel through God",
		Admin: false,
	}
	u2 := user{
		Name:  "Joshua",
		Motto: "Caputre Jericho",
		Admin: false,
	}
	u3 := user{
		Name:  "Jesus",
		Motto: "Love",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
