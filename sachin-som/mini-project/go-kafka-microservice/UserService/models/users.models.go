package models

type User struct {
	ID       string `json:"_id"           bson:"_id"`
	Fullname string `json:"fullname"      bson:"fullname"`
	Email    string `json:"email"         bson:"email"         validate:"email required"`
	Phone    string `json:"phone"         bson:"phone"         validate:"required"`
	Password string `json:"password"      bson:"password"      validate:"required"`
}
