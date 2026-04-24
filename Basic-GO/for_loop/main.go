package main

import "fmt"

func main() {
 
    // slice
	sliceData := []int{1,2,3,4,5,6,7}
    sum := 0

	for i, v := range(sliceData){
		fmt.Println("Day", i ,"views",v)
		sum = sum + v
	}
	fmt.Println(sum)
}