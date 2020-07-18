package main

import (
	"fmt"
)

func main() {

	//Maps map[string]int --> string (key) int (value)

	statePopulations := map[string]int{

		"TN":        2000000,
		"Hyd":       1000000,
		"Kerala":    50000,
		"Bangalore": 423352,
	}
	fmt.Println(statePopulations)

	//Create using make function

	//sp1 := make(map[string]int)
	//sp2 := make(map[string)int, 10)
	//Manipulating the values in the map

	fmt.Println(statePopulations["TN"])
	statePopulations["Goa"] = 234234
	fmt.Println(statePopulations)
	//deletion
	delete(statePopulations, "Goa")
	fmt.Println(statePopulations)
	
	//If we are not sure if the key is in the map or not we can use the "ok". 
	pop, ok := statePopulations["T1"]
	fmt.Println(pop, ok)
}
