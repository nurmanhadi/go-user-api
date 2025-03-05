package rest

import (
	"go-user-api/internal/model"
	"go-user-api/internal/usecase"
	"go-user-api/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type IAuthHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	ChangePassword(c *fiber.Ctx) error
}
type authHandler struct {
	authUsecase usecase.IAuthUsecase
}

func NewAuthHandler(authUsecase usecase.IAuthUsecase) IAuthHandler {
	return &authHandler{authUsecase: authUsecase}
}
func (h *authHandler) Register(c *fiber.Ctx) error {
	request := new(model.AuthRegisterRequest)
	err := c.BodyParser(&request)
	if err != nil {
		return response.ErrorR(c, 400, "failed parse json")
	}
	err = h.authUsecase.Register(request)
	if err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, 201, nil)
}
func (h *authHandler) Login(c *fiber.Ctx) error {
	request := new(model.AuthLoginRequest)
	err := c.BodyParser(&request)
	if err != nil {
		return response.ErrorR(c, 400, "failed parse json")
	}
	result, err := h.authUsecase.Login(request)
	if err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, 200, result)
}
func (h *authHandler) ChangePassword(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return response.ErrorR(c, 401, "You are not authorized to access this resource")
	}
	request := new(model.AuthChangePasswordRequest)
	err := c.BodyParser(&request)
	if err != nil {
		return response.ErrorR(c, 400, "failed parse json")
	}
	err = h.authUsecase.ChangePassword(id, request)
	if err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, 200, nil)
}
