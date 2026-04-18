package main

import "fmt"

type AppSpec struct {
    Name     string
    Replicas int
}

type FakeClient struct{}

func (c *FakeClient) Get(name string, app *AppSpec) error {
    if name == "" {
        return fmt.Errorf("app not found")
    }

    app.Name = name
    app.Replicas = 2
    return nil
}

func(c *FakeClient) Create(app *AppSpec) error {
  fmt.Printf("Create We app")
}

func main() {
 

}