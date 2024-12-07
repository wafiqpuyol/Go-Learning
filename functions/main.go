package main

import (
	"fmt"
)

/* ---- Create a HOF -----*/
func vals(a int, b func(int) int) (int, int) {
	return a, b(3)
}
func loop(a int) int {
	return 60 + a
}

/* -------- Variadic Functions -------- */
func variadic(args ...int) int {
	total := 0
	for _, val := range args {
		total += val
	}
	return total
}

func main() {
	// Create a HOF
	a, b := vals(34, loop)
	fmt.Println(a)
	fmt.Println(b)

	// Create a function with variadic arguments
	arr := []int{34, 23, 67}
	fmt.Println("Result of the func ->", variadic(arr...))
}
