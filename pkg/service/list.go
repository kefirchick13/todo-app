package service

import (
	"github.com/kefirchick13/todo-app"
	"github.com/kefirchick13/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s * TodoListService)Create(userId int, list todo.TodoList) (int, error){
	return s.repo.CreateList(userId, list)
}

func (s * TodoListService)GetAll(userId int) ([]todo.TodoList, error){
	return s.repo.GetAllLists(userId)
}

func (s * TodoListService)GetListById(userId int, listId int) (todo.TodoList, error){
	return s.repo.GetListById(userId, listId)
}

func (s * TodoListService)DeleteListById(userId int, listId int) (error){
	return s.repo.DeleteListById(userId, listId)
}

func (s * TodoListService)UpdateListById(userId int, listId int, input todo.UpdateListInput) (error){
	if err := input.ValidateListInput(); err != nil{
		return err
	}
	return s.repo.UpdateListById(userId, listId, input)
}
