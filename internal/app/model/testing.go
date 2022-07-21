package model

import "time"

func TestUser() *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}

func TestTodo() *Todo {
	return &Todo{
		Header:      "Test Header",
		Text:        "Test Text",
		CreatedDate: time.Now(),
	}
}
