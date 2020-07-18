package main

import (
	"fmt"
	"reflect"
)

//Fields have any type of data. Powerful.
type docter struct {
	Name    string
	id      int
	patient []string
}

type Animal struct {
	Name   string `required max: "100"` 
	Origin string
}

type bird struct {
	Animal
	speed  float64
	canfly bool
}

func main() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
	doc1 := docter{

		Name:    "Jefferson",
		id:      5,
		patient: []string{"Moni", "Trephy", "Sheila"},
	}

	fmt.Println(doc1)
	fmt.Println(doc1.id) //Drill down to Id's

	//Declaring an anonymous struct
	doc2 := struct{ name string }{name: "Keenu Reaves"}
	fmt.Println(doc2)

	//Unlike Map's struct are value types
	//Go does not have inheritance

	doc3 := doc2
	doc3.name = "John wick"

	fmt.Println(doc3)

	//Embedding a struct. Good idea for web framework. 

	b := bird{}
	b.Name = "Crow"
	b.Origin = "India"
	b.speed = 23.56
	b.canfly = false
	fmt.Println(b.Name)
	
	b2 := bird{
	
		Animal: Animal{Name: "Emu", Origin: "Australia"},
		speed: 40,
		canfly: false,
	
	}
	
	fmt.Println(b2.speed)
	
	//Tags to describe specific information about the name field
}
