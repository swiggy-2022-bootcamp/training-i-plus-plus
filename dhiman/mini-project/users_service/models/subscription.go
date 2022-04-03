package models

import "fmt"

type UserSubscription struct {
	modelImpl
	Medicine string `json:"title"`
	Rate     string `json:"rate"`
	NextDose string `json:"next_dose"`
}

func NewUserSubscription(id uint, medicine string, nextDose string, rate string) *UserSubscription {
	us := &UserSubscription{
		Medicine: medicine,
		NextDose: nextDose,
		Rate: rate,
	}
	us.SetId(fmt.Sprint(id))
	return us
}

func (us *UserSubscription) GetId() string {
	return us.id
}
