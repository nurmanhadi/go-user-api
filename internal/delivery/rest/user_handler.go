package rest

import (
	"go-user-api/internal/usecase"
	"go-user-api/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type IUserHandler interface {
	FindById(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
type userHandler struct {
	userUsecase usecase.IUserUsecase
}

func NewUserhandler(userUsecase usecase.IUserUsecase) IUserHandler {
	return &userHandler{userUsecase: userUsecase}
}
func (h *userHandler) FindById(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.userUsecase.FindById(id)
	if err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, 200, result)
}
func (h *userHandler) Delete(c *fiber.Ctx) error {
	userId := c.Locals("id").(string)
	id := c.Params("id")
	if userId != id {
		return response.ErrorR(c, 401, "You are not authorized to access this resource")
	}
	err := h.userUsecase.Delete(id)
	if err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, 200, nil)
}
