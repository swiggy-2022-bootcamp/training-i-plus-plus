package requests

type UserRequest struct {
	Fullname string `json:"fullname"      bson:"fullname"      validate:"required"`
	Email    string `json:"email"         bson:"email"         validate:"email required"`
	Phone    string `json:"phone"         bson:"phone"         validate:"required"`
	Password string `json:"password"      bson:"password"      validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
