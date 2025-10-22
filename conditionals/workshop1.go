package conditionals

import "fmt"

func Demo2() {
	var number1 int = 1
	var number2 int = 2
	var number3 int = 3

	var enBuyuk int = number1

	if number2 > enBuyuk {
		enBuyuk = number2
	}

	if number3 > enBuyuk {
		enBuyuk = number3
	}

	fmt.Println("en buyuk sayi : ", enBuyuk)

}
