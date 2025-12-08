package users_service

import "socialmedia/domain"

type UsersRepo interface {
	CreateUser(data domain.Users) (domain.Users, error)
	GetAllUsers(page, limit int) ([]domain.Users, error)
	GetUserByID(id string) (domain.Users, error)
	UpdateUser(data domain.Users) error
	DeleteUser(id string) error
	FindByEmail(email string) (domain.Users, error)
}
