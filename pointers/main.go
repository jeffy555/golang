package main

import (
	"fmt"
)

func main() {
	var a int = 5
	var b *int = &a
	fmt.Println(a)
	fmt.Println(b)
	s := []int{1, 2, 3, 4}
	address1 := &s[0]
	fmt.Println(address1)
	array := []int{10, 11, 12, 13, 14}
	array[3] = 52
	array2 := array
	fmt.Println(array2)
	slice := []int{15,16,17,18,19}
	slice1 := slice
	
	//Unique feature in Slice. When you change the value of an index and use the Println then you can see the value will be changed for both the slice1 and slice variable.
	//Previously for Array you cannot do that, When the array is mapped to array2 and if you change the index of array1 then it will not affect the array2, both will be considered different
	
	
	fmt.Println(slice, slice1)
	slice1[3] = 99
	fmt.Println(slice, slice1)
}
