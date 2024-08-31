package port

import (
	"context"
	"toktok-backend-v2/internal/core/domain"
	"toktok-backend-v2/internal/core/domain/frame"
)

type UserService interface {
	CreateUser(ctx context.Context, frame frame.CreateUser) error
	GetUser(ctx context.Context, frame frame.GetUser) (domain.User, error)
	UpdateUser(ctx context.Context, frame frame.UpdateUser) error
	DeleteUser(ctx context.Context, frame frame.DeleteUser) error
}

type UserRepository interface {
	CreateUser(ctx context.Context, frame frame.CreateUser) error
	GetUser(ctx context.Context, frame frame.GetUser) (domain.User, error)
	UpdateUser(ctx context.Context, frame frame.UpdateUser) error
	DeleteUser(ctx context.Context, frame frame.DeleteUser) error
}
