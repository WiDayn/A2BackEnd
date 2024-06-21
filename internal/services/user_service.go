package services

import (
	"A2BackEnd/internal/models"
	"A2BackEnd/internal/repositories"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService interface {
	RegisterUser(username, password string) error
	LoginUser(username, password string) (string, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) RegisterUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &models.User{Username: username, Password: string(hashedPassword)}
	return s.userRepository.CreateUser(user)
}

func (s *userService) LoginUser(username, password string) (string, error) {
	user, err := s.userRepository.FindUserByUsername(username)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}
	token, err := generateJWTToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func generateJWTToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret")) // Replace "secret" with your actual secret key
}
