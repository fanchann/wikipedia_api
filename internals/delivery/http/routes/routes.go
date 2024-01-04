package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/fanchann/wikipedia_api/internals/delivery/http"
)

type RoutesConfiguration struct {
	App            *fiber.App
	WikiController *http.WikiController
	Middleware     fiber.Handler
}

func (r *RoutesConfiguration) Setup() {
	r.SetupGuestRoute()
	r.SetupAuthRoute()
}

func (r *RoutesConfiguration) SetupGuestRoute() {
	r.App.Get("/vi/api/wikipedia/_all", r.WikiController.GetAllWiki)
	r.App.Get("/vi/api/wikipedia", r.WikiController.GetWikiByWord)
}

func (r *RoutesConfiguration) SetupAuthRoute() {
	r.App.Use(r.Middleware)
	r.App.Post("/vi/api/wikipedia/_new", r.WikiController.CreateNewWiki)
	r.App.Put("/vi/api/wikipedia/_edit/:wordId", r.WikiController.UpdateWiki)
}
