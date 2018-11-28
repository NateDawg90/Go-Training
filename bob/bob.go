package bob // package name must match the package name in bob_test.go

import "strings"

const testVersion = 2 // same as targetTestVersion

func Hey(s string) string {
	s = strings.TrimSpace(s)

	//check if blank
	if len(s) == 0 {
		return "Fine. Be that way!"
	}

	// Check if yelling

	if strings.ToUpper(s) == s && strings.ToLower(s) != s {
		return "Whoa, chill out!"
	}
	// var foundLetter bool = false
	// for i := 0; i < len(s); i++ {

	// 	// break if lowercase letter found
	// 	if s[i] >= 97 && s[i] <= 122 {
	// 		break
	// 	}

	// 	// at least one capital letter found
	// 	if s[i] >= 65 && s[i] <= 90 {
	// 		foundLetter = true
	// 	}

	// 	// if it doesn't break and has found a capital letter
	// 	if i == len(s)-1 && foundLetter == true {
	// 		return "Whoa, chill out!"
	// 	}
	// }

	// check if question
	if s[len(s)-1:] == "?" {
		return "Sure."
	}

	return "Whatever."

}
