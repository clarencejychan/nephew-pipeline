package db

/*
 Endpoints required for the db, this might eventually need to have handlers that do the
 writing to the database.
*/

import (
	"github.com/gin-gonic/gin"
	db_service "github.com/clarencejychan/nephew-pipeline/services/db"
	"github.com/clarencejychan/nephew-pipeline/models"
)

func Routes(route *gin.Engine, m models.MongoDatastore) {
	db := route.Group("/db") 
	{
		db.GET("/", db_service.IndexHandler(m))
		db.POST("/insert", db_service.InsertHandler(m))
	}
}