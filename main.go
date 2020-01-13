package main

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gin-gonic/gin"
	pipeline_routes "github.com/clarencejychan/nephew-pipeline/routers/pipeline"
	db_routes "github.com/clarencejychan/nephew-pipeline/routers/db"
)

func main() {
	router := gin.Default()

	// Router Groups
	pipeline_routes.Routes(router)
	db_routes.Routes(router)


	// Init the database conector
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	router.Run()
}