package entity;
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type UserExpert struct{
	 
	Userid primitive.ObjectID	`json:"userid"`
	Expertid primitive.ObjectID	`json:"expertid"`
	Accepted bool 	`json:"accepted"`
	Cost int		`json:"cost"`
	Skill string	`json:"skill"`
}