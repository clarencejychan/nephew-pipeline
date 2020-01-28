package models

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/subosito/gotenv"
	"os"
	"fmt"
)

type MongoDatastore interface {
	Close() error
	Insert(string, interface{}) error
	BulkInsert(string, []interface{}) error
}

type MongoDB struct {
	Client *mongo.Client
}

func NewDB() (*MongoDB, error) {
	// Load environment variables (.env)
	err := gotenv.Load()

	// Build MongoDB Atlas URL
	dbURL := fmt.Sprintf(
		"mongodb+srv://%s:%s@cluster0-2banx.mongodb.net/test?retryWrites=true&w=majority",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"))

	// Init the database conector
	client, err := mongo.NewClient(options.Client().ApplyURI(dbURL))
	
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

func (m *MongoDB) BulkInsert(c string, d []interface{}) error {
	collection := m.Client.Database("DB1").Collection(c)

	insertResult, err := collection.InsertMany(context.TODO(), d)
	fmt.Println("Inserted bulk documents: ", insertResult.InsertedIDs)
	return err;
}



func (m *MongoDB) Insert(c string, d interface{}) error {
	// Default to sandbox db for now with the name DB1, change when moving to prod
	collection := m.Client.Database("DB1").Collection(c)

	insertResult, err := collection.InsertOne(context.TODO(), d)

	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
	return err
}