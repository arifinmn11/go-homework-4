package routes

import (
	"go_homework_4/handler"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	app.Get("/movies", handler.GetAllMovie)
	app.Get("/movies/:slug", handler.GetSingleMovie)
	app.Post("/movies", handler.CreateNewMovie)
	app.Put("/movies/:id", handler.UpdateMovieData)
	app.Delete("/movies/:id", handler.DeleteMovieData)
}
