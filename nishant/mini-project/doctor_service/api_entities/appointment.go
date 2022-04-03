package api_entities

// Appointment
// @Description Appointment Request
type AppointmentRequest struct {
	Doctor  string
	Patient User
	From    string `example:"02 Jan 22 15:00 IST"`
	To      string `example:"02 Jan 22 16:00 IST"`
}

type User struct {
	UserId string `json:"userId,omitempty"`
	Email  string `json:"email,omitempty"`
	Name   string `json:"username,omitempty"`
}
