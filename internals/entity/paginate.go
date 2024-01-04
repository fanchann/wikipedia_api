package entity

import "gorm.io/gorm"

type IPaginate interface {
	GetOffset() int
	GetLimit() int
	GetPage() int
	GetSortBy() string
	PaginateResult(db *gorm.DB) *gorm.DB
}

type Paginate struct {
	Limit  int
	Page   int
	SortBy string
}

func NewPaginate(limit, page int, sortBy string) *Paginate {
	return &Paginate{Limit: limit, Page: page, SortBy: sortBy}
}

func (p *Paginate) GetOffset() int {
	offset := (p.GetPage() - 1) * p.GetLimit()
	return offset
}

func (p *Paginate) GetLimit() int {
	if p.Limit == 0 {
		return 30
	}
	return p.Limit
}

func (p *Paginate) GetPage() int {
	if p.Page == 0 {
		return 1
	}
	return p.Page
}

func (p *Paginate) GetSortBy() string {
	switch p.SortBy {
	case "desc":
		return "wiki_word desc"
	default:
		return "wiki_word asc"
	}
}

func (p *Paginate) PaginateResult(db *gorm.DB) *gorm.DB {
	return db.Offset(p.GetOffset()).Limit(p.GetLimit()).Order(p.GetSortBy())
}
