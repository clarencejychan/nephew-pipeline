package models

type Comment struct {
	Id     					string			`json:"id"`
	Comment 				string 			`json:"body"`
	Source					string 			`json:"-"`
	Semantic_Rating			float64 		`json:"semantic_rating"`
	Date					uint			`json:"created_utc"`
	Author					string			`json:"author"`
	Parent					string			`json:"parent_id"`
	Children				[]int 			`json:"-"`
	Subject					string			`json:"-"`
	Metadata				Metadata_Reddit	`json:"-"`
}