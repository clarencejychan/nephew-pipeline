package aggregator

import (
	"strconv"
	"time"

	"github.com/clarencejychan/nephew-pipeline/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	layoutDaily   = "2006-02-01"
	layoutMonthly = "2006-01"
	layoutYearly  = "2006"
	dailyDB       = "aggregation-daily"
	monthlyDB     = "aggregation-monthly"
	yearlyDB      = "aggregation-yearly"
)

type AggregatorPipeline struct {
	Db *models.MongoDB
}

func (p *AggregatorPipeline) UpdateSematicScores(playerId int, comments []models.Comment) {
	updateSematicScoresToDB(p.Db, playerId, comments, layoutDaily, dailyDB)
	updateSematicScoresToDB(p.Db, playerId, comments, layoutMonthly, monthlyDB)
	updateSematicScoresToDB(p.Db, playerId, comments, layoutYearly, yearlyDB)
}

func updateSematicScoresToDB(db *models.MongoDB, playerId int, comments []models.Comment, layout string, collectionName string) {
	dailysMap := make(map[string]models.Aggregation)
	for _, comment := range comments {
		dateString := time.Unix(int64(comment.Date), 0).Format(layout)
		playerIdString := strconv.Itoa(playerId)
		c, ok := dailysMap[playerIdString+dateString]
		if ok {
			c.Semantic_Rating += comment.Semantic_Rating
			dailysMap[playerIdString+dateString] = c

		} else {
			daily := models.Aggregation{
				PlayerId:        playerId,
				Date:            dateString,
				Semantic_Rating: comment.Semantic_Rating,
			}
			dailysMap[playerIdString+dateString] = daily
		}
	}

	for _, updatedDaily := range dailysMap {
		filter := bson.D{primitive.E{Key: "playerid", Value: updatedDaily.PlayerId}, primitive.E{Key: "date", Value: updatedDaily.Date}}
		var dailyToUpdate models.Aggregation
		err := db.FindOne(collectionName, filter, &dailyToUpdate)
		if err != nil {
			db.Insert(collectionName, updatedDaily)
		} else {
			dailyToUpdate.Semantic_Rating += updatedDaily.Semantic_Rating
			db.UpdateOne(collectionName, filter, bson.M{"$set": bson.M{"semantic_rating": dailyToUpdate.Semantic_Rating}})
		}
	}
}
