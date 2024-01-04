package web

type CreateWikiRequestWeb struct {
	Word        string `json:"word"`
	Description string `json:"description"`
}

type UpdateWikiRequestWeb struct {
	ID          string
	Word        string `json:"word"`
	Description string `json:"description"`
}
