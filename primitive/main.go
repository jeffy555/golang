package main

import (
	"fmt"
)

func main() {

	//Boolean type
	// Values are true or false. Not an alias for other types. Zero value is false

	var bo bool
	fmt.Println(bo)

	//String in go is an alias for byte
	//s[2] --> we cannot manipulate the value of a string

	s := "this is a string"
	fmt.Printf("%v, %T\n", s[2], s[2])

	//Convert a string to byte

	s1 := "Hello Jefferson"
	b1 := []byte(s1)
	fmt.Printf("%v, %T\n", b1, b1)

	//Rune Process UTF characters. 

	//var r rune = 'a'

}
