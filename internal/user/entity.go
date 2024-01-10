package user

import (
	"errors"
	"time"

	"github.com/raphael251/users-crud/pkg/entity"
	"github.com/raphael251/users-crud/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        entity.ID
	Name      string
	BirthDate time.Time
	Email     string
	Password  string
	Address   string
}

func NewUser(name string, birthDate time.Time, email, password, address string) (*User, error) {
	user := &User{
		ID:        entity.NewID(),
		Name:      name,
		BirthDate: birthDate,
		Email:     email,
		Address:   address,
	}

	err := user.IsValid()

	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	return user, nil
}

func (u *User) IsValid() error {
	validator := &utils.Validator{}

	if u.Name == "" {
		return errors.New("name should not be empty")
	}

	if !validator.IsEmail(u.Email) {
		return errors.New("invalid e-mail")
	}

	return nil
}
