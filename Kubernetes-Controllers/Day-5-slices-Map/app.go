package main

import "fmt"

type AppSpec struct {
    Name       string
    Replicas   int
    Containers []string
    Labels     map[string]string
}

func (a *AppSpec) reconcile() {
    fmt.Printf("Reconciling app: %s with %d replicas\n",
        a.Name, a.Replicas)

    fmt.Println("Labels:")
    for key, value := range a.Labels {
        fmt.Printf("  %s = %s\n", key, value)
    }

    fmt.Println("Containers:")
    for _, c := range a.Containers {
        fmt.Printf("  %s \n", c)
    }

    for i := 1; i <= a.Replicas; i++ {
        fmt.Printf("Ensuring pod %d is running with %d containers\n",
            i, len(a.Containers))
    }
}