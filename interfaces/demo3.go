package interfaces

import "fmt"

type ShippingCalculator interface {
	Calculate() float64
}

type StandardShipment struct {
	baseFee    float64
	weightKg   float64
	distanceKm float64
	ratePerKg  float64
	ratePerKm  float64
}

func NewStandardShipment(baseFee float64, weightKg float64, distanceKm float64, ratePerKg float64, ratePerKm float64) (StandardShipment, error) {
	if baseFee <= 0 || weightKg <= 0 || distanceKm <= 0 || ratePerKg <= 0 || ratePerKm <= 0 {
		return StandardShipment{}, fmt.Errorf("must be >0")
	}
	obj := StandardShipment{baseFee: baseFee, weightKg: weightKg, distanceKm: distanceKm, ratePerKg: ratePerKg, ratePerKm: ratePerKm}

	return obj, nil
}

func (s StandardShipment) Calculate() float64 {
	result := s.baseFee + (s.weightKg * s.ratePerKg) + (s.distanceKm * s.ratePerKm)

	return result
}

type ExpressShipment struct {
	baseFee            float64
	weightKg           float64
	distanceKm         float64
	priorityMultiplier float64
}

func NewExpressShipment(baseFee float64, weightKg float64, distanceKm float64, priorityMultiplier float64) (ExpressShipment, error) {
	if baseFee <= 0 || weightKg <= 0 || distanceKm <= 0 || priorityMultiplier < 1.1 || priorityMultiplier > 2.0 {
		return ExpressShipment{}, fmt.Errorf("invalid input: values must be >0 and priorityMultiplier must be between 1.1 and 2.0")
	}
	obj := ExpressShipment{baseFee: baseFee, weightKg: weightKg, distanceKm: distanceKm, priorityMultiplier: priorityMultiplier}

	return obj, nil
}

func (e ExpressShipment) Calculate() float64 {
	return (e.baseFee + e.weightKg*0.75 + e.distanceKm*0.40) * e.priorityMultiplier
}

type FragileShipment struct {
	baseFee          float64
	weightKg         float64
	distanceKm       float64
	fragileSurcharge float64
}

func NewFragileShipment(baseFee float64, weightKg float64, distanceKm float64, fragileSurcharge float64) (FragileShipment, error) {
	if baseFee <= 0 || weightKg <= 0 || distanceKm <= 0 || fragileSurcharge < 25 {
		return FragileShipment{}, fmt.Errorf("invalid input: values must be >0 and fragileSurcharge must be bigger than and equal 25 ")
	}
	obj := FragileShipment{baseFee: baseFee, weightKg: weightKg, distanceKm: distanceKm, fragileSurcharge: fragileSurcharge}

	return obj, nil
}

func (f FragileShipment) Calculate() float64 {
	return f.baseFee + f.weightKg*0.60 + f.distanceKm*0.35 + f.fragileSurcharge
}

func CalculateMonthlyLogistics(s []ShippingCalculator) float64 {
	total := 0.0
	for _, v := range s {
		total = total + v.Calculate()
	}
	return total
}

func Demo4() {
	std, _ := NewStandardShipment(50, 12, 120, 0.5, 0.2)
	exp, _ := NewExpressShipment(60, 8, 80, 1.5)
	frg, _ := NewFragileShipment(55, 5, 60, 30)

	collection := []ShippingCalculator{std, exp, frg}

	total := CalculateMonthlyLogistics(collection)

	fmt.Println("Total price : ", total)
}
