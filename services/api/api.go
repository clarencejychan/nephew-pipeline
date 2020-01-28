package api

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/clarencejychan/nephew-pipeline/models"
)

type PlayersResponse struct {
	Players []models.Player
}


func Update_All_Players(db models.MongoDatastore) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		url := "http://localhost:3000/api/get-all-current-players"

		resp, err := http.Get(url)
		
		if err != nil {
			c.JSON(400, gin.H{
				"error" : err.Error(), 
			})
		}
		
		// defers closing the response body until end of function, prevents resource leaks
		defer resp.Body.Close()
		body, _:= ioutil.ReadAll(resp.Body)

		var query PlayersResponse
		err = json.Unmarshal(body, &query.Players)
		if err != nil {
			c.JSON(400, gin.H{
				"error" : err.Error(),
			})
		}
		
		// Bulk Insert here
		b := make([]interface{}, len(query.Players))

		for i:= range query.Players {
			b[i] = query.Players[i]
		}

		err = db.BulkInsert("collection1", b)
		if err != nil {
			c.JSON(400, gin.H{
				"error" : err.Error(),
			})
		}
	
		c.JSON(200, gin.H{
			"inserted":string(body),
		})
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