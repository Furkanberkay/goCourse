package goroutins

import "fmt"

func EvenNumbers() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println("Çift Sayi : ", i)
	}
}

func OddNumbers() {
	for i := 1; i <= 9; i += 2 {
		fmt.Println("Tek Sayi : ", i)
	}
}
