package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoWriter struct {
	ConnectionURI  string `yaml:"connection_uri" json:"connection_uri"`
	DatabaseName   string `yaml:"database" json:"database"`
	CollectionName string `yaml:"collection" json:"collection"`
	Collection     *mongo.Collection
	Client         *mongo.Client
}

func NewMongoWriter(database, collection string) Writer {
	return &MongoWriter{
		DatabaseName:   database,
		CollectionName: collection,
	}
}

func (m *MongoWriter) Open(ctx context.Context, uri string) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	m.Client = client
	m.Collection = m.Client.Database(m.DatabaseName).Collection(m.CollectionName)

	return nil
}

func (m *MongoWriter) Write(namespace string, pod string, log []byte) error {
	logWithoutEscape := escapeRegex.ReplaceAllString(string(log), "")

	entry := LogEntry{
		Nanos:     time.Now().UnixNano(),
		Namespace: namespace,
		Pod:       pod,
		Log:       logWithoutEscape,
	}

	_, err := m.Collection.InsertOne(context.Background(), entry)
	return err
}

func (m *MongoWriter) Close(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}
