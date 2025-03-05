package config

import (
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewFiber(viper *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:     viper.GetString("app.name"),
		Prefork:     viper.GetBool("server.prefork"),
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	return app
}
