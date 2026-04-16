package main

import "fmt"

type AppSpec struct {
	Name     string
	Replicas int
}

// reconcile method
func (a *AppSpec) reconcile() error {
	if a.Name == "" {
		return fmt.Errorf("app name can't be empty")
	}

	fmt.Printf("Reconciling application: %s\n", a.Name)

	// scaling logic inside reconcile (better design)
	if a.Replicas > 3 {
		fmt.Println("Scaling application")
	} else {
		fmt.Println("Normal deployment")
	}

	for i := 1; i <= a.Replicas; i++ {
		fmt.Printf("Ensuring pod %d is running\n", i)
	}

	return nil
}

func main() {
	app := &AppSpec{
		Name:     "guestbook",
		Replicas: 4,
	}

	err := app.reconcile()
	if err != nil {
		fmt.Println("Error:", err)
	}
}