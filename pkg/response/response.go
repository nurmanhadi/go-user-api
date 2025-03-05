package response

import (
	"errors"
	"fmt"
	"go-user-api/pkg/exception"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Success(ctx *fiber.Ctx, statusCode int, data any) error {
	if data == nil {
		return ctx.Status(statusCode).JSON(fiber.Map{
			"links": fiber.Map{
				"self": ctx.OriginalURL(),
			},
		})
	}
	return ctx.Status(statusCode).JSON(fiber.Map{
		"data": data,
		"links": fiber.Map{
			"self": ctx.OriginalURL(),
		},
	})
}
func Error(ctx *fiber.Ctx, err error) error {
	if validatorErr, ok := err.(validator.ValidationErrors); ok {
		var values []string
		for _, fieldErr := range validatorErr {
			value := fmt.Sprintf("field %s is %s %s", fieldErr.Field(), fieldErr.Tag(), fieldErr.Param())
			values = append(values, value)
		}
		str := strings.Join(values, ", ")
		return ErrorR(ctx, 400, str)
	} else if errors.Is(err, exception.EmailAlreadyExists) {
		return ErrorR(ctx, 409, err.Error())
	} else if errors.Is(err, exception.EmailOrPasswordWrong) {
		return ErrorR(ctx, 400, err.Error())
	} else if errors.Is(err, exception.UserNotFound) {
		return ErrorR(ctx, 404, err.Error())
	} else if errors.Is(err, exception.PasswordWrong) {
		return ErrorR(ctx, 400, err.Error())
	}
	return ErrorR(ctx, 500, "internal server error")
}
func ErrorR(ctx *fiber.Ctx, statusCode int, err string) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"error": err,
		"links": fiber.Map{
			"self": ctx.OriginalURL(),
		},
	})
}
