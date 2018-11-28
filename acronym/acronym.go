package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 2

func Abbreviate(s string) string {
	res := ""
	prev := " "

	for _, char := range s {
		if prev == " " || prev == "-" || unicode.IsUpper(char) {
			res = res + string(char)
		}
		if string(char) == ":" {
			break
		}
		prev = string(char)
	}
	return strings.ToUpper(res)
}
