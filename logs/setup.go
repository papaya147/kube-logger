package logs

import (
	"errors"

	"github.com/papaya147/kube-logger/config"
	"k8s.io/client-go/kubernetes"
)

var namespaces []string
var logger config.Writer
var clientset *kubernetes.Clientset

func Setup(options *config.Options, cs *kubernetes.Clientset) error {
	if len(options.Namespaces) == 0 {
		return errors.New("namespaces cannot be empty")
	}
	namespaces = options.Namespaces

	switch options.Writer {
	case "console":
		logger = config.NewConsoleWriter()
	default:
		return errors.New("unknown writer")
	}

	clientset = cs

	return nil
}
