package usecase

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/raphael251/users-crud/internal/domain/entity"
	"github.com/raphael251/users-crud/internal/domain/interfaces"
	"github.com/raphael251/users-crud/internal/domain/utils"
	pkgEntity "github.com/raphael251/users-crud/pkg/entity"
)

type UpdateUserInputDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BirthDate string `json:"birth_date" validate:"datetime=2006-01-02"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Address   string `json:"address"`
}

func (user *UpdateUserInputDTO) Validate() []error {
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

type UpdateUserUseCase struct {
	UserRepository interfaces.UserRepositoryInterface
}

func NewUpdateUserUseCase(userRepository interfaces.UserRepositoryInterface) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		UserRepository: userRepository,
	}
}

func (c *UpdateUserUseCase) Execute(input UpdateUserInputDTO) error {
	if input.Email != "" {
		return &utils.UseCaseError{Type: utils.ValidationError, Message: "the e-mail cannot be changed. To use another e-mail you need to create another account."}
	}

	var birthDate time.Time

	if input.BirthDate != "" {
		parsedBirthDate, err := time.Parse("2006-01-02", input.BirthDate)
		birthDate = parsedBirthDate

		if err != nil {
			return &utils.UseCaseError{Type: utils.ValidationError, Message: "the birth date is not valid. Please see the docs."}
		}
	}

	id, err := pkgEntity.ParseID(input.ID)

	if err != nil {
		return &utils.UseCaseError{Type: utils.ValidationError, Message: "invalid user id"}
	}

	foundUser, err := c.UserRepository.FindById(id)

	if err != nil {
		return &utils.UseCaseError{Type: utils.BusinessRuleViolationError, Message: "invalid user id"}
	}

	builtUser, err := entity.BuildUser(id, input.Name, birthDate, foundUser.Email, input.Password, input.Address)

	if err != nil {
		return &utils.UseCaseError{Type: utils.InternalError, Message: "could not retrieve and build existing user"}
	}

	err = c.UserRepository.Update(builtUser)

	if err != nil {
		return &utils.UseCaseError{Type: utils.InternalError, Message: "could not update the user"}
	}

	return nil
}
