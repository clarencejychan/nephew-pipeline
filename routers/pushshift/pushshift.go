package pushshift

import (
	"github.com/gin-gonic/gin"
	pushshift_service "github.com/clarencejychan/nephew-pipeline/services/pushshift"
	"github.com/clarencejychan/nephew-pipeline/models"
)

func Routes(route *gin.Engine, m models.MongoDatastore) {
	pushshift := route.Group("/pushshift")
	{
		pushshift.GET("/", pushshift_service.GetComment(m))
	}
}