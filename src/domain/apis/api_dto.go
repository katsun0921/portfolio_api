package apis

//:TODO Change Blog
type Api struct {
  Id          string `json:"id"`
  Text        string `json:"text"`
  Link        string `json:"link"`
  Service     string `json:"service"`
  DateCreated string `json:"date_created"`
  DateUnix    int    `json:"date_unix"`
}

type Skill struct {
  Id          string `json:"id"`
  Text        string `json:"text"`
  Link        string `json:"link"`
  Service     string `json:"service"`
  DateCreated string `json:"date_created"`
  DateUnix    int    `json:"date_unix"`
}
