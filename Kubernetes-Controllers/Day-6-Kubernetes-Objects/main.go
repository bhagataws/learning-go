package main

import (
    "fmt"

    appsv1 "k8s.io/api/apps/v1"
    corev1 "k8s.io/api/core/v1"
)

func main() {
    replicas := int32(2)

    deployment := appsv1.Deployment{
        Spec: appsv1.DeploymentSpec{
            Replicas: &replicas,
            Template: corev1.PodTemplateSpec{
                Spec: corev1.PodSpec{
                    Containers: []corev1.Container{
                        {
                            Name:  "nginx",
                            Image: "nginx:latest",
                        },
                        {
                            Name:  "redis",
                            Image: "redis:latest",
                        },
                    },
                },
            },
        },
    }

    fmt.Printf("Deployment with %d replicas\n", *deployment.Spec.Replicas)

    for _, c := range deployment.Spec.Template.Spec.Containers {
        fmt.Printf("Container: %s Image: %s\n", c.Name, c.Image)
    }
}