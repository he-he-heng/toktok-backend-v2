package usecase

import (
	"context"
	"errors"
	"toktok-backend-v2/internal/core/domain"
	"toktok-backend-v2/internal/core/user/utils"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	userUsecase := userUsecase{
		userRepository: userRepository,
	}

	return &userUsecase
}

func (u *userUsecase) CreateUser(ctx context.Context, loginID string, password string) (*domain.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	return u.userRepository.CreateUser(ctx, loginID, hashedPassword)
}

func (u *userUsecase) GetUser(ctx context.Context, id int) (*domain.User, error) {
	return u.userRepository.GetUser(ctx, id)
}

func (u *userUsecase) UpdateUser(ctx context.Context, id int, loginID *string, password *string, confirmPassword *string, email *string) (*domain.User, error) {
	queriedUser, err := u.userRepository.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	if loginID != nil || password != nil {
		if confirmPassword == nil {
			return nil, errors.New("confirmPassword가 존재하지 않습니다")
		}
	}

	// 기존 패스워드와 일치하는지 확인
	err = utils.VerifyPassword(*confirmPassword, queriedUser.Password)
	if err != nil {
		return nil, err
	}

	hashedPassword := ""
	if password != nil {
		hash, err := utils.HashPassword(*password)
		if err != nil {
			return nil, err
		}
		hashedPassword = hash
	}

	updatedUser, err := u.userRepository.UpdateUser(ctx, loginID, &hashedPassword, email)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u *userUsecase) DeleteUser(ctx context.Context, id int) (*domain.User, error) {
	return u.userRepository.DeleteUser(ctx, id)
}
