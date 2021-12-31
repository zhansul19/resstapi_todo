package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	todo "github.com/zhansul19/restapi_todo"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB)*AuthPostgres{
	return &AuthPostgres{db: db}
}

func (s *AuthPostgres)CreateUser(user todo.User)(int,error){
	var id int
	query:=fmt.Sprintf("INSERT INTO %s (name,username,password_hash) VALUES ($1,$2,$3) RETURNING id",userTable)
	row:=s.db.QueryRow(query,user.Name,user.Username,user.Password)

	if err:=row.Scan(&id); err != nil {
		return 0,err
	}

	return id,nil
	
}
func (s *AuthPostgres)GetUser(username,password string)(todo.User,error){
	var user todo.User
	query:=fmt.Sprintf("SELECT ID FROM %s where username=$1 AND password_hash=$2",userTable)
	err:=s.db.Get(&user,query,username,password)
	
	return user,err
}
