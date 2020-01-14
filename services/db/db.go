package db

import (
	"github.com/gin-gonic/gin"
	"github.com/clarencejychan/nephew-pipeline/models"
)

func IndexHandler(db models.MongoDatastore) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message" : "Index Handler for DB Services",
		})
	}
	return gin.HandlerFunc(fn)
}