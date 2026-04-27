package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

// Method on Student struct
func (s Student) PrintDetails()  {
	fmt.Printf("Student Name: %s\n", s.Name)
	fmt.Printf("Student Age: %d\n", s.Age)
}

func main() {
	s := Student{
		Name: "Sonu",
		Age:  25,
	}
	// Call method
	s.PrintDetails()
}