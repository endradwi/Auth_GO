package models

import (
	"context"
	"test/lib"

	"github.com/jackc/pgx/v5"
)

type Users struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname" form:"full_name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	// created_at time.Time `json:"create_at" form:"create_at"`
	// updated_at time.Time `json:"update_at" form:"update_at"`
}

type Listusers []Users

func FindOneUser(paramId int) Users {
	conn := lib.DB()
	defer conn.Close(context.Background())
	var user Users

	conn.QueryRow(context.Background(), `
	SELECT id, email, password 
	FROM users WHERE id = $1
	`, paramId).Scan(&user.Id, &user.Email, &user.Password)
	return user
}

func FindAllUser() Listusers {
	conn := lib.DB()
	defer conn.Close(context.Background())

	rows, _ := conn.Query(context.Background(), `
SELECT id, '' as fullname, password, email
from users
`)
	users, _ := pgx.CollectRows(rows, pgx.RowToStructByName[Users])
	return users
}

func FindOneUserByEmail(email string) Users {
	conn := lib.DB()
	defer conn.Close(context.Background())
	var user Users

	conn.QueryRow(context.Background(), `
	SELECT id, email, password 
	FROM users WHERE email = $1
	`, email).Scan(&user.Id, &user.Email, &user.Password)
	return user
}

func UpdateUser(user Users) Users {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var update Users

	conn.QueryRow(context.Background(), `
	UPDATE users SET email=$1, password=$2 WHERE id = $3
	RETURNING id, email, password
	`, user.Email, user.Password, user.Id).Scan(
		&update.Id,
		&update.Email,
		&update.Password,
	)
	return update
}

func DeleteUser(id int) Users {
	conn := lib.DB()
	defer conn.Close(context.Background())
	var delete Users
	conn.QueryRow(context.Background(), `
	DELETE FROM users WHERE id=$1
	RETURNING id, email, password
	`, id).Scan(
		&delete.Id,
		&delete.Email,
		&delete.Password,
	)
	return delete
}

func InsertUser(user Users) Users {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var new Users

	conn.QueryRow(context.Background(), `
	INSERT INTO users(email,password) VALUES
	($1,$2)
	RETURNING id, email, password
	`, user.Email, user.Password).Scan(
		&new.Id, &new.Email, &new.Password)
	return new
}
