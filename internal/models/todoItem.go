package models

type TodoItem struct {
	Id int
	Title string
	Done bool
}

type NewTodoItem struct {
	Title string
}

func (n NewTodoItem) TodoItemModel() *TodoItem {
	generatedItem := new(TodoItem)
	generatedItem.Title = n.Title
	generatedItem.Done = false

	return generatedItem
}