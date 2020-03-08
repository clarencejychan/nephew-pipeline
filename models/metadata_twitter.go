package models

type Metadata_Twitter struct {
	CommentId				string			`json:"commentId"`	
	RetweetCount			uint			`json:"retweets"`
	FavoriteCount			uint			`json:"likes"`
	ReplyCount				uint			`json:"replies"`
	QuoteCount				uint			`json:"quotes"`
}