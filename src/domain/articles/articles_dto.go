package articles

const (
  TypeTwitter = "twitter"
  TypeZenn = "zenn"
)

type Article struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Type      string `json:"type"`
	Password    string `json:"password"`
}

type Articles []Article
