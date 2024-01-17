package database

import (
	"database/sql"

	"github.com/raphael251/users-crud/internal/domain/user"
	"github.com/raphael251/users-crud/pkg/entity"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// TODO: remove prepare statement
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

func (r *UserRepository) FindById(id entity.ID) (*user.User, error) {
	var user user.User

	err := r.DB.
		QueryRow(
			"SELECT id, name, birth_date, email, password, address FROM users where id = ?",
			id,
		).
		Scan(&user.ID, &user.Name, &user.BirthDate, &user.Email, &user.Password, &user.Address)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(user *user.User) error {
	_, err := r.DB.Exec(
		"UPDATE users SET name = ?, birth_date = ?, password = ?, address = ? WHERE id = ?",
		user.Name,
		user.BirthDate,
		user.Password,
		user.Address,
		user.ID,
	)

	if err != nil {
		return err
	}

	return nil
}
