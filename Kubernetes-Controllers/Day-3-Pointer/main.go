package main

import "fmt"

type AppSpec1 struct {
	Name     string `json:"name"`
	Replicas int    `json:"replicas"`
	Image    string `json:"image"`
	State    string `json:"state"`
}

func test() {
	app := &AppSpec1{
		Name:     "my-app",
		Replicas: 3,
		Image:    "nginx:latest",
		State:    "running",
	}

	// Use the app variable as needed
	fmt.Printf("%s Deployment is deployed with %d which has image: %s and state: %s \n", app.Name, app.Replicas, app.Image, app.State)
	ptras := app.Replicas
	fmt.Printf("Pointer location: %p\n", &ptras)
	fmt.Printf("Pointer value: %d\n", ptras)
}
