package main

import (
	"fmt"
	"goLesson/structs"
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

	demo := structs.StructsRun("12", "berkay", 211, "merhabalar")

	fmt.Println(demo)

}
