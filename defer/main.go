package main

import "fmt"

func print(arg int) {
	defer fmt.Println(arg)
}
func main() {
	fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	defer fmt.Println("4")
	defer print(5)
}
