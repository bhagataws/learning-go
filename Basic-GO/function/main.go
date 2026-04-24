package main

import "fmt"

func CheckEven(num int) bool {
	return num%2 == 0
}
func main() {

numbers := []int{10, 7, 22, 15, 100}

for _, num := range numbers {
	isEven := CheckEven(num)

	if isEven {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}
}

}