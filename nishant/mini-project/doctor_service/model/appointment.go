package models

// Appointment
// @Description Appointment
type Appointment struct {
	Patient string `json:"patient" form:"patient" binding:"required" bson:"patient"`
	From    int64  `json:"from" form:"from" binding:"required" bson:"from"`
	To      int64  `json:"to" bson:"to"`
}
