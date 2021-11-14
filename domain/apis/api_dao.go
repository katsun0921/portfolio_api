package apis

import (
  "errors"
  "fmt"
  "github.com/dghubble/go-twitter/twitter"
  "github.com/dghubble/oauth1"
  "github.com/joho/godotenv"
  "github.com/katsun0921/go_utils/logger"
  "github.com/katsun0921/go_utils/rest_errors"
  "github.com/mmcdole/gofeed"
  "os"
  "strconv"
)

const feedZenn = "https://zenn.dev/katsun0921/feed"

func (api *Api) GetRss() (*gofeed.Feed, rest_errors.RestErr) {
	feed, err := gofeed.NewParser().ParseURL(feedZenn)
	if err != nil {
		logger.Error("error when trying to rss", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get rss api", errors.New("response error"))
	}

	return feed, nil
}

func (api *Api) GetTwitter() ([]twitter.Tweet, rest_errors.RestErr) {

	envErr := godotenv.Load()
	if envErr != nil {
	  logger.Error("Error loading .env file", envErr)
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
  }
  
	tweets, httpResponse, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		UserID: toIntTwitterUserId,
		Count:  2,
	})

	for _, tweet := range tweets {
		fmt.Println(tweet.Text)
	}

	if err != nil {
		logger.Error(fmt.Sprintf("twitter server error %d", httpResponse.StatusCode), err)
	}
	resErr := rest_errors.NewBadRequestError("twitter api server error", errors.New("server error"))
	return tweets, resErr
}
