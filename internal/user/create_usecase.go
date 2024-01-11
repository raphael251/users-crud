package user

import "time"

type CreateUserInputDTO struct {
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Address   string    `json:"address"`
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
	user, err := NewUser(
		input.Name,
		input.BirthDate, input.Email, input.Password, input.Address)

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
