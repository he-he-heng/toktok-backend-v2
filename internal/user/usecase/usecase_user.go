package usecase

import (
	"context"
	"toktok-backend-v2/internal/domain"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func (u *userUsecase) CreateUser(ctx context.Context, loginID string, password string) (*domain.User, error) {
	u.userRepository.CreateUser(ctx, loginID, hashedPassword)
}

func (u *userUsecase) GetUser(ctx context.Context, id int) (*domain.User, error) {}

func (u *userUsecase) UpdateUser(ctx context.Context, loginID *string, hashedPassword *string, confirmPassword *string, email *string) (*domain.User, error) {
}
func (u *userUsecase) DeleteUser(ctx context.Context, id int) (*domain.User, error) {}
