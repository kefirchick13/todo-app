package service

import (
	"github.com/kefirchick13/todo-app"
	"github.com/kefirchick13/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(Username string, Password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetListById(userId int, listId int) (todo.TodoList, error)
	DeleteListById(userId int, listId int) (error)
	UpdateListById(userId int, listId int, input todo.UpdateListInput) (error)
}

type TodoItem interface {
	CreateItem(userId, listId int, input todo.TodoItem) (int, error)
	GetAllItems(userId, listId int) ([]todo.TodoItem, error)
	GetItemById(userId, itemId int)(todo.TodoItem,error)
	DeleteItemById(userId, itemId int) error
	UpdateItemById(userId, itemId int, input todo.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList: NewTodoListService(repos.TodoList),
		TodoItem: NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
