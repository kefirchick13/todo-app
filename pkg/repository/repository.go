package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kefirchick13/todo-app"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(UserName, Password string) (todo.User, error)
}

type TodoList interface {
	CreateList(userId int, list todo.TodoList) (int, error)
	GetAllLists(userId int) ([]todo.TodoList, error)
	GetListById(userId int, listId int) (todo.TodoList, error)
	DeleteListById(userId int, listId int) (error)
	UpdateListById(userId int, listId int, input todo.UpdateListInput) (error)
}

type TodoItem interface {
	CreateItem(listId int, input todo.TodoItem ) (int, error)
	GetAllItems(userId int , listId int) ([]todo.TodoItem, error)
	GetItemById(userId int , itemId int) (todo.TodoItem, error)
	DeleteItemById(userId int , itemId int) error
	UpdateItemById(userId, itemId int, input todo.UpdateItemInput) error

}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList: NewTodoListPostgres(db),
		TodoItem: NewTodoItemPostgres(db),
	}
}
