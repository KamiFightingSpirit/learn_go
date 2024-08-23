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
}
