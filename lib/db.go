package lib

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {
	connstring := "postgres://data:1@172.17.0.2:5432/data"
	// conn, _ := pgx.Connect(context.Background(), connstring)
	conn, _ := pgx.Connect(context.Background(), connstring)

	return conn
}
