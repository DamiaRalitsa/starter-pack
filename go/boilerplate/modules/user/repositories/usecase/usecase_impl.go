package usecase

import (
	"boilerplate/databases/postgresql"
	"boilerplate/modules/user/repositories/commands"
	"boilerplate/modules/user/repositories/queries"
)

type userUsecaseImpl struct {
	repo postgresql.QueryRepository
}

// NewUserUsecase constructor
func NewUserUsecase(repo postgresql.QueryRepository) UserUsecase {
	return &userUsecaseImpl{
		repo: repo,
	}
}

func (uc *userUsecaseImpl) CreateUser(data commands.Users) error {
	// user := &AddUser{}
	// uc.repo.Create(ctx, user)
	uc.repo.Create(data)
	return nil

}

func (uc *userUsecaseImpl) GetAllUser() []queries.Users {
	// fmt.Println("masuk 2")
	result := uc.repo.Get()
	var users []queries.Users

	for _, data := range result {
		user := queries.Users{
			ID:        data.ID,
			Name:      data.Name,
			Title:     data.Title,
			Status:    data.Status,
			CreatedAt: data.CreatedAt,
		}
		// fmt.Println(user)
		users = append(users, user)
	}
	return users

}
