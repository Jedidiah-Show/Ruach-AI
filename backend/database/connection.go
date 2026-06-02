package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {
	connString :=
		"postgres://ruach:ruach_password@localhost:5432/ruach_ai"

	conn, err := pgxpool.Connect(
		context.Background(),
		connString,
	)

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL")

	return conn, nil
}
