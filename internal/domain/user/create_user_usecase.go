package user

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/raphael251/users-crud/internal/domain/utils"
)

type CreateUserInputDTO struct {
	Name      string `json:"name" validate:"required"`
	BirthDate string `json:"birth_date" validate:"required,datetime=2006-01-02"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Address   string `json:"address" validate:"required"`
}

func (user *CreateUserInputDTO) Validate() []error {
	validate := validator.New()
	err := validate.Struct(user)

	if err != nil {
		errs := make([]error, 0)
		for _, e := range err.(validator.ValidationErrors) {
			errs = append(errs, fmt.Errorf("invalid field: %s", e.Field()))
		}
		return errs
	}

	return nil
}

type CreateUserOutputDTO struct {
	ID string `json:"id"`
}

type CreateUserUseCase struct {
	UserRepository UserRepositoryInterface
}

func NewCreateUserUseCase(userRepository UserRepositoryInterface) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: userRepository,
	}
}

func (c *CreateUserUseCase) Execute(input CreateUserInputDTO) (*CreateUserOutputDTO, error) {
	birthDate, err := time.Parse("2006-01-02", input.BirthDate)

	if err != nil {
		return nil, &utils.UseCaseError{Type: utils.Validation, Message: "the birth date is not valid. Please see the docs."}
	}

	user, err := NewUser(input.Name, birthDate, input.Email, input.Password, input.Address)

	if err != nil {
		return nil, err
	}

	if err := c.UserRepository.Save(user); err != nil {
		return nil, err
	}

	return &CreateUserOutputDTO{
		ID: user.ID.String(),
	}, nil
}
