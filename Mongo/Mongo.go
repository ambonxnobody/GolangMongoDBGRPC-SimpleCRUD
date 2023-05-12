package Mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDBConfig interface {
	OpenMongoDatabaseConnection() (*mongo.Client, error)
	CloseMongoDatabaseConnection(client *mongo.Client) error
}

type mongoDBConfig struct {
	env map[string]string
}

func NewMongoDBConfig(env map[string]string) MongoDBConfig {
	return &mongoDBConfig{
		env: env,
	}
}

func (m *mongoDBConfig) OpenMongoDatabaseConnection() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	url := m.env["MONGO_DB_URL"]
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (m *mongoDBConfig) CloseMongoDatabaseConnection(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	return client.Disconnect(ctx)
}
