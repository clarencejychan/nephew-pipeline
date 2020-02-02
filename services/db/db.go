package db

import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/clarencejychan/nephew-pipeline/models"
)

func IndexHandler(db models.MongoDatastore) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message" : "Index Handler for DB Services",
		})
	}
	return gin.HandlerFunc(fn)
}

// How shoudl we query this?
func GetCommentsHandler(db models.MongoDatastore) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
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
			log.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message" : comment.Id,
		})
	}
	return gin.HandlerFunc(fn)
}