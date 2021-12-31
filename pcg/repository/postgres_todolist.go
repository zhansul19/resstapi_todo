package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	todo "github.com/zhansul19/restapi_todo"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListSPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (t *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", TodoListTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTbale)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
func (t *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1", TodoListTable, usersListsTbale)
	err := t.db.Select(&lists, query, userId)

	return lists, err
}
func (t *TodoListPostgres)GetById(userId,id int)(todo.TodoList,error){
	var list todo.TodoList
	query:=fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id=$2", TodoListTable, usersListsTbale)

	err := t.db.Get(&list, query, userId,id)

	return list, err
}

func (t *TodoListPostgres)Delete(userId,id int)(error){
		query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		TodoListTable, usersListsTbale)
	_, err := t.db.Exec(query, userId, id)

	return err
}
func (t *TodoListPostgres)Update(userId,id int ,input todo.UpdateListInput)(error){
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		TodoListTable, setQuery, usersListsTbale, argId, argId+1)
	args = append(args, id, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := t.db.Exec(query, args...)
	return err
}