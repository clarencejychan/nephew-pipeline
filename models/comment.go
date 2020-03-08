package models

type Comment struct {
	Id     					string			`json:"id"`
	Comment 				string 			`json:"body"`
	Source					string 			`json:"source"`
	Semantic_Rating			float64 		`json:"semantic_rating"`
	Date					uint			`json:"created_utc"`
	Author					string			`json:"author"`
	Parent					string			`json:"parent_id"`
	Children				[]int 			`json:"-"`
	Subject					string			`json:"-"`
	Player_Id				int				`json:"player_id"`
	Metadata				Metadata_Reddit	`json:"-"`
}