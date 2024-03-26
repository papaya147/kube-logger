package logs

import (
	"context"
	"errors"

	"github.com/papaya147/kube-logger/config"
	"k8s.io/client-go/kubernetes"
)

var namespaces []string
var loggers []config.Writer
var clientset *kubernetes.Clientset

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
