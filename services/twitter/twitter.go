package twitter

import "github.com/gin-gonic/gin"
import "log"
import "os"
import "github.com/clarencejychan/nephew-pipeline/models"
import "github.com/dghubble/go-twitter/twitter"
import "golang.org/x/oauth2"
import "golang.org/x/oauth2/clientcredentials"
import "time"

type AnalysisResponse struct {
	Semantic_Rating		float64
}

func SetUpClient() *twitter.Client {
	// Load environment variables (.env) Might not be needed since main.go already sets Env variables???
	//	err := gotenv.Load("../../.env")
		
	// oauth2 configures a client that uses app credentials to keep a fresh token
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("TWITTER_CONSUMER_KEY"),
		ClientSecret: os.Getenv("TWITTER_CONSUMER_SECRET"),
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)

	// Twitter client
	client := twitter.NewClient(httpClient)
	return client
}

func GetComment(db models.MongoDatastore) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		comments, newSinceId := getTwitterCommentsForPlayer("James Harden", 0) // SinceID will need to be pulled from the Task. Something to store saying this was the last tweet read
		
		// get analysis result on each comment
		var commentsWithRating []models.Comment
		for i := range comments {
			//commentWithRating := getAnalysisResult(comments[i])
			// Temp until the analysis is done being merged in and settled
			comments[i].Semantic_Rating = 10
			commentWithRating := comments[i]
			commentsWithRating = append(commentsWithRating, commentWithRating)
		}

		// Write to db and set task to have the next Sinceid
		c.JSON(200, gin.H{
		"Lebron" : "Goat",
		"Harden" : newSinceId,
		})

	}
	return gin.HandlerFunc(fn)
}

// Currently we aim to just get any mentions of the player in question
func getTwitterCommentsForPlayer(query string, sinceId int64) ([]models.Comment, int64) {
	client := SetUpClient()
	var comments []models.Comment
	var maxId int64 = 0
	var newSinceId int64 = 0

	log.info("Query: %s\n. SinceId: %d", query, sinceId))
	for {
		search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
			Query: query,
			ResultType: "recent",
			Count: 100,
			SinceID: sinceId,
			MaxID: maxId,
		})

		if err != nil {
			log.Println(err.Error())
		}

		len := len(search.Statuses)

		var i int = 0
		for i < len {
			tweet := search.Statuses[i]
			
			// Set this so we know for the next run
			if newSinceId == 0 {
				newSinceId = tweet.ID
			}

			// Struct -> https://godoc.org/github.com/dghubble/go-twitter/twitter#Tweet
			var comment models.Comment

			maxId = tweet.ID - 1
			comment.Id = tweet.IDStr
			comment.Comment = tweet.Text
			comment.Source = "Twitter"
			epochTime, _ := time.Parse("Mon Jan 02 15:04:05 -0700 2006", tweet.CreatedAt)
			comment.Date = uint(epochTime.Unix())
			comment.Author = tweet.User.Name
			comment.Parent = tweet.InReplyToStatusIDStr
			//comment.Children = If children is just reply count, it's part of metadata
			comment.Subject = query

			// comment.Metadata = Later if there's a need for Metadata
			comments = append(comments,comment)
		}

		// We reached the end so break
		if len != 100 {
			break
		}
	}
	return comments, newSinceId
}