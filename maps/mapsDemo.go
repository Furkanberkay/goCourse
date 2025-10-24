package maps

import "fmt"

func MapsDemo() {
	demoMaps := make(map[string]string)

	demoMaps["selamla"] = "selamlar"
	demoMaps["vedalaş"] = "baybay"

	for _, v := range demoMaps {
		fmt.Println(v)
	}
}
