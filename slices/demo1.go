package slices

import "fmt"

func SlicesDemo1() {
	// Creating a slice using make
	// and demonstrating length and capacity
	// and appending elements to it
	numbers := make([]int, 0)
	numbers = append(numbers, 2, 3, 4, 5)

	fmt.Println("Slice length:", len(numbers))
	fmt.Println("Slice capacity:", cap(numbers))
	fmt.Println("Slice contents:", numbers)

	numbers = append(numbers, 6)
	fmt.Println("After appending 6:")
	fmt.Println("Slice length:", len(numbers))
	fmt.Println("Slice capacity:", cap(numbers))
	fmt.Println("Slice contents:", numbers)
}
