package main

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/clarencejychan/nephew-pipeline/models"
	api_routes "github.com/clarencejychan/nephew-pipeline/routers/api"
	db_routes "github.com/clarencejychan/nephew-pipeline/routers/db"
	"github.com/gin-gonic/gin"

	"github.com/clarencejychan/nephew-pipeline/services/pipelines/reddit"
)

func main() {
	// Initialize Logging.
	absPath, err := filepath.Abs("./logs")
	if err != nil {
		panic(err)
	}

	f, _ := os.Create(absPath + "/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	app_log, err := os.OpenFile(absPath+"/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(app_log)
	defer app_log.Close()

	// Initialize Database
	db, err := models.NewDB()
	// Eventually need to set-up a way to retry the server connection.
	if err != nil {
		log.Println(err.Error())
	}

	// Initialize the pipelines
	_ = reddit.New(db)

	// example db insert:
	// 		collection: 	the collection name
	// 		obj: 			any db interface object
	// err = db.Insert(collection, obj)

	router := gin.Default()

	// Router Groups
	db_routes.Routes(router, db)
	api_routes.Routes(router, db)

	router.Run()
}
