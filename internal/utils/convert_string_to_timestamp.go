package utils

import (
	"log"
	"strconv"
	"time"
)

func ConvertStringToTimestamp(value string) time.Time {
	tsStr := value
	sec, err := strconv.ParseInt(tsStr, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return time.Unix(sec, 0)
}
