package entity

type Wiki struct {
	ID          string `gorm:"column:wiki_id"`
	Word        string `gorm:"column:wiki_word"`
	Description string `gorm:"column:wiki_description"`
}

func (w *Wiki) TableName() string {
	return "wikipedia"
}
