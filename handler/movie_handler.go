package handler

import (
	"go_homework_4/model"
	"go_homework_4/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NewString(s string) *string {
	return &s
}

func GetAllMovie(ctx *fiber.Ctx) error {
	movies, err := service.FindAllMovie()

	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(model.ResponseFormmater{
			Error:  NewString(err.Error()),
			Result: nil})
	}

	return ctx.JSON(model.ResponseFormmater{
		Error:  nil,
		Result: movies})
}

func GetSingleMovie(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")

	movie, err := service.FindMovieBySlug(slug)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(model.ResponseFormmater{
			Error:  NewString(err.Error()),
			Result: nil})
	}

	return ctx.JSON(model.ResponseFormmater{
		Error:  nil,
		Result: movie})

}

func CreateNewMovie(ctx *fiber.Ctx) error {
	// movie, err := service.CreateMovie()
	payload := new(model.Movie)

	err := ctx.BodyParser(payload)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(model.ResponseFormmater{
			Error:  NewString(err.Error()),
			Result: nil})
	}

	movie, errorInsert := service.CreateMovie(*payload)
	if errorInsert != nil {
		return ctx.Status(http.StatusBadRequest).JSON(model.ResponseFormmater{
			Error:  NewString(errorInsert.Error()),
			Result: nil})
	}

	return ctx.Status(http.StatusCreated).JSON(model.ResponseFormmater{
		Error:  nil,
		Result: movie})
}

func UpdateMovieData(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")

	payload := new(model.Movie)

	err := ctx.BodyParser(payload)

	_, errFind := service.FindMovieBySlug(slug)
	if errFind != nil {
		return ctx.Status(http.StatusNotFound).JSON(model.ResponseFormmater{
			Error:  NewString(errFind.Error()),
			Result: nil})
	}

	if err != nil {
		return err
	}

	update, errorUpdate := service.UpdateMovie(slug, *payload)
	if errorUpdate != nil {
		return ctx.Status(http.StatusBadRequest).JSON(model.ResponseFormmater{
			Error:  NewString(errorUpdate.Error()),
			Result: nil})
	}

	return ctx.JSON(model.ResponseFormmater{
		Error:  nil,
		Result: update})
}

func DeleteMovieData(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	_, errorDeleteMovie := service.DeleteMovieBySlug(slug)
	if errorDeleteMovie != nil {
		return ctx.Status(http.StatusBadRequest).JSON(model.ResponseFormmater{
			Error:  NewString(errorDeleteMovie.Error()),
			Result: nil})
	}
	return ctx.JSON(model.ResponseFormmater{
		Error:  nil,
		Result: "Delete movie success"})

}
