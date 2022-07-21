package model

import "time"

type Todo struct {
	ID          int       `json:"id"`
	Header      string    `json:"header"`
	Text        string    `json:"text"`
	CreatedDate time.Time `json:"createdDate"`
}
