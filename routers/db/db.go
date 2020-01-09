package db

/*
 Endpoints required for the db, this might eventually need to have handlers that do the
 writing to the database.
*/

import (
	"github.com/gin-gonic/gin"
	db_service "github.com/clarencejychan/nephew-pipeline/services/db"
)

func Routes(route *gin.Engine) {
	db := route.Group("/db") 
	{
		db.GET("/", db_service.IndexHandler)
	}
}