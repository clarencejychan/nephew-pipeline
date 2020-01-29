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

// How shoudl we query this?
func GetCommentsHandler(db models.MongoDatastore) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message" : "PlaceHolder",
		})
	}
	return gin.HandlerFunc(fn)
}

func InsertHandler(db models.MongoDatastore) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var comment models.Comment

		err := c.BindJSON(&comment)
		err = db.Insert("collection1", &comment)

		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message" : comment.Id,
		})
	}
	return gin.HandlerFunc(fn)
}