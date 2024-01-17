package user

import (
	"github.com/raphael251/users-crud/pkg/entity"
)

type UserRepositoryInterface interface {
	Create(user *User) error
	FindById(id entity.ID) (*User, error)
	Update(user *User) error
}
