package main

import (
	"fmt"
	"go_homework_4/config"
	"go_homework_4/model"
	"go_homework_4/routes"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error", err.Error())
	}

	db.AutoMigrate(&model.Movie{})

	app := fiber.New()

	routes.Route(app)
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON(model.ResponseFormmater{
			Error:  &fiber.ErrNotFound.Message,
			Result: nil,
		})
	})

	app.Listen(":" + os.Getenv("APP_PORT"))
}
