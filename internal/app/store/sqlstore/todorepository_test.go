package sqlstore_test

import (
	"github.com/Dennikoff/TodoAPI/internal/app/model"
	"github.com/Dennikoff/TodoAPI/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTodoRepository_Create(t *testing.T) {
	db, del := sqlstore.TestDB(t, databaseURL)

	defer del("users")

	store := sqlstore.New(db)
	user := model.TestUser()
	assert.NoError(t, store.User().Create(user))

	todo := model.TestTodo()
	todo.UserId = user.ID
	assert.NoError(t, store.Todo().Create(todo))
	assert.NotEqual(t, 0, todo.ID)

}

func TestTodoRepository_FindByUserID(t *testing.T) {
	db, del := sqlstore.TestDB(t, databaseURL)

	defer del("users")

	store := sqlstore.New(db)
	user := model.TestUser()
	assert.NoError(t, store.User().Create(user))

	todo := model.TestTodo()
	todo.UserId = user.ID
	assert.NoError(t, store.Todo().Create(todo))
	assert.NotEqual(t, 0, todo.ID)

	todos, err := store.Todo().FindByUserID(user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, todos)
}