package util

import (
	"strconv"
	"strings"
	"time"
)

func StrToTime(s string) time.Time {
	date := strings.Split(s, "/")
	year, _ := strconv.Atoi(date[0])
	month, _ := strconv.Atoi(date[1])
	day, _ := strconv.Atoi(date[2])
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
