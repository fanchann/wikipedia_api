package repository

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/fanchann/wikipedia_api/internals/entity"
	"github.com/fanchann/wikipedia_api/models"
)

type Repository[X any] struct {
	DB *gorm.DB
}

func (r *Repository[X]) Create(db *gorm.DB, entity *X) error {
	return db.Create(entity).Error
}

func (r *Repository[X]) Update(db *gorm.DB, entity *X) error {
	return db.Save(entity).Error
}

func (r *Repository[X]) FindByID(db *gorm.DB, entity *X, id any, columnName string) error {
	return db.Where(fmt.Sprintf("%s = ?", columnName), id).Take(entity).Error
}

func (r *Repository[X]) GetPageLists(db *gorm.DB, optsPaginate *models.OptsPaginate, entities *[]X) error {
	return db.Scopes(entity.NewPaginate(optsPaginate.Limit, optsPaginate.Page, optsPaginate.SortBy).PaginateResult).Find(entities).Error
}
