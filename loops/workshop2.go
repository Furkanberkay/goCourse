package loops

import "fmt"

func Workshop2() {
	for { // while(true)
		var n int
		fmt.Print("Bir sayı giriniz (çıkmak için 0): ")
		fmt.Scanln(&n)
		if n == 0 { // çıkış koşulu
			fmt.Println("Çıkılıyor...")
			return
		}

		if n < 2 {
			fmt.Println("asal değil")
			continue
		}
		if n == 2 {
			fmt.Println("asal")
			continue
		}
		if n%2 == 0 {
			fmt.Println("asal değil")
			continue
		}

		isPrime := true
		for d := 3; d*d <= n; d += 2 {
			if n%d == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println("asal")
		} else {
			fmt.Println("asal değil")
		}
	}
}
