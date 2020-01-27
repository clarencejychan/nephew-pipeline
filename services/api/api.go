package api

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/clarencejychan/nephew-pipeline/models"
)

func Update_All_Players(db models.MongoDatastore) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		url := "http://localhost:3000/api/get-all-current-players"

		resp, _ := http.Get(url)
		// defers closing the response body until end of function, prevents resource leaks
		defer resp.Body.Close()
		body, _:= ioutil.ReadAll(resp.Body)
		
		fmt.Printf(string(body))
		c.JSON(200, resp.Body)

	}
	return gin.HandlerFunc(fn)
}

func Get_Player(db models.MongoDatastore) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message" : "Requested" + id, 
		})
	}
	return gin.HandlerFunc(fn)
}