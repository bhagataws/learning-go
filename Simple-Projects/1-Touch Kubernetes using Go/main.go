package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	home, _ := os.UserHomeDir()
	defaultKubeconfig := filepath.Join(home, ".kube", "config")

	kubeconfigEnv := os.Getenv("KUBECONFIG")

	kubeconfig := defaultKubeconfig
	if kubeconfigEnv != "" {
		kubeconfig = kubeconfigEnv
	}
	println("Using kubeconfig:", kubeconfig)

	// 🔹 Build config from kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	// 🔹 Create Kubernetes client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// 🔹 List Pods (default namespace)
	pods, err := clientset.CoreV1().Pods("local-path-storage").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	// List the Namespace
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Pods in default namespace:")
	for _, pod := range pods.Items {
		fmt.Printf("- %s (%s) - %s\n", pod.Name, pod.Status.Phase, pod.Namespace)
	}

	fmt.Println("Namespaces:")
	for _, ns := range namespaces.Items {
		fmt.Printf("- %s\n", ns.Name)
	}
}
