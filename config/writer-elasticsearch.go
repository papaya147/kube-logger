package config

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticsearchWriter struct {
	Host     string `yaml:"host" json:"host"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Index    string `yaml:"index" json:"index"`
	Client   *elasticsearch.Client
}

func NewElasticsearchWriter(host, username, password, index string) Writer {
	return &ElasticsearchWriter{
		Host:     host,
		Username: username,
		Password: password,
		Index:    index,
	}
}

func (e *ElasticsearchWriter) Open(context.Context, string) error {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{e.Host},
		Username:  e.Username,
		Password:  e.Password,
	})
	if err != nil {
		return err
	}

	if _, err := client.Ping(); err != nil {
		return err
	}

	if _, err = client.Indices.Create(e.Index); err != nil {
		return err
	}

	e.Client = client
	return nil
}

func (e *ElasticsearchWriter) Write(namespace string, pod string, log []byte) error {
	logWithoutEscape := escapeRegex.ReplaceAllString(string(log), "")

	entry := LogEntry{
		Timestamp: time.Now().UnixNano(),
		Namespace: namespace,
		Pod:       pod,
		Log:       logWithoutEscape,
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	_, err = e.Client.Index(e.Index, bytes.NewReader(data))
	return err
}

func (e *ElasticsearchWriter) Close(context.Context) error {
	return nil
}
