package model

import (
	"database/sql"
	"fmt"
	"os"
)

type PostgreSQL struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func (p *PostgreSQL) Connect() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	postgres := PostgreSQL{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Name:     dbName,
	}

	info := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		postgres.Host,
		postgres.Port,
		postgres.User,
		postgres.Password,
		postgres.Name,
	)

	db, err := sql.Open("postgres", info)

	if err != nil {
		return nil, err
	}

	return db, nil
}
