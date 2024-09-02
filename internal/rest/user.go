package rest

import (
	"context"
	"toktok-backend-v2/domain"
	"toktok-backend-v2/internal/rest/dto"
	"toktok-backend-v2/internal/rest/validator"
	"toktok-backend-v2/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUser(ctx context.Context, id int) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id int) (*domain.User, error)
}

type UserHandler struct {
	userService UserService
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	dto := dto.CreateUserReqeust{}
	if err := c.BodyParser(&dto); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	if err := validator.Get().Validate(&dto); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	_, err := h.userService.CreateUser(c.Context(), dto.ToDomain())
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}
