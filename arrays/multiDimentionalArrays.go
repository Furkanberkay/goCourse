package arrays

import "fmt"

func MultiDimensionalArrays() {
	numbers := [2][3]int{
		{1, 2, 3},
		{3, 2, 1},
	}

	for a := 0; a < 2; a++ {
		for i := 0; i < 3; i++ {
			fmt.Print(numbers[a][i])
		}
		fmt.Println()

	}

}
