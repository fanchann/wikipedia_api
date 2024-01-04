package http

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/fanchann/wikipedia_api/internals/usecases"
	"github.com/fanchann/wikipedia_api/models"
	"github.com/fanchann/wikipedia_api/models/converter"
	"github.com/fanchann/wikipedia_api/models/web"
)

type IWikiController interface {
	CreateNewWiki(ctx *fiber.Ctx) error
	UpdateWiki(ctx *fiber.Ctx) error
	GetAllWiki(ctx *fiber.Ctx) error
	GetWikiByWord(ctx *fiber.Ctx) error
}

type WikiController struct {
	WikiUsecase *usecases.WikipediaUsecases
	Log         *logrus.Logger
}

func NewWikiController(wikiUsecase *usecases.WikipediaUsecases, log *logrus.Logger) *WikiController {
	return &WikiController{WikiUsecase: wikiUsecase, Log: log}
}

func (w *WikiController) CreateNewWiki(ctx *fiber.Ctx) error {
	wikiReq := new(web.CreateWikiRequestWeb)
	if err := ctx.BodyParser(wikiReq); err != nil {
		w.Log.WithError(err).Error("failed parse body")
		return ctx.Status(400).JSON("error while parse body")
	}
	wikiResponse, err := w.WikiUsecase.CreateNewWiki(ctx.UserContext(), converter.CreateWikiWebToWikiUsecase(wikiReq))
	if err != nil {
		w.Log.WithError(err).Error("failed store to database")
		return ctx.Status(fiber.StatusInternalServerError).JSON("failed to create")
	}
	return ctx.Status(200).JSON(map[string]interface{}{
		"success": true,
		"data":    wikiResponse,
	})
}

func (w *WikiController) UpdateWiki(ctx *fiber.Ctx) error {
	id := ctx.Params("wordId")
	wikiUpdateReq := new(web.UpdateWikiRequestWeb)
	wikiUpdateReq.ID = id
	if err := ctx.BodyParser(wikiUpdateReq); err != nil {
		w.Log.WithError(err).Error("failed parse body")
		return fiber.ErrBadRequest
	}
	wikiResponse, err := w.WikiUsecase.UpdateWiki(ctx.UserContext(), converter.UpdateCreateWikiWebToWikiUsecase(wikiUpdateReq))
	if err != nil {
		w.Log.WithError(err).Error("failed parse body")
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	return ctx.Status(200).JSON(map[string]interface{}{
		"success": true,
		"data":    wikiResponse,
	})
}

func (w *WikiController) GetAllWiki(ctx *fiber.Ctx) error {
	limit := ctx.Query("limit")
	page := ctx.Query("page")
	sortBy := ctx.Query("sort")

	limitInt, _ := strconv.Atoi(limit)
	pageInt, _ := strconv.Atoi(page)

	wikiResponse := w.WikiUsecase.GetAllWikipedia(ctx.UserContext(), &models.OptsPaginate{Limit: limitInt, Page: pageInt, SortBy: sortBy})

	nextPage := buildPageURL(ctx.BaseURL(), pageInt, 1)
	previousPage := buildPageURL(ctx.BaseURL(), pageInt, -1)

	return ctx.Status(fiber.StatusFound).JSON(web.WebResponse{
		NextPage:     nextPage,
		PreviousPage: previousPage,
		Rows_ttl:     int64(len(*wikiResponse)),
		Data:         wikiResponse,
	})

}

func (w *WikiController) GetWikiByWord(ctx *fiber.Ctx) error {
	word := ctx.Query("search")
	wikiResult, ttlData, err := w.WikiUsecase.GetWikiByWord(ctx.UserContext(), word)
	if err != nil {
		w.Log.WithError(err).Error("failed get word")
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(200).JSON(struct {
		Success bool        `json:"success"`
		RowsTtl int64       `json:"rows_ttl"`
		Data    interface{} `json:"data"`
	}{Success: true, RowsTtl: ttlData, Data: wikiResult})
}

func buildPageURL(baseURL string, pageInt int, direction int) string {
	if pageInt <= 0 {
		return ""
	}
	return fmt.Sprintf("%s/vi/api/wikipedia/_all?limit=30&page=%d", baseURL, pageInt+direction)
}
