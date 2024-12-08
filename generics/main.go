package main

import "fmt"

func printClass[M []V, V string](m M, v V) {
	for idx, val := range m {
		if m[idx] == v {
			fmt.Println("value =>", val, "element found index =>", idx)
		} else {
			fmt.Println("element not found")
		}
	}
}
func main() {
	// simple generic function
	student := []string{"name", "class"}
	printClass[[]string, string](student, "name")
}
