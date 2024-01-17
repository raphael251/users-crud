package user

import (
	"github.com/raphael251/users-crud/internal/domain/utils"
	"github.com/raphael251/users-crud/pkg/entity"
)

type FindOneUserUseCaseOutputDTO struct {
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
	Email     string `json:"email"`
	Address   string `json:"address"`
}

type FindOneUserUseCase struct {
	UserRepository UserRepositoryInterface
}

func NewFindOneUserUseCase(userRepository UserRepositoryInterface) *FindOneUserUseCase {
	return &FindOneUserUseCase{
		UserRepository: userRepository,
	}
}

func (c *FindOneUserUseCase) Execute(id string) (*FindOneUserUseCaseOutputDTO, error) {
	parsedId, err := entity.ParseID(id)

	if err != nil {
		return nil, &utils.UseCaseError{Type: utils.ValidationError, Message: "invalid user id"}
	}

	user, err := c.UserRepository.FindById(parsedId)

	if err != nil {
		return nil, &utils.UseCaseError{Type: utils.NotFoundError, Message: "user not found"}
	}

	return &FindOneUserUseCaseOutputDTO{
		Name:      user.Name,
		BirthDate: user.BirthDate.Format("2006-01-02"),
		Email:     user.Email,
		Address:   user.Address,
	}, nil
}
