package apis

type Api struct {
	Title       string      `json:"title"`
	Description Description `json:"description"`
	Link        string      `json:"link"`
	Service     string      `json:"service"`
	Type        string      `json:"type"`
	DateCreated string      `json:"date_created"`
}

type Description struct {
	PlainText string `json:"plain_text"`
	Html      string `json:"html"`
}
