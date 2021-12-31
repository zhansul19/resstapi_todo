package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	userTable       = "users"
	TodoListTable   = "todo_lists"
	TodoItemTable   = "todo_item"
	usersListsTbale = "user_lists"
	ListItemsTable  = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SSLmode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DbName, cfg.Password, cfg.SSLmode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
