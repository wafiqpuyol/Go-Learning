package main

func main() {
	myAge := 23
	switch myAge {
	case 12:
		println("You are 12 years old")
	case 16:
		println("You are 16 years old")
	case 20:
		println("You are 12 years old")
	}

	// multiple case condition
	switch myAge {
	case 12, 16, 23:
		println("You are 12 years old")
	}

	// type switch
	guessType("")
}

func guessType(i interface{}) {
	println("printing interface ==>", i)
	switch i.(type) {
	case string:
		println("string")
		break
	case int:
		println("int")
		break
	case bool:
		println("boolean")
		break
	default:
		println("unknown")
	}
}
