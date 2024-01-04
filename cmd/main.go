package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/fanchann/wikipedia_api/internals/config"
)

func main() {
	viperCfg := config.NewViper()
	logCfg := config.NewLogrus()
	gormCfg := config.NewGorm(viperCfg, logCfg)
	validatorCfg := validator.New()
	app := fiber.New(fiber.Config{
		AppName: viperCfg.GetString("APP_NAME"),
		Prefork: true,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, api-key",
	}))

	config.NewAppConfiguation(&config.AppConfiguration{
		DB:       gormCfg,
		App:      app,
		Log:      logCfg,
		Validate: validatorCfg,
		Config:   viperCfg,
	})

	appPort := fmt.Sprintf(":%d", viperCfg.GetInt("APP_PORT"))
	app.Listen(appPort)
}
