package middleware

import (
	"errors"
	"go-user-api/pkg/response"
	"go-user-api/pkg/token"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type MiddlewareConfig struct {
	App   *fiber.App
	Log   *logrus.Logger
	Viper *viper.Viper
}

func (m *MiddlewareConfig) Application() {
	// monitor
	m.App.Get("/metrics", monitor.New(monitor.Config{Title: "Go User API"}))

	// logger
	// m.App.Use(logger.New(logger.Config{
	// 	Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	// }))

	// compress
	m.App.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
}
func (m *MiddlewareConfig) Auth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString, err := getTokenFromHeader(ctx)
		if err != nil {
			m.Log.WithField("error", err).Warn("failed get token from header")
			return response.ErrorR(ctx, 401, err.Error())
		}
		jwt, err := token.VerifyToken(tokenString, m.Viper)
		if err != nil {
			m.Log.WithField("error", err).Warn("failed verify jwt token")
			return response.ErrorR(ctx, 401, err.Error())
		}
		ctx.Locals("id", jwt.Id)
		return ctx.Next()
	}
}
func getTokenFromHeader(c *fiber.Ctx) (string, error) {
	header := c.Get("Authorization")
	if header == "" {
		return "", errors.New("null token Authorization")
	}
	parts := strings.Split(header, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid token format")
	}
	return parts[1], nil
}
