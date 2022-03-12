package entity;

type UserExpert struct{
	Userid int	`json:"userid"`
	Expertid int	`json:"expertid"`
	Accepted bool 	`json:"accepted"`
	Cost int		`json:"cost"`
	Skill string	`json:"skill"`
}