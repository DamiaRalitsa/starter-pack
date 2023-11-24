package server

import (
	"boilerplate/databases/postgresql"
	handlerUser "boilerplate/modules/user/handlers"
	"boilerplate/modules/user/repositories/usecase"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Start() {
	// app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &postgresql.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	// database
	db, err := postgresql.Connect(config)

	// init repo
	userRepo := postgresql.NewRepositoryImpl(db)

	// init usecase
	userUsecase := usecase.NewUserUsecase(userRepo)

	// init handlers
	userHandler := handlerUser.NewUserRestHandler(userUsecase)

	//init routes
	router := handlerUser.RouterInit(userHandler)

	app := fiber.New()

	app.Mount("/boilerplate", router)

	// commandUser.MigrateUser(db)
	// helpers.CheckError(err)

	log.Fatal(app.Listen(":9000"))

}
