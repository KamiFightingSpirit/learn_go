package banana

var BananaScopedVar = "I am a variable that is hopefully only scoped to banana package"

type Person struct {
	Name     string
	Age      int
	Location string
}

func AddBananas(a int, b int, arg string) int {
	var sum int

	if arg == "+" {
		sum = a + b
	}

	return sum
}

func MinusBanana(a int, b int, arg string) int {
	var result int

	if arg == "-" {
		result = a - b
	}

	return result
}
