package pushshift

import "github.com/gin-gonic/gin"
import "fmt"
import "io/ioutil"
import "log"
import "net/http"
import "encoding/json"
import "github.com/clarencejychan/nephew-pipeline/models"

type PushshiftQuery struct {
	Data	[]models.Comment		`json:"data"`
}

type AnalysisResponse struct {
	Semantic_Rating		float64
}

func GetComment(db models.MongoDatastore) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		// get pushshift comment
		comments, url := getPushshiftDataComment("Harden", "4d", "2d", "nba")

		// get analysis result on each comment
		var commentsWithRating PushshiftQuery
		for i := range comments.Data {
			commentWithRating := getAnalysisResult(comments.Data[i])
			commentsWithRating.Data = append(commentsWithRating.Data, commentWithRating)
		}

		// write to db

		c.JSON(200, gin.H{
			"Lebron" : url,
			"Harden" : commentsWithRating.Data[0],
			"Sucks" : commentsWithRating.Data[4].Subject,
			"TestRating" : commentsWithRating.Data[0].Semantic_Rating,
		})
	}
	return gin.HandlerFunc(fn)
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

	// process pushshift response
	var comments PushshiftQuery
	err3 := json.Unmarshal(body, &comments)
	if err3 != nil {
		log.Fatalln(err3)
	}

	// append subject to comments struct
	for i := range comments.Data {
		comments.Data[i].Subject = query
	}

	return comments, url
}

func getAnalysisResult(comment models.Comment) (models.Comment) {
	// content := comment.Comment
	// subject := comment.Subject

	// url := "***ANALYSIS API LOCALHOST PORT WITH PARAMETERS***"

	// resp, err := http.Get(url)
		
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	
	// // defers closing the response body until end of function, prevents resource leaks
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)

	// var response AnalysisResponse
	// err = json.Unmarshal(body, &response)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	result := comment
	// result.Semantic_Rating = response.Semantic_Rating
	result.Semantic_Rating = 6.4

	return result
}
