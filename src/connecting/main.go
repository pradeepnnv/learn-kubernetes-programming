package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeConfig *string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}
	kubeConfig = flag.String("kubeconfig", homeDir+"/.kube/config", "Path to the kubeconfig file to use for CLI requests")
}

func main() {
	flag.Parse()

	if *kubeConfig == "" {
		flag.Usage()
		os.Exit(1)
	}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		panic(err.Error())
	}

	clientset := kubernetes.NewForConfigOrDie(config)
	serverVersion, err := clientset.Discovery().ServerVersion()

	if err != nil {
		panic(err.Error())
	}

	log.Printf("Cluster is running k8s version %s\n", serverVersion)

	// list all namespaces
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, n := range namespaces.Items {
		fmt.Println(n.ObjectMeta.Name)
	}
}
