package model

type Address struct {
	ID       string `json:"id"`
	Location string `json:"location,omitempty"`
}
