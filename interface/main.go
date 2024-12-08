package main

import "fmt"

type Student interface {
	getGrade(student) string
}
type student struct {
	name string
	age  int
	grade
}
type grade struct {
	totalGrade float32
	avgGrade   float32
}

func (g grade) getGrade(s student) string {
	return fmt.Sprintf("Student name is %s and age is %d.\nHis grade total grade is %f and avg grade is %f\n", s.name, s.age, g.totalGrade, g.avgGrade)
}

func printStudentDetail(s Student) string {
	// assert s to student struct
	return s.getGrade(s.(student))
}

func main() {
	student := student{
		name: "Wafiq",
		age:  20,
		grade: grade{
			totalGrade: 45.5,
			avgGrade:   23.5,
		},
	}

	fmt.Println(printStudentDetail(student))
}
