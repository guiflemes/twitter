package date

import "time"

const (
	apiDateFormat string = "2006-01-02T15:04:05Z"
)

func GetNow ()time.Time{
	return time.Now().UTC()
}

func GetNowAsString () string{
	return GetNow().Format(apiDateFormat)
}