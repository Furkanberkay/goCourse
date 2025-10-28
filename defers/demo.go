package defers

import "fmt"

type car struct {
	name  string
	price int
}

func NewCar(name string, price int) car {
	return car{name: name, price: price}
}

func (c car) Save() {
	fmt.Println("kaydedildi araç ismi :", c.name)
}

func Demo() {
	car := NewCar("audi", 120)
	defer car.Save()

	car = NewCar("merco", 200)
	fmt.Println("işlem başarılı", car.name)
}
