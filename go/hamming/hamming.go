package hamming

import "errors"

// Distance calculates the hamming distance between two DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("a and b is not the same length >:c")
	}

	hammingDistance := 0
	// comparing runes instead of bytes (like a[i]) is 6 nanoseconds faster 😼
	// (this will definitely make a huge difference)
	for i, letter := range a {
		if letter != rune(b[i]) {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
