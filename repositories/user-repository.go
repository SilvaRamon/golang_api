package repository

import (
	"context"
	"hello/core"
	"hello/storage"
	"log"
)

type Repository struct {
	storage storage.Storage
}

type IUserRepository interface {
	GetUser(id string) core.User
	GetUsers() []core.User
}

func Initialize(storage storage.Storage) *Repository {
	return &Repository{storage}
}

func (r *Repository) GetUsers() []core.User {
	userList := []core.User{}

	rows, _ := r.storage.Conn.Query(context.Background(), "select u.name, u.email from users u")

	for rows.Next() {
		var name []byte
		var email []byte
		err := rows.Scan(&name, &email)

		if err != nil {
			log.Fatal(err.Error())
		}
		user := core.User{Name: string(name), Email: string(email)}
		userList = append(userList, user)
	}

	defer r.storage.Conn.Close(context.Background())

	return userList
}

func (r *Repository) GetUser(mail string) core.User {
	var name []byte
	var email []byte

	err := r.storage.Conn.
		QueryRow(context.Background(), "select u.name, u.email from users u where u.email=$1", mail).
		Scan(&name, &email)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer r.storage.Conn.Close(context.Background())

	return core.User{Name: string(name), Email: string(email)}
}
