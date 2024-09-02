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

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {

	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, repository.Wrap(err)
	}

	return user, nil
}

func (r *UserRepository) GetUser(ctx context.Context, id int) (*domain.User, error) {
	dest := domain.User{}
	err := r.db.WithContext(ctx).First(&dest, id).Error
	if err != nil {
		return nil, repository.Wrap(err)
	}

	return &dest, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {

	err := r.db.WithContext(ctx).Save(user).Error
	if err != nil {
		return nil, repository.Wrap(err)
	}

	return user, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) (*domain.User, error) {
	value := domain.User{}
	err := r.db.WithContext(ctx).Delete(&value, id).Error
	if err != nil {
		return nil, repository.Wrap(err)
	}

	return &value, nil
}
