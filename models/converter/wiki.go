package converter

import (
	"github.com/fanchann/wikipedia_api/models"
	"github.com/fanchann/wikipedia_api/models/web"
)

func CreateWikiWebToWikiUsecase(req *web.CreateWikiRequestWeb) *models.WikipediaCreateUsecase {
	return &models.WikipediaCreateUsecase{
		Word:        req.Word,
		Description: req.Description,
	}
}

func UpdateCreateWikiWebToWikiUsecase(req *web.UpdateWikiRequestWeb) *models.WikipediaUpdateUsecase {
	return &models.WikipediaUpdateUsecase{
		WordID:      req.ID,
		Word:        req.Word,
		Description: req.Description,
	}
}
