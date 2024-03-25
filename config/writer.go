package config

import "context"

type Writer interface {
	Open(context.Context, string) error
	Write(namespace, pod string, log []byte) error
	Close(context.Context) error
}
