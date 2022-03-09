package model

type UserExpert struct {
	userid   int
	expertid int
	accepted bool
	cost     int
	skill    string
}

type User struct {
	Id       int
	username string
	password string
	email    string
	phone    int
}

type Expert struct {
	id          int
	username    string
	skill       string
	email       string
	isAvailable bool
	served      int
	rating      []RatingStruct
}
