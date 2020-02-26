package twitter

import "log"
import "os"
import "github.com/clarencejychan/nephew-pipeline/models"
import "github.com/dghubble/go-twitter/twitter"
import "golang.org/x/oauth2"
import "golang.org/x/oauth2/clientcredentials"
import "time"
import "github.com/clarencejychan/nephew-pipeline/services/pipelines"

type TwitterPipeline struct {
	db *models.MongoDB
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

func New(db *models.MongoDB) models.Pipeline {
	return &TwitterPipeline{db: db}
}

func (p *TwitterPipeline) getComment(params map[string]string) (int64, error) {
	// params["sinceId"] instead of 0 when ready
	// params["query"] etc
	comments, newSinceId, err := getTwitterCommentsForPlayer("James Harden", 0) 

	// get analysis result on each comment
	analysisReq := pipelines.AnalysisRequest{
		PlayerId: 123,
		Comments: comments,
	}

	resp, err := pipelines.GetAnalysisResult(analysisReq)

	// Necessary, maybe we should put this in the bulk insert function
	x := make([]interface{}, len(resp.Comments))
	for i := range resp.Comments {
		x[i] = resp.Comments[i]
	}
	p.db.BulkInsert("comments", x)

	return newSinceId, err
}

// Currently we aim to just get any mentions of the player in question
func getTwitterCommentsForPlayer(query string, sinceId int64) ([]models.Comment, int64, error) {
	client := SetUpClient()
	var comments []models.Comment
	var maxId int64 = 0
	var newSinceId int64 = 0
	// 450 reqs per 15 min window = 0.5 request a second
	rate:= time.Second / 2
	throttle := time.Tick(rate)

	log.Println("Query: %s\n. SinceId: %d", query, sinceId)
	for {
		<- throttle
		search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
			Query: query,
			ResultType: "recent",
			Count: 100,
			SinceID: sinceId,
			MaxID: maxId,
		})

		if err != nil {
			log.Println(err.Error())
			return nil, 0, err
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
	return comments, newSinceId, nil
}

// Implement
func (p *TwitterPipeline) Run(params map[string]string) *models.PipelineResult {
	sinceId, err := p.getComment(params)
	result := new(models.PipelineResult)

	if err != nil {
		result.Status = 0
		result.Message = err.Error()
	} else {
		result.Status = 1
		result.Message = "SUCCESS" + string(sinceId)
	}
	return result
}