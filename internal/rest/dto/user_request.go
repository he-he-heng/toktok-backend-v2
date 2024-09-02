package dto

import "toktok-backend-v2/domain"

type CreateUserReqeust struct {
	LoginID  string `json:"uid" validate:"gte=6,lte=18"`
	Password string `json:"password" validate:"gte=6,lte=32"`
}

func (dto *CreateUserReqeust) ToDomain() *domain.User {
	user := domain.User{
		LoginID:  dto.LoginID,
		Password: dto.Password,
	}

	return &user
}
