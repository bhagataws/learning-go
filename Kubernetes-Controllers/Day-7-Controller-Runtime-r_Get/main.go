package main

import "fmt"

type AppSpec struct { // Initially empty struct, with 2 fields added 1 string and 1 int
	Name     string
	Replicas int
}

type FakeClient struct{} // empty struct

func (c *FakeClient) Get(name string, app *AppSpec) error { // Geting the AppSpecm checking name not be empty and the accsing values
	// Validate input first
	if name == "" {
		return fmt.Errorf("app name not found")
	}

	// Then set values
	app.Name = name
	app.Replicas = 2

	return nil
}

func (c *FakeClient) Create(app *AppSpec) error { // Creating apps, with print statement
	// Create logic here
	fmt.Printf("Creating app: %s with %d replicas\n", app.Name, app.Replicas)
	return nil
}

func (c *FakeClient) Update(app *AppSpec) error { // Update apps, with print statement
	// Update logic here
	fmt.Printf("Updating app: %s with %d replicas\n", app.Name, app.Replicas)
	return nil
}

type AppReconciler struct { // Struct who just store address of FakeClient, So that letter can use use as client insded of FakeClient
	client *FakeClient
}

func (r *AppReconciler) Reconcile(name string) error { // Actaull reconcile, with print,
	app := &AppSpec{}

	err := r.client.Get(name, app) // Must have Name values store
	if err != nil {
		return err
	}
	fmt.Printf("Fetched app: %s with %d replicas\n", app.Name, app.Replicas)
	desiredReplicas := 3

	if app.Replicas != desiredReplicas { // Compare with Replicas(AppSpec.Replicas) to desiredReplicas
		app.Replicas = desiredReplicas

		// Step 3: Update
		return r.client.Update(app)
	}

	return nil
}

// Completely not understand
func main() {
	client := &FakeClient{}
	reconciler := &AppReconciler{client: client}

	err := reconciler.Reconcile("my-app")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

//Mostly I just copy paste stuff, I don't think as this stage I can able to write this complete logic from scratch, can you explain me the main function
//
