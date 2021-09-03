package repository

import (
	"errors"
	"go_homework_4/config"
	"go_homework_4/model"

	"gorm.io/gorm"
)

func FindAll() ([]model.Movie, error) {
	var movies []model.Movie

	db, err := config.GormDatabaseConn()

	if err != nil {
		return nil, err
	}

	db.Find(&movies)

	return movies, nil

}

func FindMovieBySlug(slug string) (model.Movie, error) {
	var movie model.Movie

	db, err := config.GormDatabaseConn()

	if err != nil {
		return model.Movie{}, err
	}

	errorNotFound := db.Where("slug = ?", slug).First(&movie).Error
	errors.Is(errorNotFound, gorm.ErrRecordNotFound)

	return movie, errorNotFound
}

func Save(movie model.Movie) (model.Movie, error) {
	db, err := config.GormDatabaseConn()

	if err != nil {
		return model.Movie{}, err
	}

	db.Save(&movie)

	return movie, nil
}

func Delete(movie model.Movie) error {
	db, err := config.GormDatabaseConn()

	if err != nil {
		return err
	}

	db.Delete(&movie)

	return nil
}
