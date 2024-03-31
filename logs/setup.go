package logs

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/papaya147/kube-logger/config"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var namespaces []string
var loggers []config.Writer
var clientset *kubernetes.Clientset

type NamespacedPod struct {
	Name      string
	Namespace string
}

var pods []NamespacedPod

func Setup(options *config.Options, cs *kubernetes.Clientset) error {
	if len(options.Namespaces) == 0 {
		return errors.New("namespaces cannot be empty")
	}
	namespaces = options.Namespaces

	if options.ConsoleOptions != nil && options.ConsoleOptions.Active {
		loggers = append(loggers, config.NewConsoleWriter())
	}

	if options.MongoOptions != nil && options.MongoOptions.Active {
		logger := config.NewMongoWriter(options.MongoOptions.Database, options.MongoOptions.Collection)
		if err := logger.Open(context.Background(), options.MongoOptions.ConnectionURI); err != nil {
			return err
		}
		loggers = append(loggers, logger)
	}

	if options.ElasticsearchOptions != nil && options.ElasticsearchOptions.Active {
		logger := config.NewElasticsearchWriter(
			options.ElasticsearchOptions.Host,
			options.ElasticsearchOptions.Username,
			options.ElasticsearchOptions.Password,
			options.ElasticsearchOptions.Index,
		)
		if err := logger.Open(context.Background(), ""); err != nil {
			return err
		}
		loggers = append(loggers, logger)
	}

	clientset = cs

	if err := loadClusterPods(options); err != nil {
		return err
	}

	fmt.Println("Scraping logs from:")
	for _, pod := range pods {
		fmt.Printf("%s - %s\n", pod.Namespace, pod.Name)
	}
	fmt.Println()

	return nil
}

func loadClusterPods(options *config.Options) error {
	for _, namespace := range namespaces {
		namespacePods, err := clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			return err
		}
		for _, pod := range namespacePods.Items {
			if len(options.PodPrefixes) == 0 {
				pods = append(pods, NamespacedPod{Name: pod.Name, Namespace: namespace})
			} else {
				for _, prefix := range options.PodPrefixes {
					if strings.HasPrefix(pod.Name, prefix) {
						pods = append(pods, NamespacedPod{Name: pod.Name, Namespace: namespace})
						break
					}
				}
			}
		}
	}
	return nil
}

func write(namespace, pod string, data []byte) error {
	for _, logger := range loggers {
		if err := logger.Write(namespace, pod, data); err != nil {
			return err
		}
	}
	return nil
}
