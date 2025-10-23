package slices

import "fmt"

func Demo2() {
	// Demonstrating that slices are reference types
	s := []int{1, 2, 3, 4, 5}
	// Creating a new slice that references the same underlying array
	s2 := s
	//
	s2[0] = 100
	fmt.Println("s:", s)
	fmt.Println("s2:", s2)

}
