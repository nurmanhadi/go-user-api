package main

import (
	"context"
	"fmt"
	"go-user-api/config"
)

func main() {
	ctx := context.Background()
	viper := config.NewViper()
	db := config.NewMysql(viper)
	log := config.NewLogger()
	validation := config.NewValidator()
	app := config.NewFiber(viper)
	config.Bootstrap(&config.BootstrapConfig{
		Viper:      viper,
		DB:         db,
		App:        app,
		Validation: validation,
		Ctx:        ctx,
		Log:        log,
	})
	host := viper.GetString("server.host")
	port := viper.GetInt("server.port")
	err := app.Listen(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("failed start web server: %v", err)
	}
}
