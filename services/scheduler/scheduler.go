package scheduler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateSchedulerTask(c *gin.Context) {
		destination := c.PostForm("destination")
		time := c.PostForm("time")
		parameters := c.PostForm("parameters")
		occurence_num := c.PostForm("occurence_num")
		occurence_unit := c.PostForm("occurence_unit")

		fmt.Println("I ran")
		fmt.Println(destination)
		fmt.Println(time)
		fmt.Println(parameters)
		fmt.Println(occurence_num)
		fmt.Println(occurence_unit)
}