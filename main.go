package main

import (
	"github.com/clarencejychan/nephew-pipeline/models"
	api_routes "github.com/clarencejychan/nephew-pipeline/routers/api"
	db_routes "github.com/clarencejychan/nephew-pipeline/routers/db"
	pipeline_routes "github.com/clarencejychan/nephew-pipeline/routers/pipeline"
	pushshift_routes "github.com/clarencejychan/nephew-pipeline/routers/pushshift"
	scheduler_routes "github.com/clarencejychan/nephew-pipeline/routers/scheduler"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"path/filepath"
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

	db, err := models.NewDB()
	// Eventually need to set-up a way to retry the server connection.
	if err != nil {
		log.Println(err.Error())
	}

	// example db insert:
	// 		collection: 	the collection name
	// 		obj: 			any db interface object
	// err = db.Insert(collection, obj)

	router := gin.Default()

	// Router Groups
	pipeline_routes.Routes(router, db)
	db_routes.Routes(router, db)
	api_routes.Routes(router, db)
	pushshift_routes.Routes(router, db)
	scheduler_routes.Routes(router, db)

	router.Run()
}
