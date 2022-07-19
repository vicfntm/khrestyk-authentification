package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

const (
	usersTable = "users"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPsql(conf Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", conf.Host, conf.Port, conf.Username, conf.DBName, conf.Password, conf.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
