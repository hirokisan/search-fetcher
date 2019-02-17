package list

import (
	"regexp"
	"strings"
)

func AdjustUrl(u string) string {
	u = strings.Replace(u, "/url?q=", "", 1)
	re, _ := regexp.Compile("(&sa).+")
	u = re.ReplaceAllString(u, "")
	return u
}

func AdjustText(t string) string {
	return t
}
