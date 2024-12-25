package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(words []string) []string {
	var rhyme []string

	if len(words) == 0 {
		return rhyme
	}

	for i, word := range words {
		if i+1 >= len(words) {
			break
		}
		rhyme = append(rhyme, fmt.Sprintf("For want of a %s the %s was lost.", word, words[i+1]))
	}

	rhyme = append(rhyme, fmt.Sprintf("And all for the want of a %s.", words[0]))
	return rhyme
}

// 3992 ns/op
