package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 20 {

		i++
		if i == 19 {
			break
		}
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)

	}
	fmt.Println("-------------")
	s := []int{1, 2, 3, 4, 5, 6}
	//Key and value
	for k, v := range s {
		fmt.Println(k, v)
	}
	fmt.Println("-------------")
	statepopulations := map[string]int{

		"TN":        234230249234,
		"Bangalore": 23432123,
		"Andhra":    93249324,
		"Goa":       23423411,
	}

	for k, v := range statepopulations {

		fmt.Println(k, v)
	}
}
