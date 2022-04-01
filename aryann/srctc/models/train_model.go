package models

type Train struct {
	Source      string `json:"source,omitempty" validate:"required"`
	Destination string `json:"destination,omitempty" validate:"required"`
}
