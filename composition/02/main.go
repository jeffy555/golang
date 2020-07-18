package main

import (
	"text/template"
)

var tpl *template.Template

type ruler struct {
	Name      string
	Regin     int
	character string
}

type worshipped struct {
	God    string
	Rulers []ruler
}

type prophets struct {
	Name       string
	Good, Evil worshipped
}

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	p1 := prophets{
		Name: "Isaiah",
		Good: worshipped{
			God: "Jehovah",
			Rulers: []ruler{
				ruler{"Hezekiah",26,"Did what his right in the sight of the lord"}
			},
		},
		Evil: worshipped{
			God: "Baal",
			Rulers: []ruler{
				ruler{"Ahab",21,"Did what his wrong in the sight of the lord"}
			},
		},
	}
}
