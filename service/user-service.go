package service

import (
	"log"
	"task5/dto"
	"task5/models"
	"task5/repository"

	"github.com/mashingan/smapping"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) models.User
	Profile(userID string) models.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) models.User {
	userToUpdate := models.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser((userToUpdate))
	return updatedUser
}

func (service *userService) Profile(userID string) models.User {
	return service.userRepository.ProfileUser(userID)
}
