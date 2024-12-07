package main

import "fmt"

func main() {
	var num1 []int
	fmt.Println(num1 == nil)

	/* --------------*/
	var num2 = make([]int, 1, 5)
	num2 = append(num2, 1, 2, 3, 34, 4, 3)
	fmt.Println(num2)
	fmt.Println(len(num2))
	fmt.Println(cap(num2))

	/* --------------*/
	copySlice(num2)
}

func copySlice(num []int) {
	var copiedSlice = make([]int, len(num), cap(num))
	copy(copiedSlice, num)
	fmt.Println(copiedSlice)
	fmt.Println(cap(copiedSlice))
}
