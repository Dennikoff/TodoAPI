package teststore_test

import (
	"github.com/Dennikoff/TodoAPI/internal/app/model"
	"github.com/Dennikoff/TodoAPI/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	st := teststore.New()
	user := model.TestUser()
	assert.NoError(t, st.User().Create(user))
	assert.NotEqual(t, 0, user.ID)

}
