package delivery

import (
	"toktok-backend-v2/internal/core/domain"
	"toktok-backend-v2/internal/core/user/delivery/dto"
	"toktok-backend-v2/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

type userDelivery struct {
	userUsecase domain.UserUsecase
}

func NewUserDelivery(userUsecase domain.UserUsecase) userDelivery {
	userDelivery := userDelivery{
		userUsecase: userUsecase,
	}

	return userDelivery
}

func (d *userDelivery) CreateUser(c *fiber.Ctx) error {
	body := dto.CreateUserRequest{}
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	if err := validator.Get().Validate(&body); err != nil {
		return err
	}

	_, err := d.userUsecase.CreateUser(c.Context(), body.LoginID, body.Password)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}
