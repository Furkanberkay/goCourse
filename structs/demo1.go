package structs

func StructsRun(id string, name string, unitPrice float64, description string) demoStruct {
	exemple1 := demoStruct{id, name, unitPrice, description}
	return exemple1
}

type demoStruct struct {
	id         string
	name       string
	unitPrice  float64
	descriptin string
}
