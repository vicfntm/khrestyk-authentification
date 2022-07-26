package repository

import (
	"fmt"

	"example.com/hello/src/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, password_hash) values($1, $2) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUserFromDb(login string, password string) (models.LoginUserStruct, error) {
	var user models.LoginUserStruct
	query := fmt.Sprintf("SELECT id FROM %s WHERE name=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, login, password)

	return user, err
}
