package service

import (
	"fmt"
	"task5/dto"
	"task5/models"
	"task5/repository"

	"log"

	"github.com/mashingan/smapping"
)

type PhotoService interface {
	Insert(b dto.PhotoCreateDTO) models.Photo
	Update(b dto.PhotoUpdateDTO) models.Photo
	Delete(b models.Photo)
	All() []models.Photo
	FindPhotoByID(photoID uint64) models.Photo
	IsAllowedToEdit(userID string, photoID uint64) bool
}

type photoService struct {
	photoRepository repository.PhotoRepository
}

func NewPhotoService(photoRepo repository.PhotoRepository) PhotoService {
	return &photoService{
		photoRepository: photoRepo,
	}
}

func (service *photoService) Insert(b dto.PhotoCreateDTO) models.Photo {
	photo := models.Photo{}
	err := smapping.FillStruct(&photo, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.photoRepository.InsertPhoto(photo)
	return res
}

func (service *photoService) Update(b dto.PhotoUpdateDTO) models.Photo {
	photo := models.Photo{}
	err := smapping.FillStruct(&photo, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.photoRepository.UpdatePhoto(photo)
	return res
}

func (service *photoService) Delete(b models.Photo) {
	service.photoRepository.DeletePhoto(b)
}

func (service *photoService) All() []models.Photo {
	return service.photoRepository.AllPhoto()
}

func (service *photoService) FindPhotoByID(photoID uint64) models.Photo {
	return service.photoRepository.FindPhotoByID(photoID)
}

func (service *photoService) IsAllowedToEdit(userID string, photoID uint64) bool {
	b := service.photoRepository.FindPhotoByID(photoID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
