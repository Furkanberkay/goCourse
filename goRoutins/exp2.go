package goroutins

type Fruit struct {
	Id    string
	Name  string
	Price float32
}

func CreateFruit(id string, name string, price float32) Fruit {
	return Fruit{
		Id:    id,
		Name:  name,
		Price: price,
	}
}
