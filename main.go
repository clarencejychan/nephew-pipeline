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
	"github.com/robfig/cron/v3"

	"github.com/clarencejychan/nephew-pipeline/services/pipelines/reddit"
	// "github.com/clarencejychan/nephew-pipeline/services/pipelines/twitter"
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
	reddit_pipeline := reddit.New(db)
	test_reddit := map[string]string{
		"subject": "Harden",
		"after": "1d",
		"before": "0d",
		"subreddit": "nba",
	}

	// twitter_pipeline := twitter.New(db)
	// test_twitter := map[string]string{}

	c := cron.New()
	// "0 0 * * *" == everyday @ midnight
	c.AddFunc("0 0 * * *", func() {reddit_pipeline.Run(test_reddit)})
	// c.AddFunc("50 * * * *", twitter_pipeline.Run(test))
	c.Start()

	router := gin.Default()

	// Router Groups
	db_routes.Routes(router, db)
	api_routes.Routes(router, db)

	router.Run()
}
