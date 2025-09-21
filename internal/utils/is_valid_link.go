package utils

import "regexp"

func IsValidLink(link string) bool {
	return regexp.MustCompile(`^(https?|ftp|http)://[^\s/$.?#].[^\s]*$`).MatchString(link)
}
