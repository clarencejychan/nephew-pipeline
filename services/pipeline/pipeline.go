package pipeline

import (
	"net/http"
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