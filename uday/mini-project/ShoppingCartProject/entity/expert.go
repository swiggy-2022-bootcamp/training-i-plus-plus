package entity

type Expert struct{
	Id              int    `json:"id"`
	Username     	string	`json:"username"`
	Skill   	    string	`json:"skill"`
	Email 			string	`json:"email"`
	IsAvailable     bool	`json:"isAavailable"`
	Served			int		`json:"served"`
	// rating          []RatingStruct
}
