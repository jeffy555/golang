package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

type cars struct {
	Name  string
	Model string
}

type features struct {
	Sensor bool
	Cars   []cars
}

type Features []features

func main() {
	f := Features{
		features{
			Sensor: true,
			Cars: []cars{
				cars{
					Name:  "Swift",
					Model: "2020",
				},
			},
		},
		features{
			Sensor: false,
			Cars: []cars{
				cars{
					Name:  "Ford",
					Model: "2015",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, f)
	if err != nil {
		log.Fatalln(err)
	}
}
