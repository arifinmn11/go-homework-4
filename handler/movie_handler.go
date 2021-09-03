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
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error":  err,
			"result": nil})
	}

	return ctx.JSON(fiber.Map{
		"error":  nil,
		"result": movies})
}

func GetSingleMovie(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")

	movie, err := service.FindMovieBySlug(slug)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error":  err,
			"result": nil})
	}

	return ctx.JSON(fiber.Map{
		"error":  nil,
		"result": movie})

}

func CreateNewMovie(ctx *fiber.Ctx) error {
	// movie, err := service.CreateMovie()
	payload := new(model.Movie)

	err := ctx.BodyParser(payload)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":  err,
			"result": nil})
	}

	movie, errorInsert := service.CreateMovie(*payload)
	if errorInsert != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":  errorInsert,
			"result": nil})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"error":  nil,
		"result": movie})
}

func UpdateMovieData(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")

	payload := new(model.Movie)

	err := ctx.BodyParser(payload)

	_, errFind := service.FindMovieBySlug(slug)
	if errFind != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error":  errFind,
			"result": nil})
	}

	if err != nil {
		return err
	}

	update, errorUpdate := service.UpdateMovie(slug, *payload)
	if errorUpdate != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":  errorUpdate,
			"result": nil})
	}

	return ctx.JSON(fiber.Map{
		"error":  nil,
		"result": update})
}

func DeleteMovieData(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	_, errorDeleteMovie := service.DeleteMovieBySlug(slug)
	if errorDeleteMovie != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":  errorDeleteMovie,
			"result": nil})
	}
	return ctx.JSON(fiber.Map{
		"error":  nil,
		"result": "Delete movie success"})

}
