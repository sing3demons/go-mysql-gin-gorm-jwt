package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/sing3demons/golanh-api/dto"
	"github.com/sing3demons/golanh-api/entity"
	"github.com/sing3demons/golanh-api/repository"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	var userToUpdate entity.User
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	updateUser := service.userRepository.UpdateUser(userToUpdate)
	return updateUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}
