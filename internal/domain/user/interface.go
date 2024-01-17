package user

type UserRepositoryInterface interface {
	Create(user *User) error
}
