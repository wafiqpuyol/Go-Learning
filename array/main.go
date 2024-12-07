package main

import (
	"fmt"
)

func main() {
	// random type array
	numArr := []interface{}{1, "wafiq", 3, 55.43}
	fmt.Println(numArr)

	// fixed array with type
	var numArr1 [3]int = [3]int{1, 2, 3}
	fmt.Println(numArr1)
}
