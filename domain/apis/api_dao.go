package apis

import (
  "errors"
  "github.com/katsun0921/go_utils/logger"
  "github.com/katsun0921/go_utils/rest_errors"
  "github.com/mmcdole/gofeed"
)

const feedZenn = "https://zenn.dev/katsun0921/feed"

func (api *Api) GetRss() (*gofeed.Feed, rest_errors.RestErr) {
  feed, err := gofeed.NewParser().ParseURL(feedZenn)
  if err != nil {
    logger.Error("error when trying to prepare get user statement", err)
    return nil, rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))
  }

  return feed, nil
}
