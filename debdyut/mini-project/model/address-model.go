package model

type Address struct {
	ID       string `json:"id"`
	Location string `json:"location,omitempty" binding:"required,min=1,max=256"`
}
