package teststore_test

import (
	"github.com/Dennikoff/TodoAPI/internal/app/model"
	"github.com/Dennikoff/TodoAPI/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTodoRepository_Create(t *testing.T) {
	st := teststore.New()
	todo := model.TestTodo()
	user := model.TestUser()
	assert.NoError(t, st.User().Create(user))
	todo.UserId = user.ID
	assert.NoError(t, st.Todo().Create(todo))
	assert.NotEqual(t, 0, todo.ID)
}

func TestTodoRepository_FindByUserID(t *testing.T) {
	st := teststore.New()
	todo := model.TestTodo()
	user := model.TestUser()
	assert.NoError(t, st.User().Create(user))
	todo.UserId = user.ID
	assert.NoError(t, st.Todo().Create(todo))
	assert.NotEqual(t, 0, todo.ID)

	todos, err := st.Todo().FindByUserID(user.ID)

	assert.NoError(t, err)
	assert.NotNil(t, todos)
}