package api

/*
 Endpoints required for the db, this might eventually need to have handlers that do the
 writing to the database.
*/

import (
	"github.com/gin-gonic/gin"
	api_service "github.com/clarencejychan/nephew-pipeline/services/api"
	"github.com/clarencejychan/nephew-pipeline/models"
)

func Routes(route *gin.Engine, m models.MongoDatastore) {
	api := route.Group("/api") 
	{
		api.GET("/update_all_players", api_service.Update_All_Players(m))
		api.GET("/player/:id", api_service.Get_Player(m))
	}
}