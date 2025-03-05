package router

import (
	"go-user-api/internal/delivery/rest"
	"go-user-api/internal/delivery/rest/middleware"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	AuthHandler    rest.IAuthHandler
	UserHandler    rest.IUserHandler
	AuthMiddleware *middleware.MiddlewareConfig
}

func (r *RouteConfig) Setup() {
	api := r.App.Group("/api/v1")

	// auth
	auth := api.Group("/auth")
	auth.Post("/register", r.AuthHandler.Register)
	auth.Post("/login", r.AuthHandler.Login)
	auth.Put("/", r.AuthMiddleware.Auth(), r.AuthHandler.ChangePassword)

	// user
	user := api.Group("/users", r.AuthMiddleware.Auth())
	user.Get("/:id", r.UserHandler.FindById)
	user.Delete("/:id", r.UserHandler.Delete)
}
