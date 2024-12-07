package main

import (
	"fmt"
)

func main() {
	myMap1 := make(map[string]int)
	myMap1["John"] = 69
	myMap1["Wafiq"] = 50

	// Loop over the map using range
	for key, val := range myMap1 {
		fmt.Println(key, "---", val)
	}

	// create map on the fly
	myMap2 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(myMap2)

	// check if key exists
	_, ok := myMap2["a"]
	if ok {
		fmt.Println("key exists")
	} else {
		fmt.Println("key does not exist")
	}
}
