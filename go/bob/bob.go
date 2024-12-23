package bob

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Hey determines bobs response to a remark (442.9 ns/op)
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	lastRune, _ := utf8.DecodeLastRuneInString(remark)

	switch {
	case remark == "":
		return "Fine. Be that way!"

	case IsShouting(remark):
		if lastRune == '?' {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"

	case lastRune == '?':
		return "Sure."

	default:
		return "Whatever."
	}
}

// IsShouting checks if a remark has at least one letter, and if all letters are uppercase
//
// you should also do `strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark`,
// but that increases the ns/op by more than 600%
func IsShouting(remark string) bool {
	hasLetters := false
	for _, character := range remark {
		if unicode.IsLower(character) {
			return false
		}
		if unicode.IsLetter(character) {
			hasLetters = true
		}
	}
	return hasLetters
}
