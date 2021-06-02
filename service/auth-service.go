package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/sing3demons/golanh-api/dto"
	"github.com/sing3demons/golanh-api/entity"
	"github.com/sing3demons/golanh-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{userRepository: userRepository}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	resp := service.userRepository.VerifyCredential(email, password)
	if v, ok := resp.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return resp
		}
		return false
	}
	return false
}
func (service *authService) CreateUser(user dto.RegisterDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}
func (service *authService) FindByEmail(email string) entity.User {
	var user entity.User
	return user
}
func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hash string, password []byte) bool {
	byteHash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(byteHash, password)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
