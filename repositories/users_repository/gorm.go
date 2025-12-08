package users_repository

import (
	"context"
	"errors"
	"socialmedia/domain"

	"gorm.io/gorm"
)

type GormUsersRepository struct {
	DB  *gorm.DB
	ctx context.Context
}

func NewGormUsersRepository(db *gorm.DB) *GormUsersRepository {
	return &GormUsersRepository{
		DB:  db,
		ctx: context.Background(),
	}
}

func (r *GormUsersRepository) CreateUser(data domain.Users) (domain.Users, error) {
	err := r.DB.WithContext(r.ctx).Create(&data).Error
	if err != nil {
		return domain.Users{}, err
	}

	return data, nil
}

func (r *GormUsersRepository) GetAllUsers(page, limit int) ([]domain.Users, error) {

	var users []domain.Users
	err := r.DB.WithContext(r.ctx).Order("username ASC").Offset((page - 1) * limit).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *GormUsersRepository) GetUserByID(id string) (domain.Users, error) {

	var user domain.Users
	err := r.DB.WithContext(r.ctx).Where("id=?", id).First(&user).Error
	if err != nil {
		return domain.Users{}, err
	}

	return user, nil
}

func (r *GormUsersRepository) UpdateUser(data domain.Users) error {
	row := r.DB.WithContext(r.ctx).Where("id=?", data.ID).Updates(data)
	if err := row.Error; err != nil {
		return err
	}

	if row.RowsAffected == 0 {
		return errors.New("no rows affected, user doesnt exists")
	}

	return nil
}

func (r *GormUsersRepository) DeleteUser(id string) error {
	row := r.DB.WithContext(r.ctx).Delete(domain.Users{})
	if err := row.Error; err != nil {
		return err
	}

	if row.RowsAffected == 0 {
		return errors.New("no rows affected, user doesnt exists")
	}

	return nil
}

func (r *GormUsersRepository) FindByEmail(email string) (domain.Users, error) {
	var user domain.Users

	err := r.DB.WithContext(r.ctx).Where("email=?", email).First(&user).Error
	if err != nil {
		return domain.Users{}, err
	}

	return user, nil
}
