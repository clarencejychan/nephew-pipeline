package aggregator

import (
	"github.com/clarencejychan/nephew-pipeline/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func updateSematicScores(db models.MongoDatastore, playerId int, comments []models.Comment) {
	//Daily Aggregating
	var dailys map[uint]models.Daily
	for _, comment := range comments {
		c, ok := dailys[comment.Date]
		if ok {
			c.Semantic_Rating += comment.Semantic_Rating

		} else {
			daily := models.Daily{
				PlayerId:        playerId,
				Date:            comment.Date,
				Semantic_Rating: comment.Semantic_Rating,
			}
			dailys[comment.Date] = daily
		}
	}

	for _, updatedDaily := range dailys {
		filter := bson.D{primitive.E{Key: "playerId", Value: updatedDaily.PlayerId}, primitive.E{Key: "date", Value: updatedDaily.Date}}
		var dailyToUpdate models.Daily
		err := db.FindOne("aggregation-daily", filter, &dailyToUpdate)
		if err != nil {
			db.Insert("aggregation-daily", updatedDaily)
		} else {
			dailyToUpdate.Semantic_Rating += updatedDaily.Semantic_Rating
			db.UpdateOne("aggregation-daily", filter, dailyToUpdate)
		}
	}
}
