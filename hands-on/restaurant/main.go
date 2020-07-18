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

type dish struct {
	Food1 string
	Food2 string
	Food3 string
}

type items struct {
	Name   string
	Timing string
	Dish   []dish
}

type Items []items

func main() {
	i := Items{
		items{
			Name:   "Hotel Jeffy",
			Timing: "Breakfast",
			Dish: []dish{
				dish{
					Food1: "Idly/Vadakari",
					Food2: "Poori",
					Food3: "Dosai",
				},
			},
		},
		items{

			Name:   "Hotel Jeffy",
			Timing: "Lunch",
			Dish: []dish{
				dish{
					Food1: "Chicken Biriyani",
					Food2: "Fish",
					Food3: "Prawn",
				},
			},
		},
		items{
			Name:   "Hotel Jeffy",
			Timing: "Dinner",
			Dish: []dish{
				dish{
					Food1: "Buriito",
					Food2: "Pizza",
					Food3: "Burger",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, i)
	if err != nil {
		log.Fatalln(err)
	}
}
