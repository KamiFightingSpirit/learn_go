package main

import (
	"fmt"
	public "learn_go/package1"
	banana "learn_go/types"
)

func main() {
	fmt.Println("hello world")
	result := public.PublicTestDeHuff()
	fmt.Println(result)

	banana_res := banana.AddBananas(4, 5, "+")
	fmt.Println(banana_res)

	banana_res2 := banana.MinusBanana(4, 5, "-")
	fmt.Println(banana_res2)

	banana_var_import := banana.BananaScopedVar
	fmt.Println(banana_var_import)

	var a, b, c = 1, 2, "Coders"
	fmt.Println(a, b, c)

	//block scoping
	{
		var blockVariable int = 5
		fmt.Println(blockVariable)
	}

	//out of scope here
	// fmt.Println(blockVariable)

	//arrays
	var intArray [3]int = [3]int{1, 2, 3}
	fmt.Println(intArray)

	//pointers in Go
	var originalVar int = 42
	var pointerVar = &originalVar

	originalVar = 43
	fmt.Println(originalVar, *pointerVar, pointerVar)

	//structs in Go
	var Person = banana.Person{Name: "Andrew", Age: 37, Location: "San Francisco, CA"}
	fmt.Println(Person)
}
