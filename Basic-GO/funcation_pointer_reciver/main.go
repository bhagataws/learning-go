package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

// Pointer receiver method
func (s *Student) UpdateAge(age int) {
	s.Age = age
}

func UpdateNumber(num *int) {
	fmt.Printf("you number is %d \n", *num)
}
func main() {

	// Create student object
	student1 := Student{
		Name: "Sonu",
		Age:  25,
	}

	// Before update
	fmt.Printf("Before Update: %+v\n", student1)

	// Update age
	student1.UpdateAge(30)
	fmt.Printf("After Update: %+v\n", student1)
	x := 10
	UpdateNumber(&x)
}