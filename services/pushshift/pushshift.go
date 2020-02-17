package pushshift

import "github.com/gin-gonic/gin"
import "fmt"
import "io/ioutil"
import "log"
import "net/http"
import "encoding/json"
import "github.com/clarencejychan/nephew-pipeline/models"
import "bytes"

type PushshiftQuery struct {
	Data	[]models.Comment		`json:"data"`
}

type AnalysisRequset struct {
	Topic	string					`json:"topic"`
	Comments []models.Comment		`json:"comments"`
}


type AnalysisResponse struct {
	PlayerId	int					`json:"playerId`
	Comments	[]models.Comment	`json:"comments`
}

func GetComment(db models.MongoDatastore) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		// get pushshift comment
		comments, topic, url := getPushshiftDataComment("Harden", "4d", "2d", "nba")

		// get analysis result on each comment
		analysisReq := AnalysisRequset {
			Topic: topic,
			Comments: comments.Data,
		}

		resp:= getAnalysisResult(analysisReq)

		// write to db

		c.JSON(200, gin.H{
			"url": url,
			"result": resp,
		})
	}
	return gin.HandlerFunc(fn)
}

func getPushshiftDataComment(query string, after string, before string, sub string) (PushshiftQuery, string, string) {
	url := fmt.Sprintf("https://api.pushshift.io/reddit/search/comment/?q=%s&size=5&after=%s&before=%s&subreddit=%s", 
					  query, after, before, sub)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
	}

	// defers closing the response body until end of function, prevents resource leaks
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}

	var comments PushshiftQuery
	err = json.Unmarshal(body, &comments)
	if err != nil {
		log.Println(err.Error())
	}

	return comments, query, url
}

func getAnalysisResult(req AnalysisRequset) (AnalysisResponse) {
	url := "http://127.0.0.1:5000/get-sentiments"

	b, err := json.Marshal(req)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(b))
		
	if err != nil {
		log.Fatalln(err)
	}
	
	// defers closing the response body until end of function, prevents resource leaks
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var response AnalysisResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	return response
}