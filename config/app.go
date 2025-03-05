package config

import (
	"context"
	"database/sql"
	"go-user-api/internal/delivery/rest"
	"go-user-api/internal/delivery/rest/middleware"
	"go-user-api/internal/delivery/rest/router"
	"go-user-api/internal/repository"
	"go-user-api/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	Viper      *viper.Viper
	DB         *sql.DB
	App        *fiber.App
	Validation *validator.Validate
	Ctx        context.Context
	Log        *logrus.Logger
}

func Bootstrap(config *BootstrapConfig) {
	// repository
	userRepository := repository.NewUserRepository(config.DB, config.Ctx, config.Log)

	// usecase
	authUsecase := usecase.NewAuthUsecase(userRepository, config.Validation, config.Log, config.Viper)
	userUsecase := usecase.NewUserUsecase(userRepository, config.Validation, config.Log)

	// handler
	authHandler := rest.NewAuthHandler(authUsecase)
	userHandler := rest.NewUserhandler(userUsecase)

	// middleware
	middleware := &middleware.MiddlewareConfig{
		App:   config.App,
		Log:   config.Log,
		Viper: config.Viper,
	}
	middleware.Application()

	// router
	route := &router.RouteConfig{
		App:            config.App,
		AuthHandler:    authHandler,
		UserHandler:    userHandler,
		AuthMiddleware: middleware,
	}
	route.Setup()
}
