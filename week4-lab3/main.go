package main

import (
	"errors"
	"fmt"
)

type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Year  int     `json:"year"`
	GPA   float64 `json:"gpa"`
}

func (s *Student) IsHonor() bool {
	return s.GPA >= 3.50
}
func (s *Student) Validte() error {
	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.GPA < 0 || s.GPA > 4 {
		return errors.New("gpa must be between 0-4")
	}
	if s.Year < 1 || s.Year > 4 {
		return errors.New("year must be between 1-4")
	}
	return nil
}

func main() {
	// var st Student = Student{ID: "1", Name: "Thanaphat", Email: "Sukchuen_t@su.ac.th", Year: 2, GPA: 3.75}
	students := []Student{
		{ID: "1", Name: "Thanaphat", Email: "Sukchuen_t@su.ac.th" , Year: 2, GPA: 3.75},
		{ID: "2", Name: "John Doe", Email: "john@example.com", Year: 3, GPA: 3.80},
	}
	newStudent := Student{ID: "3", Name: "Jane Smith", Email:"john@samit", Year: 1, GPA: 3.90}
	students = append(students, newStudent)

	for i,student := range students{
		fmt.Printf("%d Honor = %v\n",i,student.IsHonor())
		fmt.Printf("%d Validate = %v\n",i, student.Validte())
	}
}
