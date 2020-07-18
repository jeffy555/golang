package main

import (
	"fmt"
)

//Not only struct you can do any type of custom to implement interface
func main() {
	//Polymorpheic behaviour
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Jeffy"))

	sa1 := student{
		person: person{
			First: "Jefferson",
			last:  "Immanuel",
		},
		pass: true,
	}
	sa2 := student{
		person: person{
			First: "Trephy",
			last:  "Manuel",
		},

		pass: false,
	}
	p1 := person{

		First: "Jeff",
		last:  "Immanuel",
	}
	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.Speak()
	sa2.Speak()
	bar(sa1)
	bar(sa2)
	bar(p1)
	p1.Speak()

}

//In Go we don't explicitly implement interfaces, Instead we do it implicitly.

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

//Creating a method and has the significance of Writer interface

func (cw ConsoleWriter) Write(data []byte) (int, error) {

	n, err := fmt.Println(string(data))
	return n, err
}

//Embedding interface, So the WriterCloser will be implemented if an object as Close method and Write method on it.
// A value can be of more than one type
//Because of a method Speak() attached to the student function then the value is also of type human interface and also of type student.

type human interface {
	Speak()
}

func bar(h human) {
	//Assertion
	switch h.(type) {
	case person:
		fmt.Println("I am passed to a Person of Person", h.(person).First)
	case student:
		fmt.Println("I am passed to a  student of student", h.(student).First)
	}
}

type person struct {
	First string
	last  string
}

type student struct {
	person
	pass bool
}

func (p person) Speak() {

	fmt.Println("I am a person", p.First, p.last)
}
func (s student) Speak() {

	fmt.Println("I am a student", s.First, s.last)
}
