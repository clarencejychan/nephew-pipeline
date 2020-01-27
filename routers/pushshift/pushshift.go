package pushshift

import (
	"github.com/gin-gonic/gin"
	pushshift_service "github.com/clarencejychan/nephew-pipeline/services/pushshift"
)

func Routes(route *gin.Engine) {
	pushshift := route.Group("/pushshift")
	{
		pushshift.GET("/", pushshift_service.IndexHandler)
	}
}