package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	pipeline_routes 			"github.com/clarencejychan/nephew-pipeline/routers/pipeline"
	db_routes 					"github.com/clarencejychan/nephew-pipeline/routers/db"
	"github.com/clarencejychan/nephew-pipeline/models"
)

func main() {
	db, err := models.NewDB()
	// Eventually need to set-up a way to retry the server connection.
	if err != nil {
		fmt.Println(err.Error())
	}

	// example db insert:
	// 		collection: 	the collection name 
	// 		obj: 			any db interface object
	// err = db.Insert(collection, obj)

	router := gin.Default()

	// Router Groups
	pipeline_routes.Routes(router, db)
	db_routes.Routes(router, db)

	router.Run()
}