package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/fanchann/wikipedia_api/internals/delivery/http"
	"github.com/fanchann/wikipedia_api/internals/delivery/http/middlewares"
	"github.com/fanchann/wikipedia_api/internals/delivery/http/routes"
	"github.com/fanchann/wikipedia_api/internals/repository"
	"github.com/fanchann/wikipedia_api/internals/usecases"
)

type AppConfiguration struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func NewAppConfiguation(appConfig *AppConfiguration) {
	wikiRepo := repository.NewWikiRepository(appConfig.Log)

	wikiUsecase := usecases.NewWikipediaUsecases(wikiRepo, appConfig.DB, appConfig.Log, appConfig.Validate)

	wikiController := http.NewWikiController(wikiUsecase, appConfig.Log)

	authMiddleware := middlewares.NewAuthMiddlewares(appConfig.Log)

	routeConfiguration := routes.RoutesConfiguration{
		App:            appConfig.App,
		WikiController: wikiController,
		Middleware:     authMiddleware,
	}

	routeConfiguration.Setup()
}
