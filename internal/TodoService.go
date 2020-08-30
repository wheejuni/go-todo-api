package internal

import "./models"

var todoItems []*models.TodoItem

func AddTodo(title string) {
	todoItem := new(models.TodoItem)

	todoItem.Id = len(todoItems)
	todoItem.Done = false
	todoItem.Title = title

	todoItems = append(todoItems, todoItem)
}

func SetAsDone(itemNumber int) {

}