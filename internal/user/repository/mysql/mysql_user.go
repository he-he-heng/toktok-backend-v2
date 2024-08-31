package mysql

import (
	"toktok-backend-v2/internal/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	userRepository := userRepository{
		db: db,
	}

	return &userRepository
}

func (r *userRepository) CreateUser(loginID string, hashedPassword string) (*domain.User, error) {
	value := domain.User{
		LoginID:  loginID,
		Password: hashedPassword,
	}

	err := r.db.Create(&value).Error
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func (r *userRepository) GetUser(id int) (*domain.User, error) {
	dest := domain.User{}
	err := r.db.First(&dest, id).Error
	if err != nil {
		return nil, err
	}

	return &dest, nil
}

func (r *userRepository) UpdateUser(loginID *string, hashedPassword *string, email *string) (*domain.User, error) {
	value := domain.User{}

	if loginID != nil {
		value.LoginID = *loginID
	}

	if hashedPassword != nil {
		value.Password = *hashedPassword
	}

	if email != nil {
		value.Email = email
	}

	err := r.db.Save(&value).Error
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func (r *userRepository) DeleteUser(id int) (*domain.User, error) {
	value := domain.User{}
	err := r.db.Delete(&value, id).Error
	if err != nil {
		return nil, err
	}

	return &value, nil
}
