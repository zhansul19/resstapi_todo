package repository

import (
	"github.com/jmoiron/sqlx"
	todo "github.com/zhansul19/restapi_todo"
)

type Authorization interface {
	CreateUser(user todo.User)(int,error)
	GetUser(username,password string)(todo.User,error)
}
type TodoList interface {
	Create(userId int,list todo.TodoList)(int,error)
	GetAll(userId int)([]todo.TodoList,error)
	GetById(userId,id int)(todo.TodoList,error)
	Delete(userId,id int)(error)
	Update(userId,id int ,input todo.UpdateListInput)(error)
}
type TodoItem interface {
	Create(listId int,input todo.TodoItem)(int,error)
	GetAll(userId,listId int)([]todo.TodoItem,error)
	GetItemsById(userId,itemId int)(todo.TodoItem,error)
	Delete(userId,itemId int)(error)
	UpdateItem(userId,itemId int ,input todo.UpdateItemInput)(error)
}
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList: NewTodoListSPostgres(db),
		TodoItem: NewTodoItemPostgres(db),
	}
}
