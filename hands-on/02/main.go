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
}

type region struct {
	Region string
	Hotels []hotel
}

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	h := region{
		Region: "Southern",
		Hotels: []hotel{
			hotel{
				Name:    "Hote California",
				Address: "Wisconsin street",
				City:    "NewYork",
				Zip:     "52423",
			},

			hotel{
				Name:    "H",
				Address: "4",
				City:    "L",
				Zip:     "95612",
			},
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
