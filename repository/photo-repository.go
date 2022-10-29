package repository

import (
	"task5/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	InsertPhoto(b models.Photo) models.Photo
	UpdatePhoto(b models.Photo) models.Photo
	DeletePhoto(b models.Photo)
	AllPhoto() []models.Photo
	FindPhotoByID(PhotoID uint64) models.Photo
}

type photoConnection struct {
	connection *gorm.DB
}

func NewPhotoRepository(dbConn *gorm.DB) PhotoRepository {
	return &photoConnection{
		connection: dbConn,
	}
}

func (db *photoConnection) InsertPhoto(b models.Photo) models.Photo {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *photoConnection) UpdatePhoto(b models.Photo) models.Photo {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *photoConnection) DeletePhoto(b models.Photo) {
	db.connection.Delete(&b)
}

func (db *photoConnection) FindPhotoByID(photoID uint64) models.Photo {
	var photo models.Photo
	db.connection.Preload("User").Find(&photo, photoID)
	return photo
}

func (db *photoConnection) AllPhoto() []models.Photo {
	var photos []models.Photo
	db.connection.Preload("User").Find(&photos)
	return photos
}
