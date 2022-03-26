package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "swagger-example/docs"
	"github.com/gofiber/fiber/v2"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	// app.Get("/swagger/*", fiberSwagger.New(fiberSwagger.Config{ // custom
	// 	URL: "http://localhost:8080/swagger/doc.json",
	// 	DeepLinking: false,
	// }))

	app.Get("api/users/:id", ShowUser)

	app.Listen(":8080")
}

// ShowUser godoc
// @Summary Show a user
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/users/{id} [get]
func ShowUser(c *fiber.Ctx) error {
	return c.JSON(User{
		Id: c.Params("id"),
	})
}

type User struct {
	Id string
}

type HTTPError struct {
	status  string
	message string
}