package maps

import "fmt"

func Workshop(slice ...int) {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	toplam := 0
	var x = 4

	for _, v := range numbers {
		if v%2 == 0 {
			toplam += v
		}
	}

	fmt.Println(x)

}
