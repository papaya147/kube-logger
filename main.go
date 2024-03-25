package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/papaya147/kube-logger/config"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	options, err := config.NewEKSOptions("./options.yaml")
	if err != nil {
		panic(err)
	}

	clientset, err := config.NewEKSClientset(options)
	if err != nil {
		panic(err)
	}

	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	podName := pods.Items[0].Name
	namespace := "default"

	podLogs, err := clientset.CoreV1().Pods(namespace).GetLogs(podName, &corev1.PodLogOptions{
		SinceTime: &metav1.Time{Time: time.Now().Add(-time.Minute * 30)},
	}).DoRaw(context.Background())
	if err != nil {
		log.Fatalf("Error getting logs for pod %s: %v", podName, err)
	}

	outFile, err := os.Create("halfhour.txt")
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outFile.Close()

	_, err = outFile.Write(podLogs)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	log.Println("Logs written to out.txt successfully")
}
