package workloadExplore

import (
	"context"
	"flag"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func ExplorePods() {
	fmt.Println("Exploring pods in cluster")

	kubeconfig := flag.String("kubeconfig", "C:\\Users\\Ananya\\.kube\\config", "kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("The kubeconfig cannot be loaded: %v\n", err)
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("The clientset cannot be built: %v\n", err)
		os.Exit(1)
	}

	pod, err := clientset.CoreV1().Pods("default").Get(context.Background(), "api-gateway-6cd9d594fb-pq7hh", metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Could not fetch pod: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(pod.Labels["app"])
}
