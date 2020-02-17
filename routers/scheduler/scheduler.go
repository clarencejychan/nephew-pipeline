package scheduler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	scheduler_service "github.com/clarencejychan/nephew-pipeline/services/scheduler"
)

func Routes(route *gin.Engine) {
	scheduler := route.Group("/scheduler") 
	{
		route.LoadHTMLGlob("templates/*")
		scheduler.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
		scheduler.POST("/create", scheduler_service.CreateSchedulerTask)
	}
}