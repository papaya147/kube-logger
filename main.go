package main

import (
	"github.com/papaya147/kube-logger/config"
	"github.com/papaya147/kube-logger/logs"
	"k8s.io/client-go/kubernetes"
)

func main() {
	options, err := config.NewOptions("./options.yaml", "./options.json")
	if err != nil {
		panic(err)
	}

	var clientset *kubernetes.Clientset
	switch options.ClusterProvider {
	case "eks":
		clientset, err = config.NewEKSClientset(options.EKSOptions)
		if err != nil {
			panic(err)
		}
	default:
		panic("unknown cluster provider")
	}

	if clientset == nil {
		panic("options to create clientset has not been configured")
	}

	if err := logs.Setup(options, clientset); err != nil {
		panic(err)
	}

	if err := logs.Scrape(); err != nil {
		panic(err)
	}
}
