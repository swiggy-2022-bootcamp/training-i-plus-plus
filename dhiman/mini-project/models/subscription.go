package models

type Subscription struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Medicine string `json:"title"`
	Rate     string `json:"rate"`
	NextDose string `json:"next_dose"`
}
