package web

type WebResponse struct {
	NextPage     string      `json:"next_page"`
	PreviousPage string      `json:"previous_page"`
	Rows_ttl     int64       `json:"rows_ttl"`
	Data         interface{} `json:"data"`
}
