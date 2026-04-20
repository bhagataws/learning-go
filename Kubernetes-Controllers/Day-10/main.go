package main

import "fmt"

type AppSpec struct {
	Name     string
	Replicas int
}

type FakeClient struct {
	store map[string]*AppSpec
}

func NewFakeClient() *FakeClient {
	return &FakeClient{
		store: make(map[string]*AppSpec),
	}
}

func (c *FakeClient) Get(name string, app *AppSpec) error {
	data, ok := c.store[name]
	if !ok {
		return fmt.Errorf("not found")
	}
	app.Name = data.Name
	app.Replicas = data.Replicas
	return nil
}

func (c *FakeClient) Create(app *AppSpec) error {
	fmt.Printf("Creating: %s (%d replicas)\n", app.Name, app.Replicas)
	c.store[app.Name] = app
	return nil
}

func (c *FakeClient) Update(app *AppSpec) error {
	fmt.Printf("Updating: %s (%d replicas)\n", app.Name, app.Replicas)
	c.store[app.Name] = app
	return nil
}

type AppReconciler struct {
	client *FakeClient
}

func (r *AppReconciler) Reconcile(name string) error {
	app := &AppSpec{}

	// Step 1: Get App
	err := r.client.Get(name, app)
	if err != nil {
		return err
	}

	fmt.Printf("Fetched App: %s (%d replicas)\n", app.Name, app.Replicas)

	// Step 2: Desired state
	desiredReplicas := 3

	deploymentName := app.Name + "-deployment"

	desiredDeployment := &AppSpec{
		Name:     deploymentName,
		Replicas: desiredReplicas,
	}

	// Step 3: Check if deployment exists
	existing := &AppSpec{}
	err = r.client.Get(deploymentName, existing)

	if err != nil {
		// Not found → Create
		fmt.Println("Deployment not found, creating...")
		return r.client.Create(desiredDeployment)
	}

	// Step 4: Compare and Update
	if existing.Replicas != desiredDeployment.Replicas {
		fmt.Println("Deployment exists but replicas mismatch, updating...")
		existing.Replicas = desiredDeployment.Replicas
		return r.client.Update(existing)
	}

	fmt.Println("Deployment already in desired state")
	return nil
}

func main() {
	client := NewFakeClient()

	// Seed initial App (CR)
	client.store["my-app"] = &AppSpec{
		Name:     "my-app",
		Replicas: 5,
	}

	reconciler := &AppReconciler{client: client}

	// First run → create deployment
	reconciler.Reconcile("my-app")

	// Second run → update (or no-op)
	reconciler.Reconcile("my-app")
}
