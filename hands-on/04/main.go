package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name   string
	Id     string
	Aadhar string
	PAN    string
}

type company struct {
	Name   string
	Person []person
}

type Company []company

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	h := Company{
		company{
			Name: "MrCooper",
			Person: []person{
				person{
					Name:   "Jefferson Immanuel",
					Id:     "5344232",
					Aadhar: "53432",
					PAN:    "23553",
				},

				person{
					Name:   "Thomas Xavier",
					Id:     "23423",
					Aadhar: "86754",
					PAN:    "8452",
				},
			},
		},
		company{
			Name: "Servion",
			Person: []person{
				person{
					Name:   "Moni Immanuel",
					Id:     "5344232",
					Aadhar: "53432",
					PAN:    "23553",
				},

				person{
					Name:   "Trephy Immanuel",
					Id:     "23423",
					Aadhar: "86754",
					PAN:    "8452",
				},
			},
		},
		company{
			Name: "Teacher",
			Person: []person{
				person{
					Name:   "Sheila Immanuel",
					Id:     "5344232",
					Aadhar: "53432",
					PAN:    "23553",
				},

				person{
					Name:   "David Rajan",
					Id:     "23423",
					Aadhar: "86754",
					PAN:    "8452",
				},
			},
		},
	}
	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
