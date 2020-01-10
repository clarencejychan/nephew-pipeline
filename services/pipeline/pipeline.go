package pipeline

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func IndexHandler(c *gin.Context) {
	fmt.Print("hello")
	c.JSON(200, gin.H{
		"message" : "pong",
	})
}