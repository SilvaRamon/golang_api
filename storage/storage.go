package storage

import (
	"context"
	"fmt"
	"hello/core"
	"os"

	"github.com/jackc/pgx/v4"
)

const (
	host = ""
	db   = ""
	user = ""
	pass = ""
)

type Storage struct {
	Conn *pgx.Conn
}

type IUserRepository interface {
	GetUsers() []core.User
}

func Initialize() *Storage {
	var connStr string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, pass, db)

	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &Storage{Conn: conn}
}
