package usecases

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/fanchann/wikipedia_api/internals/entity"
	"github.com/fanchann/wikipedia_api/internals/repository"
	"github.com/fanchann/wikipedia_api/internals/utils"
	"github.com/fanchann/wikipedia_api/models"
)

type IWikipediaUsecases interface {
	GetAllWikipedia(ctx context.Context, lists *models.OptsPaginate) []*models.WikipediaUsecaseResponse
	GetWikiByWord(ctx context.Context, word string) ([]*models.WikipediaUsecaseResponse, error)
	CreateNewWiki(ctx context.Context, req *models.WikipediaCreateUsecase) (*models.WikipediaUsecaseResponse, error)
	UpdateWiki(ctx context.Context, req *models.WikipediaUpdateUsecase) (*models.WikipediaUsecaseResponse, error)
}

type WikipediaUsecases struct {
	WikiRepo *repository.WikiRepository
	DB       *gorm.DB
	Log      *logrus.Logger
	Validate *validator.Validate
}

func NewWikipediaUsecases(wikiRepo *repository.WikiRepository, db *gorm.DB, log *logrus.Logger, validate *validator.Validate) *WikipediaUsecases {
	return &WikipediaUsecases{
		WikiRepo: wikiRepo,
		DB:       db,
		Log:      log,
		Validate: validate,
	}
}

func (w *WikipediaUsecases) GetAllWikipedia(ctx context.Context, lists *models.OptsPaginate) *[]models.WikipediaUsecaseResponse {
	tx := w.DB.Begin()
	defer tx.Rollback()

	var wikis []entity.Wiki
	if err := w.WikiRepo.GetPageLists(tx, lists, &wikis); err != nil {
		w.Log.Warnf("data wikipedia not found :%v", err)
		return nil
	}
	var wikiDatas []models.WikipediaUsecaseResponse

	for _, wiki := range wikis {
		wikiDatas = append(wikiDatas, models.WikipediaUsecaseResponse{
			WordID:      wiki.ID,
			Word:        wiki.Word,
			Description: wiki.Description,
		})
	}

	if err := tx.Commit().Error; err != nil {
		w.Log.Warnf("Error while commit : %v", err)
		return nil
	}

	return &wikiDatas
}

func (w *WikipediaUsecases) GetWikiByWord(ctx context.Context, word string) (*[]models.WikipediaUsecaseResponse, int64, error) {
	tx := w.DB.Begin()
	defer tx.Rollback()

	var wikis []entity.Wiki
	var ttl int64
	ttl, err := w.WikiRepo.FindWikiByWord(tx, &wikis, models.OptsPaginate{}, word)
	if err != nil {
		return nil, 0, fiber.ErrNotFound
	}

	var wikiDatas []models.WikipediaUsecaseResponse

	for _, wiki := range wikis {
		wikiDatas = append(wikiDatas, models.WikipediaUsecaseResponse{
			WordID:      wiki.ID,
			Word:        wiki.Word,
			Description: wiki.Description,
		})
	}

	if err := tx.Commit().Error; err != nil {
		w.Log.Warnf("Error while commit : %v", err)
		return nil, 0, fiber.ErrInternalServerError
	}

	return &wikiDatas, ttl, nil
}

func (w *WikipediaUsecases) CreateNewWiki(ctx context.Context, req *models.WikipediaCreateUsecase) (*models.WikipediaUsecaseResponse, error) {
	tx := w.DB.Begin()
	defer tx.Rollback()

	if err := w.Validate.Struct(req); err != nil {
		w.Log.Warnf("Error while parse body")
		return nil, fiber.ErrBadRequest
	}

	wikiCreateReq := &entity.Wiki{
		ID:          utils.IDGenerator(),
		Word:        req.Word,
		Description: req.Description,
	}
	if err := w.WikiRepo.Create(tx, wikiCreateReq); err != nil {
		w.Log.Warnf("error while store to database %v", err)
		return nil, fiber.ErrBadRequest
	}

	if err := tx.Commit().Error; err != nil {
		w.Log.Warnf("Error while commit : %v", err)
		return nil, fiber.ErrInternalServerError
	}

	return &models.WikipediaUsecaseResponse{
		WordID:      wikiCreateReq.ID,
		Word:        wikiCreateReq.Word,
		Description: wikiCreateReq.Description,
	}, nil
}

func (w *WikipediaUsecases) UpdateWiki(ctx context.Context, req *models.WikipediaUpdateUsecase) (*models.WikipediaUsecaseResponse, error) {
	tx := w.DB.Begin()
	defer tx.Rollback()
	if err := w.Validate.Struct(req); err != nil {
		w.Log.Warnf("Error while parse body")
		return nil, fiber.ErrBadRequest
	}

	var wikiResult entity.Wiki
	if err := w.WikiRepo.FindByID(tx, &wikiResult, req.WordID, "wiki_id"); err != nil {
		return nil, fiber.ErrNotFound
	}
	wikiResult.ID = req.WordID
	wikiResult.Word = req.Word
	wikiResult.Description = req.Description

	if err := w.WikiRepo.Repository.Update(tx, &wikiResult); err != nil {
		return nil, fiber.ErrBadRequest
	}

	if err := tx.Commit().Error; err != nil {
		w.Log.Warnf("Error while commit : %v", err)
		return nil, fiber.ErrInternalServerError
	}

	return &models.WikipediaUsecaseResponse{
		WordID:      wikiResult.ID,
		Word:        wikiResult.Word,
		Description: wikiResult.Description,
	}, nil
}
