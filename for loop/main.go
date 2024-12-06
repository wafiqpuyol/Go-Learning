package main

import "fmt"

func main() {

	// Traditional for loop
	for i := 0; i < 10; i++ {
		println(i)
		fmt.Println(i)
	}

	// Range loop
	for r := range 40 {
		println(r)
	}

	// Infinite loop
	for {
		fmt.Println("Wafiq")
	}
}
