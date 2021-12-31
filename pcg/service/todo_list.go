package service

import (
	todo "github.com/zhansul19/restapi_todo"
	"github.com/zhansul19/restapi_todo/pcg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList)*TodoListService{
	return &TodoListService{repo: repo}
}

func (t* TodoListService)Create(userId int,list todo.TodoList)(int,error){
	return t.repo.Create(userId,list)
}
func (t* TodoListService)GetAll(userId int)([]todo.TodoList,error){
	return t.repo.GetAll(userId)
}
func (t* TodoListService)GetById(userId,id int)(todo.TodoList,error){
	return t.repo.GetById(userId,id)
}
func (t* TodoListService)Delete(userId,id int)(error){
	return t.repo.Delete(userId,id)
}
func(t*TodoListService)Update(userId,id int ,input todo.UpdateListInput)(error){
	if err:=input.Validate();err!=nil{
		return err
	}
	return t.repo.Update(userId,id,input)
}

