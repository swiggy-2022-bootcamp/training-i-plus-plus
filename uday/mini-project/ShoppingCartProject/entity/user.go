package entity
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type User struct{
	Id 	primitive.ObjectID	`bson:"_id"`
	Username string `json:"username"`
	Password string  `json:"password"`
	Email string	`json:"email"`
	Phone int		`json:"phone"`
	Location int    `json:"location"`
}

