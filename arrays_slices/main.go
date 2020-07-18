package main

import (
	"fmt"
)

func main() {

	// Array declaration
	grades := [4]int{45, 23, 66, 43} //Type of the data we going to store
	fmt.Println(grades)
	var students [3]string
	students[0] = "Lisa"
	students[1] = "Moni"
	students[2] = "Jeffy"
	fmt.Printf("Students: %v", students)
	fmt.Printf("\nStudents: %v", students[2])
	fmt.Printf("\nNumber of Students: %v\n", len(students))

	//In GO we are actually creating an array of values. Other languages they will have a pointer to it.
	//a := [...]int{2, 4, 5}

	//Slice is initialized as a literal and it looks exactly like an array --> a :=[]int{1,2,4}
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(a)

	//Different way of creating a slice
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b1 := a1[:5]  //Slice first 5 element
	c1 := a1[3:]  //Slice from 3rd element to last
	d1 := a1[3:6] //slice from 4th to 7th element
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1)

	//Make function to create an array contains 2 or 3 arguments, (datatype, len, capacity)

	m1 := make([]int, 3, 6)
	fmt.Println(m1)
	fmt.Printf("Length: %v\n", len(m1))
	fmt.Printf("Capacity: : %v\n", cap(m1))

	//Empty Slice: If append with a value then the capacity will be updated and added to the slice

	m2 := []int{}
	m2 = append(m2, 3, 4, 5, 6, 7, 8, 23, 56, 88)
	fmt.Println("Slice of M2", m2)
	fmt.Printf("Slice of M2 - Length: %v\n", len(m2))
	fmt.Printf("Slice of M2 - Capacity: : %v\n", cap(m2))

	m3 := []int{10, 20, 30}

	//Below append will not work. We need apply m4 := append(m2, m3...) to append another slice into the slice m4. Appending between slices is possible only through this step

	//m4 := append(m2, m3)
	m4 := append(m2, m3...)
	fmt.Println("Slice of M4", m4)
	m5 := m4[5:]
	fmt.Println("Slice of M5: ",m5)
	
	//Appending two slices and concatenating some index values
	
	m6 := append(m2[:4], m5[2:]...)
	fmt.Println("Slice of M6: ",m6)
	
	
	
}
