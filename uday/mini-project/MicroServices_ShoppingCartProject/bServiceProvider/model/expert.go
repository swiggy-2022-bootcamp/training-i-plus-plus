package model
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Expert struct{
	Id 	primitive.ObjectID	`bson:"_id"`
	Username     	string	`json:"username"`
	Skill   	    string	`json:"skill"`
	Email 			string	`json:"email"`
	IsAvailable     bool	`json:"isavailable"`
	Served			int		`json:"served"`
	Rating 			float64 `json:"rating"`
	Location        int     `json:"location"`
	Reviews         []RatingStruct  `json:"reviews"`
}

type RatingStruct struct{
	Rating int `json:"rating"`
	Review string `json:"review"`
}