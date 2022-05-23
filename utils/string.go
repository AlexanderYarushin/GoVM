package utils

import (
	"strconv"
	"strings"
)

func ClearSpace(str string) string {
	for strings.Index(str, "  ") != -1 {
		str = strings.Replace(str, "  ", " ", -1)
	}

	str = strings.TrimSpace(str)
	return str
}

func IsInt(data string) bool {
	if _, err := strconv.Atoi(data); err == nil {
		return true
	}
	return false
}
