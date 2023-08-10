package main

import (
	"context"
	"fmt"
	"log"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func getNamespace() string {
	// 獲取 Pod 的命名空間（該信息在 Pod 中的環境變數中）
	namespace := os.Getenv("NAMESPACE")
	if namespace == "" {
		// 如果未能獲取到命名空間，則預設為 "default"
		namespace = "default"
	}
	return namespace
}

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Error creating in-cluster config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	namespace := getNamespace()

	fmt.Printf("Listing Pods in namespace: %s\n", namespace)

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing Pods: %v", err)
	}

	fmt.Printf("Pods in the cluster:\n")
	for _, pod := range pods.Items {
		fmt.Printf("- %s\n", pod.Name)
	}
}
