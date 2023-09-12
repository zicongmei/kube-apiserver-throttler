package main

import (
	"context"
	"log"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	interval = 100 * time.Millisecond
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	ctx := context.Background()
	nsb, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		panic(err.Error())
	}
	ns := string(nsb)
	for {
		go apiserverCall(ns, clientset, ctx)
		time.Sleep(interval)
	}
}

func apiserverCall(
	namespace string,
	clientset *kubernetes.Clientset,
	ctx context.Context) {
	start := time.Now()
	_, err := clientset.CoreV1().ServiceAccounts(namespace).List(ctx, metav1.ListOptions{})
	finish := time.Now()
	if err != nil {
		log.Printf("error after %s: %v", finish.Sub(start).String(), err)
	}
	log.Printf("Success call after %s.", finish.Sub(start).String())
}
