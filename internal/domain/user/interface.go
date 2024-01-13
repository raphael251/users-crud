package user

type UserRepositoryInterface interface {
	Save(user *User) error
}
