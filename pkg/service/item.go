package service

import (
	"github.com/kefirchick13/todo-app"
	"github.com/kefirchick13/todo-app/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{
	repo: repo,
	listRepo: listRepo,
	}
}

func(s *TodoItemService) CreateItem(userId int, listId int, input todo.TodoItem) (int, error){
	_, err := s.listRepo.GetListById(userId, listId)
	if err != nil{
		// list does not exist or belongs to user
		return 0, err
	}

	return s.repo.CreateItem(listId, input)
}

func(s *TodoItemService) GetAllItems(userId int, listId int) ([]todo.TodoItem, error){
	_, err := s.listRepo.GetListById(userId, listId)
	if err != nil{
		// list does not exist or belongs to user
		return nil, err
	}

	return s.repo.GetAllItems(userId, listId)
}

func (s *TodoItemService)GetItemById(userId, itemId int)(todo.TodoItem,error){
	return s.repo.GetItemById(userId, itemId)
}

func (s *TodoItemService)UpdateItemById(userId, itemId int, input todo.UpdateItemInput) error{
	return s.repo.UpdateItemById(userId, itemId, input)
}

func (s *TodoItemService)DeleteItemById(userId, itemId int) error{
	return s.repo.DeleteItemById(userId, itemId)
}