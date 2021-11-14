package apis

type Api struct {
	Title       string `json:"title"`
	PlainText   string `json:"plain_text"`
	Link        string `json:"link"`
	Service     string `json:"service"`
	DateCreated string `json:"date_created"`
}
