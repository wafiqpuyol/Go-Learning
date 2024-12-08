package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	Age       int
	Class     string
	CreatedAt time.Time
	Result    result
}

// Struct embedding
type result struct {
	totalGrade float32
	avgGrade   float32
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
	// print struct through custom function
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

	// create struct with method
	fmt.Println(student1.getClass())

	// embedding struct
	embeddedStruct := Student{
		Name:      "Ishan",
		Age:       69,
		Class:     "13th",
		CreatedAt: time.Now(),
		Result: result{
			totalGrade: 45.5,
			avgGrade:   23.5,
		},
	}
	fmt.Println("struct embedding ->", embeddedStruct)

	fmt.Println(student1.getClassFromStructMethod())
}

func changeOrAddFieldValue(s *Student) *Student {
	s.CreatedAt = time.Now()
	s.Age = 78
	s.Class = "3rd"
	return s
}

func printStruct(s Student) Student {
	return s
}

func (S Student) getClass() string {
	return S.Class
}

func (s Student) getClassFromStructMethod() string {
	return s.getClass()
}
