package main

import (
	"fmt"
)

func main() {
	numbers := 65
	guess := 12
	if guess < 1 {

		fmt.Println("The number your guessing must be greater than 1")
	} else if guess > 100 {
		fmt.Println("The number your guessing must be less than 100")
	} else {

		if guess < 30 {

			fmt.Println("Too low")
		}

		if guess > 70 {
			fmt.Println("Too high")
		}

		if guess == numbers {
			fmt.Println("Perfect")
		}

	}

	//Switch statement
	i := 10
	switch {

	case i <= 10:
		fmt.Println("i is lesser then 10")
		fallthrough
	case i == 10:
		fmt.Println("i is equals 10")
	case i > 10:
		fmt.Println("i is lesser than 10")
	default:
		fmt.Println("i is default")
	}

	//If we want to check multiple conditions or if two conditions are about to match then you can use keyword "Fallthrough"
	// But if we are using fallthrough then it is the duty of a programmer to handle everything properly in the code.
	//Break statement that we use in the below switch statement will help us to save the program memory time to come out of the loop once the condition is satisfied.
	//If you need to go through another condition to get the value then you can use fallthrough

	var n interface{} = 23.5
	switch n.(type) {

	case int:
		fmt.Println("Int")
		break
	case float64:
		fmt.Println("Float")
		break
	case string:
		fmt.Println("String")
		break
	}
}
