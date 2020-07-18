package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

type hotels []hotel

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	h := hotels{
		hotel{
			Name:    "Hote California",
			Address: "Wisconsin street",
			City:    "NewYork",
			Zip:     "52423",
			Region:  "Southern",
		},
		hotel{
			Name:    "H",
			Address: "4",
			City:    "L",
			Zip:     "95612",
			Region:  "Northern",
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
