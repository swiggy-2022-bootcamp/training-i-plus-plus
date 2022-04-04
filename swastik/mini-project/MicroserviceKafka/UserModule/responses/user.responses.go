package responses

type MessageResponse struct {
	Message string `json:"message"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type AddressResponse struct{
	State 	string	`json:"state"`
	City 	string	`json:"city"`
	Pincode	int		`json:"pincode"`
}

type UserUserResponse struct{
	Name 	string 			`json:"name"`
	Age 	int				`json:"age"`
	Address AddressResponse	`json:"address"`
	Username string			`json:"username"`
	Password string			`json:"password"`
}
