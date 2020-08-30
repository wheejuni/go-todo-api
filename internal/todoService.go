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

func AddTodo(title string) {
	todoItem := new(models.TodoItem)

	todoItem.Id = len(todoItems)
	todoItem.Done = false
	todoItem.Title = title

	todoItems = append(todoItems, todoItem)
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