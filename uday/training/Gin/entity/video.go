package entity

type Person struct{
	FirstName string `json:"firstname" `
	LastName string `json:"lastname"`
	Age int8 `json:"age" binding:"gte=1,lte=130"`
	Email string `json:"email" validate:"required,email"`
}
// "vallidate":"email" binding:"required"
type Video struct{
	Title string `json:"title" form:"title" `
	Description string `json:"description"`
	URL string `json:"url"`	
}