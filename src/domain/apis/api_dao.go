package apis

import (
	"errors"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
  "github.com/katsun0921/go_utils/logger"
	"github.com/katsun0921/go_utils/rest_errors"
	"github.com/katsun0921/portfolio_api/src/constants"
	"github.com/katsun0921/portfolio_api/src/domain/articles"
  "github.com/katsun0921/portfolio_api/src/lib/google_sheets_api"
	"github.com/mmcdole/gofeed"
	"os"
	"strconv"
)

type TUserTimelineParams struct {
	UserID  int64
	Count   int
	SinceID int64
}

func (api *Api) GetFeedApi(service string) (*gofeed.Feed, rest_errors.RestErr) {
	var url string
	switch service {
	case constants.ZENN:
		url = constants.FeedZenn
	default:
		return nil, rest_errors.NewInternalServerError("error not found rss service", errors.New("rss service error"))
	}
	feed, err := gofeed.NewParser().ParseURL(url)
	if err != nil {
		logger.Error("error when trying to rss", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get rss api", errors.New("response error"))
	}

	return feed, nil
}

func (api *Api) GetTwitterApi() ([]twitter.Tweet, rest_errors.RestErr) {

	envErr := godotenv.Load()
	if envErr != nil {
		logger.Error("Error loading .env file", envErr)
		return nil, rest_errors.NewInternalServerError("Error env file", envErr)
	}

	twitterApiKey := os.Getenv("TWITTER_API_KEY")
	twitterApiKeySecret := os.Getenv("TWITTER_API_KEY_SECRET")
	twitterAccessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	twitterAccessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	twitterUserId := os.Getenv("TWITTER_USER_ID")

	config := oauth1.NewConfig(twitterApiKey, twitterApiKeySecret)
	token := oauth1.NewToken(twitterAccessToken, twitterAccessTokenSecret)
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// twitter client
	client := twitter.NewClient(httpClient)

	// Status Show
	toIntTwitterUserId, errUserId := strconv.ParseInt(twitterUserId, 10, 64)
	if errUserId != nil {
		logger.Error("Error loading User Id from .env file", errUserId)
		return nil, rest_errors.NewInternalServerError("twitter server error to user id", errUserId)
	}

	result := &articles.Article{}

	articleId, dbErr := result.FindByLatestArticleId(constants.TWITTER)
	if dbErr != nil {
		logger.Error("Error loading Twitter Sever", dbErr)
		return nil, rest_errors.NewInternalServerError("twitter server", dbErr)
	}

	// Set env value if not find article id
	if articleId == "" {
		articleId = os.Getenv("TWITTER_LAST_ID")
	}

	articleIdInt64, _ := strconv.ParseInt(articleId, 10, 64)

	p := TUserTimelineParams{
		UserID:  toIntTwitterUserId,
		Count:   constants.ArticlesMaxCount,
		SinceID: articleIdInt64,
	}

	tweets, httpResponse, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		UserID:  p.UserID,
		Count:   p.Count,
		SinceID: p.SinceID,
	})

	if err != nil {
		logger.Error(fmt.Sprintf("twitter server error %d", httpResponse.StatusCode), err)
		return nil, rest_errors.NewInternalServerError(fmt.Sprintf("twitter server error %d", httpResponse.StatusCode), err)
	}
	return tweets, nil
}

func (skill *Skill) GetGoogleSheetsApi() ([][]interface{}, rest_errors.RestErr) {
  api := google_sheets_api.Main()

  fmt.Println(api)
  return api, nil
}
