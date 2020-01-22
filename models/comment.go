package models


type Comment struct {
	Id     					uint			`json:"id"`
	Comment 				string 			`json:"comment"`
	Source					string 			`json:"source"`
	Semantic_Rating			float64 		`json:"semantic_rating"`
	Date					uint			`json:"date"`
	Author					uint			`json:"author"`
	Length					uint			`json:"length"`
	Parent					uint			`json:"parent"`
	Children				[]int 			`json:"children"`
	//Metadata				Metadata_Reddit	`json:"metadata"`
}