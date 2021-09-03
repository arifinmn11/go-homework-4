package main

import (
	"fmt"
	"go_homework_4/config"
	"go_homework_4/model"
	"go_homework_4/routes"

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
	// app.Get("/users/:id", handler.GetSingleUser)
	// app.Post("/users", handler.CreateNewUser)
	// app.Put("/users/:id", handler.UpdateUserData)
	// app.Delete("/users/:id", handler.DeleteUserData)

	app.Listen(":8081")
}
