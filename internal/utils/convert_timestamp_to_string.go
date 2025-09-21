package utils

import (
	"fmt"
	"time"
)

func ConvertTimestampToString(value time.Time) string {
	return fmt.Sprint(value.Unix())
}
