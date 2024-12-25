package isbn

// IsValidISBN checks if a string is a valid ISBN (178 ns/op)
func IsValidISBN(isbn string) bool {
	comparisonNumber := 0
	digit := 0
	for _, character := range isbn {
		switch {
		case character == '-':
			continue

		case character >= '0' && character <= '9':
			comparisonNumber += int(character-'0') * (10 - digit)

		case character == 'X':
			if digit != 9 {
				return false
			}
			comparisonNumber += 10 * (10 - digit)

		default:
			return false
		}

		digit++
	}
	return comparisonNumber%11 == 0 && digit == 10
}
