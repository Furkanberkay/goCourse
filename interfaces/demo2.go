package interfaces

import "fmt"

type CreditCalculate interface {
	Calculate() float64
}

type Mortgage struct {
	creditPaymentTotal float64
	rate               float64
}

func NewMortgage(creditPaymentTotal float64, rate float64) (Mortgage, error) {
	if creditPaymentTotal <= 0 || rate <= 0 {
		return Mortgage{}, fmt.Errorf("creditPaymentTotal and rate must be  > 0")
	}
	return Mortgage{creditPaymentTotal: creditPaymentTotal, rate: rate}, nil
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

func NewCar(creditPaymentTotal float64, carInfo string, rate float64) (Car, error) {
	if creditPaymentTotal <= 0 || rate <= 0 {
		return Car{}, fmt.Errorf("creditPaymentTotal  rate and carInfo  must be  > 0")
	}
	return Car{creditPaymentTotal: creditPaymentTotal, carInfo: carInfo, rate: rate}, nil
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 12
}

func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 12
}

func CalculateMonthlyPayment(c []CreditCalculate) float64 {
	var total float64 = 0

	for _, v := range c {
		total += v.Calculate()
	}
	return total
}

func Demo2() {
	car, _ := NewCar(12000, "merco s", 200)
	mortgage, _ := NewMortgage(20000, 200)

	credits := []CreditCalculate{car, mortgage}

	total := CalculateMonthlyPayment(credits)
	fmt.Println("toplam borç : ", total)
}
