package usecase

import (
	"boilerplate/modules/user/repositories/commands"
	"boilerplate/modules/user/repositories/queries"
)

// UserUsecase abstraction
type UserUsecase interface {
	CreateUser(data commands.Users) error
	GetAllUser() []queries.Users
}
