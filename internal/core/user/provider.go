package user

import (
	"toktok-backend-v2/internal/core/domain"
	"toktok-backend-v2/internal/core/user/delivery"
	"toktok-backend-v2/internal/core/user/repository"
	"toktok-backend-v2/internal/core/user/usecase"
	"toktok-backend-v2/internal/database"
)

type userProvider struct {
	UserRepository domain.UserRepository
	UserUsecase    domain.UserUsecase
	UserDelivery   delivery.UserDelivery
}

func Provide(db *database.Database) *userProvider {
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userDelivery := delivery.NewUserDelivery(userUsecase)

	userProvider := userProvider{
		UserRepository: userRepository,
		UserUsecase:    userUsecase,
		UserDelivery:   userDelivery,
	}

	return &userProvider
}
