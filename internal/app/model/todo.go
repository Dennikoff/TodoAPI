package model

import "time"

type Todo struct {
	ID          int       `json:"id"`
	UserId      int       `json:"user_id"`
	Header      string    `json:"header"`
	Text        string    `json:"text"`
	CreatedDate time.Time `json:"created_date"`
}
