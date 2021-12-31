package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	todo "github.com/zhansul19/restapi_todo"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}
func (t *TodoItemPostgres) Create(listId int, item todo.TodoItem) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}
	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title,description) values($1,$2) returning id", TodoItemTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s(list_id, items_id)values($1,$2)", ListItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()

}
func (t *TodoItemPostgres) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description, tl.done FROM %s tl INNER JOIN %s li on li.items_id=tl.id INNER JOIN %s ul on ul.list_id=li.list_id WHERE li.list_id = $1 AND ul.user_id= $2",
		TodoItemTable, ListItemsTable, usersListsTbale)
	if err := t.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}
func (t *TodoItemPostgres) GetItemsById(userId, itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description, tl.done FROM %s tl INNER JOIN %s li on li.items_id=tl.id INNER JOIN %s ul on ul.list_id=li.list_id WHERE tl.id = $1 AND ul.user_id= $2",
		TodoItemTable, ListItemsTable, usersListsTbale)
	if err := t.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}
	return item, nil
}
func (t *TodoItemPostgres) Delete(userId, id int) error {
	query := fmt.Sprintf("DELETE FROM %s ti USING %s li,%s ul WHERE ti.id = li.items_id AND li.list_id=ul.list_id AND ul.user_id=$1 AND ti.id=$2",
		TodoItemTable, ListItemsTable, usersListsTbale)
	_, err := t.db.Exec(query, userId, id)

	return err
}
func (t *TodoItemPostgres) UpdateItem(userId, id int, input todo.UpdateItemInput) error {
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
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id = li.items_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		TodoItemTable, setQuery, ListItemsTable, usersListsTbale, argId, argId+1)
	args = append(args, userId, id)

	_, err := t.db.Exec(query, args...)
	return err
}
