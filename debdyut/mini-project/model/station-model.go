package model

// Station represents a single user profile.
// ID should be globally unique.
type Station struct {
	ID      string  `json:"id"`
	Name    string  `json:"name,omitempty" binding:"required,min=1,max=256"`
	Address Address `json:"address,omitempty" binding:"required"`
}
