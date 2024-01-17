package database

import (
	"database/sql"

	"github.com/raphael251/users-crud/internal/domain/entity"
	pkgEntity "github.com/raphael251/users-crud/pkg/entity"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *entity.User) error {
	_, err := r.DB.Exec(
		"insert into users(id, name, birth_date, email, password, address) values(?, ?, ?, ?, ?, ?)",
		user.ID, user.Name, user.BirthDate, user.Email, user.Password, user.Address,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindById(id pkgEntity.ID) (*entity.User, error) {
	var user entity.User

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

func (r *UserRepository) Update(user *entity.User) error {
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
