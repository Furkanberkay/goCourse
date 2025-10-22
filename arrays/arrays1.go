package arrays

import "fmt"

func ArraysDemo1() {
	numbers := [...]int{2, 3, 4, 5}

	theBigestNumber := numbers[0]

	for n := 0; n < len(numbers); n++ {
		if numbers[n] > theBigestNumber {
			theBigestNumber = numbers[n]
		}
		fmt.Println(theBigestNumber)

	}
}
