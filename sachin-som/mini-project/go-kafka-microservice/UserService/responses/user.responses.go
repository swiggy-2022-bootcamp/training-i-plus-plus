package responses

type MessageResponse struct {
	Message string `json:"message"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	Fullname string `json:"fullname"      bson:"fullname"      validate:"required"`
	Email    string `json:"email"         bson:"email"         validate:"email required"`
	Phone    string `json:"phone"         bson:"phone"         validate:"required"`
	Password string `json:"password"      bson:"password"      validate:"required"`
}
