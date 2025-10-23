package functions

import (
	"fmt"
)

func DortIslem(number1 int, number2 int) (int, int, int, float32) {
	toplam := number1 + number2
	cikarma := number1 - number2
	carpma := number1 * number2
	bolme := float32(number1 / number2)

	return toplam, cikarma, carpma, bolme
}

func VariadicFunction(slice ...int) (int, error) {
	total := 0

	if len(slice) == 0 {
		return 0, fmt.Errorf("sıfır")
	}
	for _, v := range slice {
		total += v
	}
	return total, nil
}
