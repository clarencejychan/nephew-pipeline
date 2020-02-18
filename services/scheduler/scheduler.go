package scheduler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func CreateSchedulerTask(c *gin.Context) {
		destination := c.PostForm("destination")
		time := c.PostForm("time")
		parameters := c.PostForm("parameters")
		occurence_num := c.PostForm("occurence_num")
		occurence_unit := c.PostForm("occurence_unit")

		c.JSON(http.StatusOK, gin.H{
			"destination": destination,
			"time": time,
			"parameters" : parameters,
			"occurence_num" : occurence_num,
			"occurence_unit" : occurence_unit,
	})
}