package model

// Station represents a single user profile.
// ID should be globally unique.
type Station struct {
	ID      string  `json:"id"`
	Name    string  `json:"name,omitempty"`
	Address Address `json:"address,omitempty"`
}
