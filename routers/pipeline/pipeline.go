package pipeline

// Endpoints that are required for the pipeline

import (
	"github.com/gin-gonic/gin"
	pipeline_service "github.com/clarencejychan/nephew-pipeline/services/pipeline"
	"github.com/clarencejychan/nephew-pipeline/models"
)

func Routes(route *gin.Engine, m models.MongoDatastore) {
	pipeline := route.Group("/pipeline") 
	{
		pipeline.GET("/", pipeline_service.IndexHandler(m))
	}
}