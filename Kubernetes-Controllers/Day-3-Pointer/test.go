package main

import "fmt"

type AppSpec struct {
	Name     string
	Replicas int
	Image    string
	State    string
}

func main() {
	app := &AppSpec{
		Name:     "my-app",
		Replicas: 3,
		Image:    "nginx:latest",
		State:    "running",
	}

	fmt.Printf("App: %s, Replicas: %d, Image: %s, State: %s\n",
		app.Name, app.Replicas, app.Image, app.State)

	// Take pointer of replicas
	replicaPtr := &app.Replicas

	// Modify using pointer
	*replicaPtr = 5

	fmt.Printf("After update: %d replicas\n", app.Replicas)

	test()
}
