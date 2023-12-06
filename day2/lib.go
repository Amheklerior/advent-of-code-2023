package day2

import (
	"fmt"
	"regexp"
	"strings"
)

func getCountsFor(str, color string) []string {
	digitsRegex := regexp.MustCompile(`\d+`)
	r := regexp.MustCompile(fmt.Sprintf(`\d+ %s`, color))
	return digitsRegex.FindAllString(strings.Join(r.FindAllString(str, -1), " "), -1)
}
