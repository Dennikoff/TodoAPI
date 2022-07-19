package sqlstore_test

import (
	"github.com/Dennikoff/TodoAPI/internal/app/model"
	"github.com/Dennikoff/TodoAPI/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, del := sqlstore.TestDB(t, databaseURL)

	defer del("users")

	user := model.TestUser()

	store := sqlstore.New(db)
	assert.NoError(t, store.User().Create(user))
	assert.NotNil(t, user.ID)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, del := sqlstore.TestDB(t, databaseURL)
	defer del("users")

	user := model.TestUser()
	store := sqlstore.New(db)

	_, err := store.User().FindByEmail(user.Email)
	assert.Error(t, err)

	assert.NoError(t, store.User().Create(user))

	us, err := store.User().FindByEmail(user.Email)
	assert.NotNil(t, us)
	assert.NoError(t, err)
}
