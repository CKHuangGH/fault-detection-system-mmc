package main

import (
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Printf("Error creating in-cluster config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating Kubernetes client: %v", err)
	}

	namespace, err := detectNamespace()
	if err != nil {
		fmt.Printf("Error detecting namespace: %v", err)
	}

	fmt.Printf("Current Pod's namespace: %s\n", namespace)

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: "app=vcluster",
	})
	if err != nil {
		fmt.Printf("Error listing Pods: %v", err)
	}

	fmt.Printf("Pods in the cluster:\n")
	for _, pod := range pods.Items {
		fmt.Printf("- %s\n", pod.Name)
	}

}

func detectNamespace() (string, error) {
	namespacePath := "/var/run/secrets/kubernetes.io/serviceaccount/namespace"
	nsBytes, err := os.ReadFile(namespacePath)
	if err != nil {
		return "", err
	}
	return string(nsBytes), nil
}
