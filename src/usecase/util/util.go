package util

import (
	"log"
	"time"
)

func FormatDateTime(datetime time.Time) string {
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		location = time.UTC
		log.Fatal(err)
	}
	localDateTime := datetime.In(location)
	return localDateTime.Format("2006-01-02T15:04:05")
}
