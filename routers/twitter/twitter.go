package twitter

import (
	"github.com/gin-gonic/gin"
	twitter_service "github.com/clarencejychan/nephew-pipeline/services/twitter"
	"github.com/clarencejychan/nephew-pipeline/models"
)

func Routes(route *gin.Engine, m models.MongoDatastore) {
	twitter := route.Group("/twitter")
	{
		twitter.GET("/", twitter_service.GetComment(m))
	}
}