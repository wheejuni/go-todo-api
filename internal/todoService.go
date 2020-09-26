package TodoService

import (
	"./models"
)

var todoItems []*models.TodoItem

func GetTodo(id int) *models.TodoItem {
	for _, element := range todoItems {
		if element.Id == id {
			return element
		}
		continue
	}
	return new(models.TodoItem)
}

func AddTodo(todoItem *models.NewTodoItem) *models.TodoItem {
	lastIndex := len(todoItems)

	generatedItem := todoItem.TodoItemModel()
	generatedItem.Id = 1

	if lastIndex > 1 {
		generatedItem.Id = lastIndex
	}

	todoItems = append(todoItems, generatedItem)
	return todoItems[generatedItem.Id - 1]
}

func DeleteTodo(itemNumber int) *models.TodoItem {
	return nil
}

func SetAsDone(itemNumber int) {
	for _, element := range todoItems {
		if element.Id == itemNumber {
			element.Done = true
			return
		}
		continue
	}
}