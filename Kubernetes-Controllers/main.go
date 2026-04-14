package main

import "fmt"

func main() {
	fmt.Printf("****Welcome to the Kubernetes controllers application****\n")
	replicas := 1
	applicationName := "K8s-controllers"

	if replicas == 0 {
		fmt.Printf("No pods to create \n")

	} else {
		fmt.Printf("Deploying %s \n", applicationName)
		for i := 1; i <= replicas; i++ {
			fmt.Printf("Starting %s pods %d\n", applicationName, i)
		}
	}
}
