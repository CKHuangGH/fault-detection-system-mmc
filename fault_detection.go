package main

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// 使用 InClusterConfig() 函數建立集群內部的配置
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Error creating in-cluster config: %v", err)
	}

	// 建立 Kubernetes 客戶端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	// 在這裡您可以使用 clientset 進行各種 Kubernetes 操作
	// 例如列出 Pods，創建 Deployments，查詢 Services 等

	// 這裡是一個簡單的示例，列出集群中的 Pods
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing Pods: %v", err)
	}

	fmt.Printf("Pods in the cluster:\n")
	for _, pod := range pods.Items {
		fmt.Printf("- %s\n", pod.Name)
	}
}
