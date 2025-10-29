package stringfunctions

import (
	"fmt"
	s "strings"
)

func Demo1() {
	isim := "Berkay"

	fmt.Println(s.Count(isim, "b"))
	fmt.Println(s.Contains(isim, "e"))
	fmt.Println(s.Index(isim, "y"))
}

func Demo2() {
	ad := "furkan"
	fmt.Println(s.HasPrefix(ad, "ber"))
	fmt.Println(s.HasSuffix(ad, "kan"))

	harfler := []string{"a", "c", "d", "t"}

	sonuc := s.Join(harfler, "")

	fmt.Println(sonuc)
}
