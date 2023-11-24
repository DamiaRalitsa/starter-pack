package postgresql

import (
	"boilerplate/helpers"
	"boilerplate/modules/user/repositories/commands"
	"boilerplate/modules/user/repositories/queries"
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepositoryImpl(Db *gorm.DB) QueryRepository {
	return &Repository{Db: Db}
}

func Connect(config *Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	// 	// open database
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.CheckError(err)

	fmt.Println("Connected!")

	// defer func() {
	// 	// close database
	// 	DbInstance, _ := DB.DB()
	// 	DbInstance.Close()
	// }()

	return DB, nil

	// 	// close database
	// 	defer DB.Close()

	// 	// check db
	// 	err = DB.Ping()
	// 	helpers.CheckError(err)

	// 	fmt.Println("Connected!")
}

func (r *Repository) Create(data commands.Users) error {

	result := r.Db.Create(&data)
	if result != nil {
		// fmt.Println(data)
		return nil
	}
	return errors.New("failed to create data")

	// 	err := r.DB.Create(data).Error
	// 	// helpers.CheckErrorBadRequest(ctx, err)
	// 	if err != nil {
	// 		ctx.Status(http.StatusBadRequest).JSON(
	// 			&fiber.Map{"message": "could not create data"})
	// 		return err
	// 	}

	// 	ctx.Status(http.StatusOK).JSON(&fiber.Map{
	// 		"message": "data has been created successfully"})

	return nil

}

func (r *Repository) Get() []queries.Users {
	// fmt.Println("masuk 3")
	var users []queries.Users
	r.Db.Find(&users)
	// fmt.Println(users)
	return users
}

func (r *Repository) GetById(id int) (queries.Users, error) {
	// fmt.Println("masuk 4")
	var users queries.Users
	result := r.Db.Find(&users, id)
	if result != nil {
		// fmt.Println(users)
		return users, nil
	}
	return users, errors.New("user not found")
}

// func (r *Repository) Create(ctx *fiber.Ctx, data interface{}) error {

// 	err := r.DB.Create(data).Error
// 	// helpers.CheckErrorBadRequest(ctx, err)
// 	if err != nil {
// 		ctx.Status(http.StatusBadRequest).JSON(
// 			&fiber.Map{"message": "could not create data"})
// 		return err
// 	}

// 	ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"message": "data has been created successfully"})
// 	return nil

// }

// func (r *Repository) Get(ctx *fiber.Ctx, data []interface{}) error {

// 	err := r.DB.Find(data).Error
// 	helpers.CheckErrorBadRequest(ctx, err)

// 	ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"message": "data fetched successfully",
// 		"data":    data})
// 	return nil

// }

// func (r *Repository) GetAllDataUsers(ctx *fiber.Ctx) error {
// 	rows, err := r.DB.Exec(`SELECT * FROM users`)
// 	helpers.CheckError(err)

// 	defer rows.Close()
// 	for rows.Next() {
// 		var user queries.User

// 		err = rows.Scan(user)
// 		helpers.CheckError(err)

// 		fmt.Println(name, roll_number)

// 		ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 			"message": "data fetched successfully",
// 			"data":    data})
// 		return nil
// 	}

// 	helpers.CheckError(err)
// }
