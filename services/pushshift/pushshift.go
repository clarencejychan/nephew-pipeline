package pushshift

import "github.com/gin-gonic/gin"
import "fmt"
import "io/ioutil"
import "log"
import "net/http"
import "encoding/json"
import comment "github.com/clarencejychan/nephew-pipeline/models"

type PushshiftQuery struct {
	Data	[]comment.Comment		`json:"data"`
}

func getPushshiftDataComment(query string, after string, before string, sub string) (PushshiftQuery, string) {
	url := fmt.Sprintf("https://api.pushshift.io/reddit/search/comment/?q=%s&size=5&after=%s&before=%s&subreddit=%s", 
					  query, after, before, sub)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	// defers closing the response body until end of function, prevents resource leaks
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatalln(err2)
	}

	var pushShiftQueryResult PushshiftQuery
	err3 := json.Unmarshal(body, &pushShiftQueryResult)
	if err3 != nil {
		log.Fatalln(err3)
	}

	for i := range pushShiftQueryResult.Data {
		pushShiftQueryResult.Data[i].Subject = query
	}

	return pushShiftQueryResult, url
}

func IndexHandler(c *gin.Context) {
	pushShiftQueryResult, url := getPushshiftDataComment("Harden", "4d", "2d", "nba")

	c.JSON(200, gin.H{
		"Lebron" : url,
		"Harden" : pushShiftQueryResult.Data[0],
		"Sucks" : pushShiftQueryResult.Data[4],
	})
}