package conditionals

import "fmt"

func Demo1() {
	withdrawAmount := 200.0
	balance := 600.0

	if withdrawAmount <= balance {
		fmt.Println("parayı alıınız")
		balance = balance - withdrawAmount
		fmt.Println("güncel bakiyeniz : " + fmt.Sprintf("%.2f", balance))
	} else {

		fmt.Println("yetersiz bakiye")
	}

}
