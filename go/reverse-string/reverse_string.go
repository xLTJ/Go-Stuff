package reverse

import "strings"

func Reverse(input string) string {
	inputSplice := []rune(input)

	var builder strings.Builder
	builder.Grow(len(inputSplice))
	for i := len(inputSplice) - 1; i >= 0; i-- {
		builder.WriteRune(inputSplice[i])
	}
	return builder.String()
}
