package config

import (
	"context"
	"regexp"
)

type LogEntry struct {
	Timestamp int64  `bson:"timestamp"`
	Namespace string `bson:"namespace"`
	Pod       string `bson:"pod"`
	Log       string `bson:"log"`
}

var escapePattern = "\x1b[^m]*m"
var escapeRegex = regexp.MustCompile(escapePattern)

type Writer interface {
	Open(context.Context, string) error
	Write(namespace, pod string, log []byte) error
	Close(context.Context) error
}
