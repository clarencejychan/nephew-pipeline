package models

type Player struct {
	Id     					uint			`json:"id"`
	First_Name				string			`json:"first_name"`
	Last_Name				string			`json:"last_name"`
	Nicknames				[]string		`json:"-"`
	Current_Team			[]uint			`json:"current_team"`
	Past_Teams				[]uint			`json:"past_teams"`
}