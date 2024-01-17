package database

import (
	"database/sql"

	"github.com/raphael251/users-crud/internal/domain/user"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *user.User) error {
	stmt, err := r.DB.Prepare("insert into users(id, name, birth_date, email, password, address) values(?, ?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Name, user.BirthDate, user.Email, user.Password, user.Address)

	if err != nil {
		return err
	}

	return nil
}
