package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LogEntry struct {
	Namespace string `bson:"namespace"`
	Pod       string `bson:"pod"`
	Log       []byte `bson:"log"`
}

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

func (m *MongoWriter) Open(uri string) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return err
	}

	m.Client = client
	m.Collection = m.Client.Database(m.DatabaseName).Collection(m.CollectionName)

	return nil
}

func (m *MongoWriter) Write(namespace string, pod string, log []byte) error {
	entry := LogEntry{
		Namespace: namespace,
		Pod:       pod,
		Log:       log,
	}

	_, err := m.Collection.InsertOne(context.Background(), entry)
	return err
}

func (m *MongoWriter) Close() error {
	return m.Client.Disconnect(context.Background())
}
