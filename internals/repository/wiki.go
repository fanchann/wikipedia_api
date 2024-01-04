package repository

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/fanchann/wikipedia_api/internals/entity"
	"github.com/fanchann/wikipedia_api/models"
)

type WikiRepository struct {
	Repository[entity.Wiki]
	Log *logrus.Logger
}

func NewWikiRepository(log *logrus.Logger) *WikiRepository {
	return &WikiRepository{Log: log}
}

func (w *WikiRepository) FindWikiByWord(db *gorm.DB, entities *[]entity.Wiki, optsPaginate models.OptsPaginate, word string) (int64, error) {
	var ttl int64
	db.Model(entities).Where("wiki_word LIKE ?", fmt.Sprintf("%%%s%%", word)).Count(&ttl)
	return ttl, db.Table("wikipedia").Scopes(entity.NewPaginate(optsPaginate.Limit, optsPaginate.Page, optsPaginate.SortBy).PaginateResult).Where("wiki_word LIKE ?", fmt.Sprintf("%%%s%%", word)).Find(entities).Error
}
