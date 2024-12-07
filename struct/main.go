package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	Age       int
	Class     string
	createdAt time.Time
}

func main() {
	// Create basic struct
	student1 := Student{
		Name:  "John",
		Age:   20,
		Class: "10th",
	}
	fmt.Println(student1)

	
	// modify one struct field using pointer
	ar := &student1
	fmt.Println("Before ->", student1, ar)
	ar.Class = "5th"
	fmt.Println("After ->", student1, ar)


	// print struct to custom function
	fmt.Println(printStruct(student1))


	// modify struct field using pointer & function
	fmt.Println(student1, changeOrAddFieldValue(&student1))
	fmt.Println(ar)


	// create inline struct
	teacher := struct {
		passYear int
		name     string
		age      int
	}{
		passYear: 2020,
		name:     "John",
		age:      20,
	}
	fmt.Println("Inline struct ->", teacher)
}

func changeOrAddFieldValue(s *Student) *Student {
	s.createdAt = time.Now()
	s.Age = 78
	s.Class = "3rd"
	return s
}

func printStruct(s Student) Student {
	return s
}
