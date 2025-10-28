package errorhandling

import (
	"fmt"
	"os"
)

func Demo() {
	f, err := os.Open("test.txt")

	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("dosya yolu hatası : ", pErr.Path)
			return
		} else {
			fmt.Println("dosya bulunamadı")
		}

	} else {
		fmt.Println("dosya :", f.Name())
	}
}
