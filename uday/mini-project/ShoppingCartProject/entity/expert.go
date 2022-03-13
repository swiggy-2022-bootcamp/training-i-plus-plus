package entity

type Expert struct{
	Id              int    `json:"id"`
	Username     	string	`json:"username"`
	Skill   	    string	`json:"skill"`
	Email 			string	`json:"email"`
	IsAvailable     bool	`json:"isAavailable"`
	Served			int		`json:"served"`
	Rating 			float64 `json:"rating"`
	Reviews          []RatingStruct 
}

type RatingStruct struct{
	Rating int `json:"rating"`
	Review string `json:"review"`
}