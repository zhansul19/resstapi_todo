package service

import (
	todo "github.com/zhansul19/restapi_todo"
	"github.com/zhansul19/restapi_todo/pcg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
	listRepo repository.TodoList
}
func NewTodoItemService(repo repository.TodoItem,listrepo repository.TodoList)*TodoItemService{
	return&TodoItemService{repo: repo,listRepo: listrepo}
}
func (t* TodoItemService)Create(userId,listId int,item todo.TodoItem)(int,error){
	_, err := t.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return t.repo.Create(listId, item)
}
func(t*TodoItemService)GetAll(userId,listId int)([]todo.TodoItem,error){
	return t.repo.GetAll(userId,listId)
}
func(t*TodoItemService)GetItemsById(userId,itemId int)(todo.TodoItem,error){
	return t.repo.GetItemsById(userId,itemId)
}
func(t*TodoItemService)Delete(userId,itemId int)(error){
	return t.repo.Delete(userId,itemId)
}
func(t*TodoItemService)UpdateItem(userId,itemId int ,input todo.UpdateItemInput)(error){
	if err:=input.Validate();err != nil {
		return err
	}
	return t.repo.UpdateItem(userId,itemId,input)
}

