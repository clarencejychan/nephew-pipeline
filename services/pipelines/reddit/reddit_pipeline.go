package reddit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/clarencejychan/nephew-pipeline/services/pipelines"
	"github.com/clarencejychan/nephew-pipeline/models"
)

type PushshiftQuery struct {
	Data []models.Comment `json:"data"`
}

type RedditPipeline struct {
	db *models.MongoDB
}

func (p *RedditPipeline) getComment(params map[string]string) error {
	// get pushshift comment
	comments, _, err := getPushshiftDataComment(
		params["subject"], 
		params["after"], 
		params["before"], 
		params["subreddit"])

	// get analysis result on each comment
	analysisReq := pipelines.AnalysisRequest{
		PlayerId: 123,
		Comments: comments.Data,
	}

	resp, err := pipelines.GetAnalysisResult(analysisReq)

	for i, _ := range resp.Comments {
		resp.Comments[i].Player_Id = resp.PlayerId
		resp.Comments[i].Source = "Reddit"
	}

	// Necessary, maybe we should put this in the bulk insert function
	x := make([]interface{}, len(resp.Comments))
	for i := range resp.Comments {
		x[i] = resp.Comments[i]
	}
	p.db.BulkInsert("comments", x)

	return err
}

func getPushshiftDataComment(query string, after string, before string, sub string) (PushshiftQuery, string, error) {
	url := fmt.Sprintf("https://api.pushshift.io/reddit/search/comment/?q=%s&size=5&after=%s&before=%s&subreddit=%s",
		query, after, before, sub)
	resp, err := http.Get(url)

	// defers closing the response body until end of function, prevents resource leaks
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var comments PushshiftQuery
	err = json.Unmarshal(body, &comments)

	return comments, url, err
}

func New(db *models.MongoDB) models.Pipeline {
	return &RedditPipeline{db: db}
}

// Implement
func (p *RedditPipeline) Run(params map[string]string) *models.PipelineResult {
	err := p.getComment(params)
	result := new(models.PipelineResult)

	if err != nil {
		result.Status = 0
		result.Message = err.Error()
	} else {
		result.Status = 1
		result.Message = "SUCCESS"
	}
	return result
}
