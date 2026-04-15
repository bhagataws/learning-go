package main

import "fmt"

type AppSpec struct {
	Name     string `json:"name"`
	Replicas int    `json:"replicas"`
	Image    string `json:"image"`
	State    string `json:"state"`
}

func main() {
	app := AppSpec{
		Name:     "my-app",
		Replicas: 3,
		Image:    "nginx:latest",
		State:    "running",
	}

	// Use the app variable as needed
	fmt.Printf("%s Deployment is deployed with %d which has image: %s and state: %s \n", app.Name, app.Replicas, app.Image, app.State)

	for i := 1; i <= app.Replicas; i++ {
		fmt.Printf("Create pods of %s-%d\n", app.Name, i)
	}

}
