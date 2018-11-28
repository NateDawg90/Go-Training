package pangram

import (
	"strings"
)

const testVersion = 1

func IsPangram(s string) bool {
	alphabet := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	for _, char := range strings.Split(strings.ToLower(s), "") {
		for idx, letter := range alphabet {
			if letter == char {
				alphabet = append(alphabet[:idx], alphabet[idx+1:]...)
				break
			}
		}
	}

	return len(alphabet) == 0
}
