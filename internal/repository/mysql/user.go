package mysql

import (
	"context"
	"toktok-backend-v2/domain"
	"toktok-backend-v2/internal/repository"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	userRepository := UserRepository{
		db: db,
	}

	return &userRepository
}

func (r *UserRepository) CreateUser(ctx context.Context, loginID string, hashedPassword string) (*domain.User, error) {
	value := domain.User{
		LoginID:  loginID,
		Password: hashedPassword,
	}

	err := r.db.WithContext(ctx).Create(&value).Error
	if err != nil {
		return nil, repository.Wrap(err)
	}

	return &value, nil
}

func (r *UserRepository) GetUser(ctx context.Context, id int) (*domain.User, error) {
	dest := domain.User{}
	err := r.db.WithContext(ctx).First(&dest, id).Error
	if err != nil {
		return nil, repository.Wrap(err)
	}

	return &dest, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, loginID *string, hashedPassword *string, email *string) (*domain.User, error) {
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

	err := r.db.WithContext(ctx).Save(&value).Error
	if err != nil {
		return nil, repository.Wrap(err)
	}

	return &value, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) (*domain.User, error) {
	value := domain.User{}
	err := r.db.WithContext(ctx).Delete(&value, id).Error
	if err != nil {
		return nil, repository.Wrap(err)
	}

	return &value, nil
}
