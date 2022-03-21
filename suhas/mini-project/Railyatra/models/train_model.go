package models

type Train struct {
	Station1 string `json:"station1,omitempty" validate:"required"`
	Station2 string `json:"station2,omitempty" validate:"required"`
}
