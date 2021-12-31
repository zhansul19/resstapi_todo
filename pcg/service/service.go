package service

import (
	todo "github.com/zhansul19/restapi_todo"
	"github.com/zhansul19/restapi_todo/pcg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList,error)
	GetById(userId,id int)(todo.TodoList,error)
	Delete(userId,id int)(error)
	Update(userId,id int ,input todo.UpdateListInput)(error)
}
type TodoItem interface {
	Create(userId,listId int,input todo.TodoItem)(int,error)
	GetAll(userId,listId int)([]todo.TodoItem,error)
	GetItemsById(userId,itemId int)(todo.TodoItem,error)
	Delete(userId,itemId int)(error)
	UpdateItem(userId,itemId int ,input todo.UpdateItemInput)(error)
}
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
		
	}
}
