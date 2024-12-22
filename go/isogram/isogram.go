package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines whether a given word or phrase is an isogram.
func IsIsogram(word string) bool {
	existingLetters := map[rune]bool{}
	for _, letter := range strings.ToLower(word) {
		if !unicode.IsLetter(letter) {
			continue
		}

		if existingLetters[letter] {
			return false
		}

		existingLetters[letter] = true
	}
	return true
}
