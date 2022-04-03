package model;
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type UserExpert struct{
	Userid primitive.ObjectID	`json:"userid"`
	Expertid primitive.ObjectID	`json:"expertid"`
	Accepted bool 	`json:"accepted"`
	Cost int		`json:"cost"`
	Skill string	`json:"skill"`
	OrderedAt string `json:orderedat`
	AcceptedAt string `json:acceptedat` 
	Status string `json:status`
}