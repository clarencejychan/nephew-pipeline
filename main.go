package main

import (
	"github.com/gin-gonic/gin"
	"github.com/clarencejychan/nephew-pipeline/routers/pipeline"
	"github.com/clarencejychan/nephew-pipeline/routers/db"
)

func main() {
	router := gin.Default()

	// Router Groups
	pipeline.Routes(router)
	db.Routes(router)


	router.Run()
}