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

type item struct {
	Name, Desc string
	Price      float64
}

type meal struct {
	Meal string
	Item []item
}

type menu []meal

type restaurant struct {
	Name string
	Menu menu
}

type restaurants []restaurant

func main() {
	res := restaurants{
		restaurant{
			Name: "Hotel Jeffy",
			Menu: menu{
				meal{
					Meal: "Breakfast",
					Item: []item{
						item{
							Name:  "Idly",
							Desc:  "Rice pan cake",
							Price: 65.02,
						},
						item{
							Name:  "Dosai",
							Desc:  "Made with maavu",
							Price: 55.02,
						},
					},
				},
				meal{
					Meal: "Snacks",
					Item: []item{
						item{
							Name:  "Biscuits",
							Desc:  "Made with Original cocoa",
							Price: 165.32,
						},
						item{
							Name:  "Milk halwa",
							Desc:  "Made with Vanilla",
							Price: 55.32,
						},
					},
				},
				meal{
					Meal: "Dinner",
					Item: []item{
						item{
							Name:  "Biriyani",
							Desc:  "Made with Chicken",
							Price: 365.32,
						},
						item{
							Name:  "Chicken fried rice",
							Desc:  "Made with Chicken",
							Price: 145.32,
						},
					},
				},
			},
		},
		restaurant{
			Name: "Hotel Moni",
			Menu: menu{
				meal{
					Meal: "Breakfast",
					Item: []item{
						item{
							Name:  "Poori",
							Desc:  "flour",
							Price: 35.02,
						},
					},
				},
				meal{
					Meal: "Snacks",
					Item: []item{
						item{
							Name:  "Bhajji",
							Desc:  "Oil bajji",
							Price: 25.32,
						},
					},
				},
				meal{
					Meal: "Dinner",
					Item: []item{
						item{
							Name:  "Mutton Biriyani",
							Desc:  "Made with Mutton",
							Price: 465.32,
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, res)
	if err != nil {
		log.Fatalln(err)
	}
}
