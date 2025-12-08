package users_service

import (
	"errors"
	"socialmedia/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	usersRepo UsersRepo
	jwtSecret string
}

type Service interface {
	Register(data domain.Users) (domain.Users, error)
	Login(email, password string) (string, error)
	GetAllUsers(page, limit int) ([]domain.Users, error)
	GetUserByID(id string) (domain.Users, error)
}

func NewUsersService(usersRepo UsersRepo, jwtSecret string) Service {
	return &UsersService{
		usersRepo: usersRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *UsersService) Register(data domain.Users) (domain.Users, error) {
	isExists, _ := s.usersRepo.FindByEmail(data.Email)
	if isExists.Email != "" {
		return domain.Users{}, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.Users{}, err
	}

	data.ID = uuid.NewString()
	data.Password = string(hashedPassword)
	data.Role = "user"

	return s.usersRepo.CreateUser(data)
}
func (s *UsersService) Login(email, password string) (string, error) {
	isExists, err := s.usersRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(isExists.Password), []byte(password))
	if err != nil {
		return "", err
	}

	type MapClaims struct {
		ID   string `json:"id"`
		Role string `json:"role"`
		*jwt.RegisteredClaims
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MapClaims{
		ID:   isExists.ID,
		Role: isExists.Role,
		RegisteredClaims: &jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})

	signedString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return signedString, nil
}
func (s *UsersService) GetAllUsers(page, limit int) ([]domain.Users, error) {
	return s.usersRepo.GetAllUsers(page, limit)
}
func (s *UsersService) GetUserByID(id string) (domain.Users, error) {
	return s.usersRepo.GetUserByID(id)
}
