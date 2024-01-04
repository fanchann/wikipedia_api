package models

type WikipediaCreateUsecase struct {
	Word        string `validate:"required"`
	Description string `validate:"required"`
}

type WikipediaUpdateUsecase struct {
	WordID      string
	Word        string `validate:"required"`
	Description string `validate:"required"`
}

type WikipediaUsecaseResponse struct {
	WordID      string `json:"word_id"`
	Word        string `json:"word"`
	Description string `json:"description"`
}
