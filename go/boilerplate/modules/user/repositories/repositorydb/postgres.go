package repositorydb

// type Repository struct {
// 	DB *gorm.DB
// }

// func (r *Repository) CreateUserData(ctx *fiber.Ctx) error {
// 	user := commands.AddUser{}
// 	err := r.DB.Create(user).Error
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

// func (r *Repository) GetUserData(ctx *fiber.Ctx) error {
// 	user := []queries.User{}
// 	err := r.DB.Find(user).Error
// 	helpers.CheckErrorBadRequest(ctx, err)

// 	ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"message": "data fetched successfully",
// 		"data":    user})
// 	return nil

// }

// func GetAllDataUsers() {
// rows, err := DB.Query(`SELECT "Name", "Roll_Number" FROM "Students"`)
// helpers.CheckError(err)

// defer rows.Close()
// for rows.Next() {
// 	var name string
// 	var roll_number int

// 	err = rows.Scan(&name, &roll_number)
// 	helpers.CheckError(err)

// 	fmt.Println(name, roll_number)
// }

// helpers.CheckError(err)
// }
