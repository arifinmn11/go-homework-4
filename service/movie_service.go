package service

import (
	"go_homework_4/model"
	"go_homework_4/repository"
)

func FindAllMovie() ([]model.Movie, error) {
	result, err := repository.FindAll()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func FindMovieBySlug(slug string) (model.Movie, error) {
	result, err := repository.FindMovieBySlug(slug)
	if err != nil {
		return model.Movie{}, err
	}
	return result, nil
}

func CreateMovie(movie model.Movie) (model.Movie, error) {
	inserted, err := repository.Save(movie)

	if err != nil {
		return model.Movie{}, err
	}

	return inserted, nil
}

func UpdateMovie(slug string, payload model.Movie) (model.Movie, error) {
	movie, errBySlug := repository.FindMovieBySlug(slug)
	if errBySlug != nil {
		return model.Movie{}, errBySlug
	}

	movie.Image = payload.Image
	movie.Title = payload.Title
	movie.Slug = payload.Slug

	update, errUpdate := repository.Save(movie)
	if errUpdate != nil {
		return model.Movie{}, errBySlug
	}

	return update, nil
}

func DeleteMovieBySlug(slug string) (interface{}, error) {
	movie, errBySlug := repository.FindMovieBySlug(slug)
	if errBySlug != nil {
		return nil, errBySlug
	}

	errDelete := repository.Delete(movie)
	if errDelete != nil {
		return nil, errDelete
	}

	return nil, nil
}
