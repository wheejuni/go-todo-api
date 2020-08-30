package TodoService

import "./models"

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

func AddTodo(todoItem *models.NewTodoItem) {
	generatedItem := todoItem.TodoItemModel()
	generatedItem.Id = len(todoItems) - 1

	todoItems = append(todoItems, generatedItem)
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