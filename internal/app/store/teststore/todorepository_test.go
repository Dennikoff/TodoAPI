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
	assert.NoError(t, st.Todo().Create(todo))
	assert.NotEqual(t, 0, todo.ID)
}
