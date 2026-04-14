package main

import "fmt"

func main() {
	fmt.Printf("****Welcome to the Kubernetes controllers application****\n")
	replicas := 3
	applicationName := "K8s-controllers"

	fmt.Printf("Deploying %v \n", applicationName)
	for i := 1; i <= replicas; i++ {
		fmt.Printf("Starting %v pods %d\n", applicationName, i)
	}
}
