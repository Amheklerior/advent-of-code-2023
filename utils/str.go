package utils

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func ToInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("Error converting '%v' into integer", str)
	}
	return value
}

func GetOccurrence(str, regex string) string {
	return regexp.MustCompile(regex).FindString(str)
}

func GetOccurrences(str, regex string) []string {
	return regexp.MustCompile(regex).FindAllString(str, -1)
}

func ExtractPrefix(line, prefixRegex string) (string, string) {
	prefix := GetOccurrence(line, prefixRegex)
	withoutPrefix := strings.TrimPrefix(line, prefix)
	return withoutPrefix, prefix
}

func ExtractSuffix(line, suffixRegex string) (string, string) {
	suffix := GetOccurrence(line, suffixRegex)
	withoutSuffix := strings.TrimSuffix(line, suffix)
	return withoutSuffix, suffix
}
