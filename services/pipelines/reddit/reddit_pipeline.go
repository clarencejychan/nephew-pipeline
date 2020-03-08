package reddit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/clarencejychan/nephew-pipeline/models"
	"github.com/clarencejychan/nephew-pipeline/services/pipelines/aggregator"
)

type PushshiftQuery struct {
	Data []models.Comment `json:"data"`
}

type AnalysisResponse struct {
	PlayerId int              `json:"playerId`
	Comments []models.Comment `json:"comments`
}

type AnalysisRequest struct {
	PlayerId int              `json:"playerId"`
	Comments []models.Comment `json:"comments"`
}

type RedditPipeline struct {
	db *models.MongoDB
}

func (p *RedditPipeline) getComment(params map[string]string) error {
	// get pushshift comment
	comments, _, err := getPushshiftDataComment("Harden", "4d", "2d", "nba")

	// get analysis result on each comment
	analysisReq := AnalysisRequest{
		PlayerId: 123,
		Comments: comments.Data,
	}

	resp, err := getAnalysisResult(analysisReq)

	for i, _ := range resp.Comments {
		resp.Comments[i].Player_Id = resp.PlayerId
	}

	// Necessary, maybe we should put this in the bulk insert function
	x := make([]interface{}, len(resp.Comments))
	for i := range resp.Comments {
		x[i] = resp.Comments[i]
	}
	p.db.BulkInsert("comments", x)

	aggregator_pipeline := aggregator.AggregatorPipeline{Db: p.db}
	aggregator_pipeline.UpdateSematicScores(resp.PlayerId, resp.Comments)

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

func getAnalysisResult(req AnalysisRequest) (AnalysisResponse, error) {
	url := "http://127.0.0.1:5000/get-sentiments"

	b, err := json.Marshal(req)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(b))

	// defers closing the response body until end of function, prevents resource leaks
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var response AnalysisResponse
	err = json.Unmarshal(body, &response)

	return response, err
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
