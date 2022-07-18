package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 30)),
	)
}

func (u *User) BeforeCreate() error {
	pas, err := bcrypt.GenerateFromPassword([]byte(u.Password), 4)
	if err != nil {
		return err
	}
	u.EncryptedPassword = string(pas)
	return nil
}
