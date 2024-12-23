package pangram

import (
	"strings"
)

// IsPangram checks if a string is a pangram
func IsPangram(input string) bool {
	lowercaseInput := strings.ToLower(input)
	for char := 'a'; char <= 'z'; char++ {
		if !strings.ContainsRune(lowercaseInput, char) {
			return false
		}
	}
	return true
}
