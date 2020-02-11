package models

type Player struct {
	Id     					uint			`json:"playerId"`
	Name					string			`json:"fullName"`
	First_Name				string			`json:"firstName"`
	Last_Name				string			`json:"lastName"`
	Nicknames				[]string		`json:"-"`
	Current_Team			uint			`json:"teamId"`
	Past_Teams				[]uint			`json:"-"`
}