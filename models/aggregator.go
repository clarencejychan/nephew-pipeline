package models

type Daily struct {
	PlayerId     			int				`json:"playerId"`
	Date					uint			`json:"created_utc"`
	Semantic_Rating			float64 		`json:"semantic_rating"`
}

type Monthly struct {
	PlayerId     			int				`json:"playerId"`
	Date					uint			`json:"created_utc"`
	Semantic_Rating			float64 		`json:"semantic_rating"`
}

type Yearly struct {
	PlayerId     			int				`json:"playerId"`
	Date					uint			`json:"created_utc"`
	Semantic_Rating			float64 		`json:"semantic_rating"`
}