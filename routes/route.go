package routes

import (
	"go_homework_4/handler"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	app.Get("/movie", handler.GetAllMovie)
	app.Get("/movie/:slug", handler.GetSingleMovie)
	app.Post("/movie", handler.CreateNewMovie)
	app.Put("/movie/:slug", handler.UpdateMovieData)
	app.Delete("/movie/:slug", handler.DeleteMovieData)
}
