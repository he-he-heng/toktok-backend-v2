package rest

import (
	"context"
	"toktok-backend-v2/domain"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUser(ctx context.Context, id int) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id int) (*domain.User, error)
}

type UserHandler struct {
	userHandler UserService
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {

}
