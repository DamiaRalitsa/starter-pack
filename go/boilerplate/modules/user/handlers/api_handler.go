package user

import (
	"boilerplate/helpers"
	commandUser "boilerplate/modules/user/repositories/commands"
	"boilerplate/modules/user/repositories/usecase"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// UserRestHandler handler
type UserRestHandler struct {
	uc usecase.UserUsecase
	// userCommandUsecase commandUser.UserCommandUsecase
	// userQueryUsecase   queryUser.UserQueryUsecase
}

// NewUserRestHandler create new rest handler
func NewUserRestHandler(uc usecase.UserUsecase) *UserRestHandler {
	return &UserRestHandler{uc: uc}
}

// func RouterInit(app *fiber.App) {

// 	api := app.Group("/boilerplate")
// 	// Create User
// 	api.Post("/user", h.CreateUserHandler)
// 	// Get All Users
// 	api.Get("/user", h.GetUserHandler)
// }

// func GetDetailUser(ctx *fiber.Ctx) error {
// 	return ctx.JSON(fiber.Map{
// 		"hello": "world",
// 	})
// }

//Init Router
func RouterInit(h *UserRestHandler) *fiber.App {
	router := fiber.New()
	router.Route("/user", func(router fiber.Router) {
		router.Post("", h.CreateUserHandler)
		router.Get("", h.GetUserHandler)
	})
	return router
}

func (h *UserRestHandler) CreateUserHandler(ctx *fiber.Ctx) error {
	userPayload := commandUser.Users{}
	// user.CreatedAt = time.Now()
	err := ctx.BodyParser(&userPayload)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	// helpers.CheckErrorBodyParse(ctx, err)
	// fmt.Println(user)
	// err = h.userCommandUsecase.CreateUser(ctx)
	// fmt.Println(err)
	// if err != nil {
	// 	return err
	// }
	// helpers.CheckError(err)

	h.uc.CreateUser(userPayload)

	finalResponse := helpers.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Success Create User Data!",
		Data:    nil,
	}
	ctx.Status(http.StatusOK).JSON(finalResponse)

	// ctx.Status(http.StatusOK).JSON(&fiber.Map{
	// 	"message": "Success Create User Data!"})
	return nil
}

func (h *UserRestHandler) GetUserHandler(ctx *fiber.Ctx) error {
	// user := &[]queryUser.User{}
	// err := h.userQueryUsecase.GetAllUser(ctx)
	// fmt.Println(err)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("masuk 1")
	response := h.uc.GetAllUser()
	// fmt.Println(response)
	finalResponse := helpers.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Success Get User Data!",
		Data:    response,
	}
	ctx.Status(http.StatusOK).JSON(finalResponse)
	return nil
}
