package repository

import (
	"context"
	"toktok-backend-v2/internal/core/domain"
	"toktok-backend-v2/internal/database"
)

type UserRepository struct {
	db *database.Database
}

func NewUserRepository(db *database.Database) domain.UserRepository {
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
		return nil, err
	}

	return &value, nil
}

func (r *UserRepository) GetUser(ctx context.Context, id int) (*domain.User, error) {
	dest := domain.User{}
	err := r.db.WithContext(ctx).First(&dest, id).Error
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return &value, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) (*domain.User, error) {
	value := domain.User{}
	err := r.db.WithContext(ctx).Delete(&value, id).Error
	if err != nil {
		return nil, err
	}

	return &value, nil
}
