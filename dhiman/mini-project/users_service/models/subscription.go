package models

import "fmt"

// Data model for an Client's Subscription to a medicine.
type ClientSubscription struct {
	modelImpl
	Medicine string `json:"title"`
	Rate     string `json:"rate"`
	NextDose string `json:"next_dose"`
}

// Generate a new Client Subscription with the given data.
func NewClientSubscription(id uint, medicine string, nextDose string, rate string) *ClientSubscription {
	us := &ClientSubscription{
		Medicine: medicine,
		NextDose: nextDose,
		Rate: rate,
	}
	us.SetId(fmt.Sprint(id))
	return us
}

// Get the ID of this Subscription.
func (us *ClientSubscription) GetId() string {
	return us.id
}
