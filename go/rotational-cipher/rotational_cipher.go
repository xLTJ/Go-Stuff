package rotationalcipher

import (
	"strings"
	"unicode"
)

// RotationalCipher performs the rotational cipher on a string (1069 ns/op)
func RotationalCipher(plain string, shiftKey int) string {
	var builder = strings.Builder{}
	for _, character := range plain {
		if !unicode.IsLetter(character) {
			builder.WriteRune(character)
			continue
		}

		isUppercase := unicode.IsUpper(character)
		character += int32(shiftKey)

		switch {
		case isUppercase:
			if character > 'Z' {
				character -= 26
			}
		default:
			if character > 'z' {
				character -= 26
			}
		}
		builder.WriteRune(character)
	}
	return builder.String()
}
