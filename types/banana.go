package banana

func AddBananas(a int, b int, arg string) int {
	var sum int

	if arg == "+" {
		sum = a + b
	}

	return sum
}
