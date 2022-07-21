package main

import (
	"fmt"
	"github.com/Dennikoff/TodoAPI/internal/app/model"
)

func main() {
	u := model.User{}
	fmt.Println(u.Email, u.EncryptedPassword, u.Password, u.ID, u.TodoId)
}
