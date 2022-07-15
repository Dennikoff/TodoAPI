package model_test

import (
	"github.com/Dennikoff/TodoAPI/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name string
		user func() *model.User
		ok   bool
	}{
		{
			name: "valid_case",
			user: func() *model.User {
				return model.TestUser()
			},
			ok: true,
		},
		{
			name: "empty pass",
			user: func() *model.User {
				u := model.TestUser()
				u.Password = ""
				return u
			},
			ok: false,
		},
		{
			name: "empty email",
			user: func() *model.User {
				u := model.TestUser()
				u.Email = ""
				return u
			},
			ok: false,
		},
		{
			name: "incorrect email 1",
			user: func() *model.User {
				u := model.TestUser()
				u.Email = "user@gmail"
				return u
			},
			ok: false,
		},
		{
			name: "incorrect email 2",
			user: func() *model.User {
				u := model.TestUser()
				u.Email = "usergmail.ru"
				return u
			},
			ok: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.ok {
				assert.NoError(t, tc.user().Validate())
			} else {
				assert.Error(t, tc.user().Validate())
			}
		})
	}
}
