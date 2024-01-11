package database

import (
	"database/sql"

	"github.com/raphael251/users-crud/internal/user"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Save(user *user.User) error {
	return nil
}
