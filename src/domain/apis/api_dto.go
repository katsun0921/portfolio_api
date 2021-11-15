package apis

type Api struct {
  Text        string `json:"text"`
  Link        string `json:"link"`
  Service     string `json:"service"`
  DateCreated string `json:"date_created"`
  DateUnix    int    `json:"date_unix"`
}
