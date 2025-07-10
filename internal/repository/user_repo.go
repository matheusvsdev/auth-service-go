package repository

import (
	"database/sql"

	"github.com/matheusvsdev/auth-service-go/internal/domain"
)

type UserRepository struct {
	DB *sql.DB
}

func (ur *UserRepository) CreateUser(user domain.User) error {
	_, err := ur.DB.Exec(`
	INSERT INTO users (name, email, password_hash, provider, plan) VALUES ($1, $2, $3, $4, $5)
	`, user.Name, user.Email, user.PasswordHash, user.Provider, user.Plan)

	return err
}

func (ur *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	row := ur.DB.QueryRow(`
		SELECT id, name, email, password_hash, provider, plan
		FROM users WHERE email = $1
	`, email)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Provider, &user.Plan)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
