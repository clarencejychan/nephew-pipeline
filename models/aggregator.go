package models

type Aggregation struct {
	PlayerId        int     `json:"playerId"`
	Date            string  `json:"date"`
	Semantic_Rating float64 `json:"semantic_rating"`
}
