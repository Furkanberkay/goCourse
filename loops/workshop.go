package loops

import "fmt"

func Workshop1() {
	tutulanSayi := 34
	tahminEdilenSayi := 0

	fmt.Print("bir sayi tutunuz")
	fmt.Scanln(&tahminEdilenSayi)

	for tahminEdilenSayi != tutulanSayi {

		if tahminEdilenSayi < tutulanSayi {
			fmt.Println("tahmini büyüt")
			fmt.Scanln(&tahminEdilenSayi)
		}

		if tahminEdilenSayi > tutulanSayi {
			fmt.Println("tahmini küçült")
			fmt.Scanln(&tahminEdilenSayi)
		}
	}

	fmt.Println("tebrikler bildiniz sayi : ", tahminEdilenSayi)
}
