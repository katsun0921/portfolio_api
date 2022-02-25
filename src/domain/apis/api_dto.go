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
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	NameJp      string        `json:"name_jp"`
	Programming []Programming `json:"programming"`
}

type Programming struct {
	Id        string `json:"language_id"`
	Language  string `json:"language"`
	Level     string `json:"level"`
	ColorCode string `json:"color_code"`
}

type Workexpress struct {
	Id          string   `json:"id"`
	Company     string   `json:"company"`
	Project     string   `json:"project"`
	JobType     string   `json:"job_type"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	Description string   `json:"description"`
	Skills      []string `json:"skills"`
}
