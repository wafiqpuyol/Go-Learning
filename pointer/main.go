package main

import (
	"fmt"
)

/*
take two var preVal and newVal. Change preVal by newVal using pointer
*/
func double(x *int) *int {
	*x = *x * 2
	return x
}

func main() {
	prevVal := 10
	newVal := double(&prevVal)
	fmt.Println(*newVal, prevVal)
	*newVal++
	fmt.Println(*newVal, prevVal)
}
