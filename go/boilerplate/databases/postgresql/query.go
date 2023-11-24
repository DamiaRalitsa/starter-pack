package postgresql

import (
	"boilerplate/modules/user/repositories/commands"
	"boilerplate/modules/user/repositories/queries"
)

type QueryRepository interface {
	Create(data commands.Users) error
	Get() []queries.Users
	GetById(id int) (queries.Users, error)
}
