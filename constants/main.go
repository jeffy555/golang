package main

import (
	"fmt"
)

//Inner const always overshadows the variable in the outer constant. I created a constant outside the function and the same inside function. Inside the main function wins.
//var myConst1 int = 44

func main() {

	//General Constant

	//const myConst

	//Type Constant

	const myConst1 int = 42
	//myConst1 = 45
	//Cannot re-assign to a constant value
	fmt.Printf("%v, %T\n", myConst1, myConst1)

	//Adding a const variable to a normal variable

	const a1 int = 45
	const myConst2 int = 53
	fmt.Printf("%v, %T\n", a1+myConst2, a1+myConst2)

	//Adding a const variable with a different type : it will not work, we need to convert it to int type

	const a2 = 53
	var f2 float32 = 54.234
	var f3 int
	f3 = int(f2)

	fmt.Printf("%v, %T\n", a2+f3, a2+f3)

	//Embedded constant
	//Iota : counter

	const (
		i1 = iota
		i2
		i3
	)

	fmt.Printf("%v\n", i1)
	fmt.Printf("%v\n", i2)
	fmt.Printf("%v\n", i3)

	// It will still work since we applied a pattern so that Iota can recognize and work accordingly. Iota will scope to the constant block.

	const (
		dog = iota
		cat
		lion
	)

	//var animal int = lion
	//var animal int
	var animal int = dog
	fmt.Printf("%v\n", animal == dog)

	//var animal int = lion: Value of Iota will be true in this case, Because we have an int of animal equal to the lion.
	//var animal int: Value of Iota will be false in this case.
	// True if it is assigned to iota value

	const (
		_ = iota + 6
		i4
		i5
		i6
	)
	fmt.Printf("%v\n", i6)
	
	//Iota can be started with zero value and then iota will follow the procedure and add it accordingly.
	
}
