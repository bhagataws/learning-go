package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	home, _ := os.UserHomeDir()
	kubeConfigPath := filepath.Join(home, ".kube/config")

	if _, err := os.Stat(kubeConfigPath); err == nil {
		fmt.Println("Kubeconfig file exists at:", kubeConfigPath)
	} else {
		fmt.Println("Kubeconfig file does not exist at:", kubeConfigPath)
		os.Exit(1)

	}

}
