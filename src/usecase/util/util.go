package util

import (
	"log"
	"strconv"
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

func FormatDateSeed() int {
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		location = time.UTC
		log.Fatal(err)
	}
	localDateTime := time.Now().In(location)
	date := localDateTime.Format("20060102")
	seed, err := strconv.Atoi(date)
	if err != nil {
		panic(err)
	}
	return seed
}
