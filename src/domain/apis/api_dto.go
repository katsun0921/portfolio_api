package apis

type Api struct {
	Id          int    `json:"id"`
	Text        string `json:"text"`
	Link        string `json:"link"`
	Service     string `json:"service"`
	DateCreated string `json:"date_created"`
}
