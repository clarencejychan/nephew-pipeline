package models

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatastore interface {
	Close() error
}

type MongoDB struct {
	Client *mongo.Client
}

func NewDB() (*MongoDB, error) {
	// Init the database conector
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	
	// WithTimeout creates a goroutine that is retained unless cancel is called. 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// cancels as soon as we exit out of the NewDB function.
	defer cancel()
	err = client.Connect(ctx)

	return &MongoDB{client}, err
}

func (m *MongoDB) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := m.Client.Disconnect(ctx)
	return err
}