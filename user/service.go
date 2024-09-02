package user

import (
	"context"
	"toktok-backend-v2/domain"
	"toktok-backend-v2/pkg/errors"
	"toktok-backend-v2/utils"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUser(ctx context.Context, id int) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id int) (*domain.User, error)
}

type Service struct {
	userRepository UserRepository
}

func NewService(userRepository UserRepository) *Service {
	service := Service{
		userRepository: userRepository,
	}

	return &service
}

func (s *Service) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, errors.Wrap(domain.ErrInternalServerError, err)
	}
	user.Password = hashedPassword

	createdUser, err := s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
