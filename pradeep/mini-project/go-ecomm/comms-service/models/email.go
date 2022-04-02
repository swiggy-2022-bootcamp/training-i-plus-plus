package models

import "time"

type Email struct {
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	Message     string    `json:"message"`
	PublishedAt time.Time `json:"publishedAt"`
}
