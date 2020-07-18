package main

import (
	"fmt"
)

func main() {
	m1 := mystruct{

		greet:   "Hello",
		luckyme: "Your lucky",
	}
	m1.Apps()
	for i := 0; i < 5; i++ {
		Call("Hello Team", i)
	}

	greeting := "Hello"
	wishme := "Jeffy"
	GreetMe(greeting, wishme)
	//The value of the variable is changed in the function GreetMe but still if we see the below println statement it will print the old one.
	//The value that we pass in the function goes to that function alone. Passing the data by Value.
	//But when we pass this as a pointer then we can manipulate the value from the function to the variable in main.. See next example
	fmt.Println(wishme)

	g1 := "Hello"
	w2 := "Moni"
	Pointers(&g1, &w2)
	fmt.Println(g1, w2)

	s1 := Sum(1, 2, 3, 4, 5)
	fmt.Println("The sum of the value is", s1)

	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

}

func Call(msg string, idx int) {
	fmt.Println(msg, idx)
}

func GreetMe(greeting, wishme string) {
	fmt.Println(greeting, wishme)
	//I am re-assigning the variable wishme to Trephy and printing it out.
	wishme = "Trephy"
	fmt.Println(wishme)
}

func Pointers(g1, w2 *string) {

	fmt.Println(*g1, *w2)
	*w2 = "Sheila"
}

func Sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {

		result += v

	}
	return result

}

func divide(a, b float64) (float64, error) {

	if b == 0.0 {

		return 0.0, fmt.Errorf("Error cannot be divided by 0")

	}
	return a / b, nil
}

//Method is basically a function they just have a little bit syntax change, providing a context that the function is executing in.

type mystruct struct {
	greet   string
	luckyme string
}

func (m1 mystruct) Apps() {

	fmt.Println(m1.greet, m1.luckyme)

}
