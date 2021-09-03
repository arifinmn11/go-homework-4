package handler

import (
	"go_homework_4/model"
	"go_homework_4/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetAllMovie(ctx *fiber.Ctx) error {
	movies, err := service.FindAllMovie()

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"status": 200,
		"data":   movies})
}

func GetSingleMovie(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")

	movie, err := service.FindMovieBySlug(slug)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{
		"data": movie})

}

func CreateNewMovie(ctx *fiber.Ctx) error {
	// movie, err := service.CreateMovie()
	payload := new(model.Movie)

	err := ctx.BodyParser(payload)

	if err != nil {
		return err
	}

	movie, errorInsert := service.CreateMovie(*payload)
	if errorInsert != nil {
		return errorInsert
	}

	return ctx.JSON(fiber.Map{
		"inserted_record": movie})
}

func UpdateMovieData(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")

	payload := new(model.Movie)

	err := ctx.BodyParser(payload)

	if err != nil {
		return err
	}

	update, errorUpdate := service.UpdateMovie(slug, *payload)
	if errorUpdate != nil {
		return errorUpdate
	}

	return ctx.JSON(fiber.Map{
		"updated_record": update})
}

func DeleteMovieData(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	_, errorDeleteMovie := service.DeleteMovieBySlug(slug)
	if errorDeleteMovie != nil {
		return errorDeleteMovie
	}
	return ctx.JSON(fiber.Map{
		"message": "Delete movie success"})

}
