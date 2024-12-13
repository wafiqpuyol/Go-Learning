package main

import (
	"fmt"
	"strings"
)

type transform func(string) string

func Uppercase(str string) string {
	return strings.ToUpper(str)
}

func Suffix(sfx string) transform {
	return func(str string) string {
		return str + sfx
	}
}

func Transformer(str string, fn transform) string {
	return fn(str)
}

func main() {
	fmt.Println(Transformer("wafiq", Uppercase))
	fmt.Println(Transformer("wafiq", Suffix("_OP")))
}
