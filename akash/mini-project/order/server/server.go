package server

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"order.akash.com/db"
)

const (
	Port = "7779"
)

var (
	repo db.OrderRepository = db.NewMongoRepository()
)

func Start(repository db.OrderRepository) {

	repo = repository

	app := fiber.New()

	app.Get("/order/all", func(c *fiber.Ctx) error {
		log.Println("/order/all request received")
		return c.JSON(repo.FindAll())
	})

	log.Fatal(app.Listen(":" + Port))
}
