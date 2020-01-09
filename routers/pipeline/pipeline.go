package pipeline

// Endpoints that are required for the pipeline

import (
	"github.com/gin-gonic/gin"
	pipeline_service "github.com/clarencejychan/nephew-pipeline/services/pipeline"
)

func Routes(route *gin.Engine) {
	pipeline := route.Group("/pipeline") 
	{
		pipeline.GET("/", pipeline_service.IndexHandler)
	}
}