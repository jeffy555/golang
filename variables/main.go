//https://play.golang.org/p/VMq2pw5iNnG

package main

import (
	"fmt"
	"strconv"
)

var A int = 42

func main() {

	// 3 types of Variable declaration

	// First type is valuable if we want to use this variable to loop in or to use in if statements

	//var i int
	//i = 42
	//fmt.Println(i)
	//i = 27
	//i = 5
	//fmt.Println(i)

	//Second type will have more control over the declared variable.

	//var i int = 42.8 This is not possible in the second type, If you assign the variable as int then you have the control over the value. If you change the value

	//then you need to change the type as well.

	//fmt.Printf("%v, %T", i, i)

	//Third type: You just declaring a variable and your not mentioning the type of the variable. It is picking up automatically. see the below example.
	//I defined a float type but did not mention the type in the declaration, but the output clearly shows the type of the variable as well.

	//i := 42.894
	//fmt.Printf("%v, %T", i, i)

	//var (
	//	actorName string = "Jefferson Immanuel"
	//	series    string = "GOT"
	//	ID        int    = 5
	//)

	//Below example will not work in variable. variable i is already initialized, you can either map it to new value using i = 4 or create a new variable.

	//	var i int = 5
	//i := 5
	//	fmt.Println(i)

	//If variables are declared but not used then it will throw an error stating the variable is not used. Error b declared not used. This is also similar
	//to the variables in uppercase and lowercase, you cannot declare a variable in uppercase and call in the print statement with lowercase character. Same case
	//even if the variable is declared outside the main function.
	//If the variable I is initialized outside main function and then small i is also initialized inside the main function then
	//the variables are considered to be a two different variables

	//Always use uppercase variables for best practices.

	//	var a int = 4
	//	fmt.Println(a)
	//	b := 45

	//Convert a variable from one type to another

	var i float32 = 42.43
	fmt.Printf("%v, %T\n", i, i)
	var j int
	j = int(i)
	fmt.Printf("%v, %T", j, j)

	//Convert a int to String using a package called "strconv"

	var n int = 5
	fmt.Printf("%v, %T\n", n, n)
	var m string
	m = strconv.Itoa(n)
	fmt.Printf("%v, %T\n", m, m)
}
