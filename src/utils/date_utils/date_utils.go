package date_utils

import (
	"github.com/katsun0921/portfolio_api/src/constants"
	"time"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(constants.DateLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(constants.DBLayout)
}

func SetUnixParse(layout string, value string) int {
	t, _ := time.Parse(layout, value)
	return int(t.Unix())
}
