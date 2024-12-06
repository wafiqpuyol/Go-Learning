package main

import (
	"fmt"
)

func main() {
	const myName = "wafiq"
	const Range = 51
	var myAge = 50

	for i := range Range {
		if myName == "wafiq" {
			myAge = i
			if myAge == 50 {
				fmt.Println("wafiq", myAge)
			}
		}
	}

	for i := 0; i < 20; i++ {
		if increment := i + 1; increment > 10 {
			fmt.Println("My Age is", increment)
		}
	}
}
