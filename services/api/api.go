package api

import (
	"context"
	"strings"
	"net/http"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/clarencejychan/nephew-pipeline/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func Get_Player_By_ID(db models.MongoDatastore) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(400, gin.H{
				"error" : err.Error(),
			})
		}

		filter := bson.D{primitive.E{Key: "id", Value: id}}

		var queryResult models.Player

		err = db.FindOne("collection1", filter, &queryResult)

		if err != nil {
			c.JSON(400, gin.H{
				"error" : err.Error(),
			})
		}

		result, err := json.Marshal(queryResult)

		if err != nil {
			c.JSON(400, gin.H{
				"error" : err.Error(),
			})
		}

		c.JSON(200, gin.H{
			"message" : string(result),
		})
	}
	return gin.HandlerFunc(fn)
}

func Get_Player_By_Name(db models.MongoDatastore) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		name := strings.Title(strings.ToLower(c.Param("name")))
		
		filter := bson.D{
			{"$or",
				bson.A{
					bson.D{{"first_name", name}},
					bson.D{{"last_name", name}},
				},
			},
		}
		
		//filter := bson.D{primitive.E{Key: "first_name", Value: name}}

		o := options.Find()
		o.SetLimit(10)

		cursor, err := db.FindAll("collection1", o, filter)

		if err != nil {
			c.JSON(400, gin.H{
				"error" : err.Error(),
			})
		}

		var results []models.Player

		for cursor.Next(context.Background()) {
			var player models.Player
			err := cursor.Decode(&player)
			if err != nil {
				c.JSON(400, gin.H{
					"error" : err.Error(),
				})
			}

			results = append(results, player)
		}

		cursor.Close(context.Background());

		if err != nil {
			c.JSON(400, gin.H{
				"error" : err.Error(),
			})
		}

		result, err := json.Marshal(results)

		c.JSON(200, gin.H{
			"message" : string(result),
		})
	}
	return gin.HandlerFunc(fn)
}