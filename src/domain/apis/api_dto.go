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
	Id     string     `json:"id"`
	Job    string     `json:"job"`
	Skills []Language `json:"skills"`
}

type Language struct {
	Name  string `json:"name"`
	Level string `json:"level"`
}

type WorkExpress struct {
	Id          string `json:"id"`
	Company     string `json:"company"`
	Project     string `json:"project"`
	JobType     string `json:"job_type"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Description string `json:"description"`
	Skills      string `json:"skills"`
}
