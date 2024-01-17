package interfaces

import (
	"github.com/raphael251/users-crud/internal/domain/entity"
	pgkEntity "github.com/raphael251/users-crud/pkg/entity"
)

type UserRepositoryInterface interface {
	Create(user *entity.User) error
	FindById(id pgkEntity.ID) (*entity.User, error)
	Update(user *entity.User) error
}
