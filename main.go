package main

import (
	"fmt"
	"goLesson/channels"
)

func main() {
	// veriables.Demo1()
	// conditionals.Demo1()
	// conditionals.Demo2()
	// loops.Demo1()
	// loops.Workshop1()
	// loops.Workshop2()
	// arrays.ArraysDemo1()
	// arrays.MultiDimensionalArrays()
	// slices.SlicesDemo1()
	// slices.Demo2()
	// result, _ := functions.VariadicFunction(1, 2, 3, 4)
	// fmt.Println(result)

	// maps.MapsDemo()
	// numberDemo := 10
	// var numberfunc = pointers.NumberAddOne(&numberDemo)

	// fmt.Println(numberDemo)
	// fmt.Println(numberfunc)

	// demo := structs.StructsRun("12", "berkay", 211, "merhabalar")

	// fmt.Println(demo)

	// var customer = structs.AddCustomer("berkay", "ozcan", 23)
	// customer.Ageup()
	// fmt.Println(customer.Age)

	// go goroutins.EvenNumbers()
	// goroutins.OddNumbers()

	// time.Sleep(2 * time.Second)
	// fmt.Println("bitti")
	//

	// var f1 = goroutins.CreateFruit("1221", "ayva", 123.2)

	// var fruidName = f1.Name
	// var fruidId = f1.Id
	// var fruidPrice = f1.Price

	// fmt.Println("meyve id'si : ", fruidId, "meyve ismi :", fruidName, "meyve fiyatı : ", fruidPrice)

	datas := channels.GetData()

	fmt.Println(datas)

}
