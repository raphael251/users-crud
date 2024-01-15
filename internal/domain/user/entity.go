package user

import (
	"time"

	"github.com/raphael251/users-crud/pkg/entity"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	return user, nil
}
