package pipelines

import (
	"github.com/clarencejychan/nephew-pipeline/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)
type AnalysisResponse struct {
	PlayerId int              `json:"playerId`
	Comments []models.Comment `json:"comments`
}

type AnalysisRequest struct {
	PlayerId int              `json:"playerId"`
	Comments []models.Comment `json:"comments"`
}

func GetAnalysisResult(req AnalysisRequest) (AnalysisResponse, error) {
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
