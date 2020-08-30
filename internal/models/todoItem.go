package models

type TodoItem struct {
	Id int
	Title string
	Done bool
}

type NewTodoItem struct {
	Title string
}

func (n *NewTodoItem) toTodoItemModel() TodoItem {
	return TodoItem{
		Title: n.Title,
		Done: false}
}